package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/duke-git/lancet/v2/strutil"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

// DBType database type
type DBType string

const (
	// dbMySQL Gorm Drivers mysql || postgres || sqlite || sqlserver
	dbMySQL      DBType = "mysql"
	dbPostgres   DBType = "postgres"
	dbSQLite     DBType = "sqlite"
	dbSQLServer  DBType = "sqlserver"
	dbClickHouse DBType = "clickhouse"
)

var (
	dsn     string
	outPath string
	tables  string
	db      string

	updateTimeFieldName string
	createTimeFieldName string
	deleteTimeFieldName string

	isGenCRUD bool
)

func init() {
	flag.StringVar(&dsn, "dsn", "", `consult[https://gorm.io/docs/connecting_to_the_database.html]`)
	flag.StringVar(&db, "db", "mysql", `input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]`)
	flag.StringVar(&outPath, "outPath", "./dal", `specify a directory for output`)
	flag.StringVar(&tables, "tables", "", `enter the required data table or leave it blank`)
	flag.StringVar(&updateTimeFieldName, "updateTimeField", "", `auto update time field name`)
	flag.StringVar(&createTimeFieldName, "createTimeField", "", `auto create time field name`)
	flag.StringVar(&deleteTimeFieldName, "deleteField", "", `delete time field name`)
	flag.BoolVar(&isGenCRUD, "isGenCRUD", false, `is gen curd func`)

	flag.Parse()

}

func main() {
	db, err := connectDB(DBType(db), dsn)
	if err != nil {
		log.Fatalln("connect db server fail:", err)
	}

	// 生成实例
	g := gen.NewGenerator(gen.Config{
		// 模型包名路径
		OutPath:      outPath + "/query",
		ModelPkgPath: outPath + "/model",

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: false, // generate pointer when field is nullable

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false, // detect integer field's unsigned type, adjust generated data type
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag
		// 生成单元测试
		WithUnitTest: false,
	})

	// 自定义字段的数据类型
	// 统一数字类型为int64, 兼容protobuf
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(
		map[string]func(columnType gorm.ColumnType) (dataType string){
			"int":       func(columnType gorm.ColumnType) (dataType string) { return "int64" },
			"int2":      func(columnType gorm.ColumnType) (dataType string) { return "int64" },
			"int4":      func(columnType gorm.ColumnType) (dataType string) { return "int64" },
			"int8":      func(columnType gorm.ColumnType) (dataType string) { return "int64" },
			"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int64" },
			"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
			"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
			"tinyint": func(columnType gorm.ColumnType) (dataType string) {
				// mysql tinyint(1) 类型 改为 bool
				ct, _ := columnType.ColumnType()
				if strings.HasPrefix(ct, "tinyint(1)") {
					return "bool"
				}
				return "int64"
			},
			// 统一日期类型为 sql.NullTime
			"datetime": func(columnType gorm.ColumnType) (dataType string) { return "*time.Time" },
			"json":     func(columnType gorm.ColumnType) (dataType string) { return "datatypes.JSON" },
		})

	// 自定义模型结体字段的标签
	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	// jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
	//	toStringField := `balance, `
	//	if strings.Contains(toStringField, columnName) {
	//		return columnName + ",string"
	//	}
	//	return columnName
	// })
	// 将非默认字段名的字段定义为自动时间戳和软删除字段;

	g.UseDB(db)

	models, err := genModels(g, db, tables)
	if err != nil {
		log.Fatalln("get tables info fail:", err)
	}

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	g.ApplyBasic(models...)
	//g.Execute()

}

func modelOpt() (modelOpts []gen.ModelOpt) {
	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
	if deleteTimeFieldName != "" {
		softDeleteField := gen.FieldType(deleteTimeFieldName, "gorm.DeletedAt")
		modelOpts = append(modelOpts, softDeleteField)
	}
	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	if updateTimeFieldName != "" {
		autoUpdateTimeField := gen.FieldGORMTag(updateTimeFieldName, func(tag field.GormTag) field.GormTag {
			return field.GormTag{"column": []string{updateTimeFieldName}, "type": []string{"int unsigned"}, "": []string{"autoUpdateTime"}}
		})
		modelOpts = append(modelOpts, autoUpdateTimeField)
	}
	if createTimeFieldName != "" {
		autoCreateTimeField := gen.FieldGORMTag(createTimeFieldName, func(tag field.GormTag) field.GormTag {
			return field.GormTag{"column": []string{createTimeFieldName}, "type": []string{"int unsigned"}, "": []string{"autoCreateTime"}}
		})
		modelOpts = append(modelOpts, autoCreateTimeField)
	}

	return modelOpts

}

func connectDB(t DBType, dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return nil, fmt.Errorf("dsn cannot be empty")
	}

	switch t {
	case dbMySQL:
		return gorm.Open(mysql.Open(dsn))
	case dbPostgres:
		return gorm.Open(postgres.Open(dsn))
	case dbSQLite:
		return gorm.Open(sqlite.Open(dsn))
	case dbSQLServer:
		return gorm.Open(sqlserver.Open(dsn))
	case dbClickHouse:
		return gorm.Open(clickhouse.Open(dsn))
	default:
		return nil, fmt.Errorf("unknow db %q (support mysql || postgres || sqlite || sqlserver for now)", t)
	}
}

// genModels is gorm/gen generated models
func genModels(g *gen.Generator, db *gorm.DB, tableSting string) (models []interface{}, err error) {
	var tables []string
	if tableSting == "" {
		// Execute tasks for all tables in the database
		tables, err = db.Migrator().GetTables()
		if err != nil {
			return nil, fmt.Errorf("GORM migrator get all tables fail: %w", err)
		}
	} else {
		tables = strings.Split(tableSting, ",")
	}

	// Execute some data table tasks
	models = make([]interface{}, len(tables))
	for i, tableName := range tables {
		model := g.GenerateModel(tableName, modelOpt()...)
		models[i] = model
		if isGenCRUD {
			template.Must(template.New("test").Funcs(map[string]any{
				"CamelCase":  strutil.CamelCase,
				"LowerFirst": strutil.LowerFirst,
			}).Parse(interfaceTmpl)).Execute(os.Stdout, model)
		}

	}

	return models, nil
}

//  -dsn "root:gY3-Okw-Z9H-M2F@tcp(192.168.16.7:3306)/proxy_bmt?charset=utf8mb4&parseTime=True&loc=Local"  -updateTimeField "mtime" -createTimeField "ctime" -outPath "../../../internal/dal"

const interfaceTmpl = `
{{ $modelStructName := .ModelStructName }}



var _ I{{.ModelStructName}} = (*{{.ModelStructName}}Dao)(nil)

type I{{.ModelStructName}} interface {
	{{ range .Fields }}
	WhereBy{{.Name}}({{.ColumnName | CamelCase }} *{{.Type}}) func(dao gen.Dao) gen.Dao{{ end }}

	Create(m *model.{{.ModelStructName}}) (*model.{{.ModelStructName}}, error)
	Delete(whereFunc ...func(dao gen.Dao) gen.Dao) error
	DeletePhysical(whereFunc ...func(dao gen.Dao) gen.Dao) error
	Update(m interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (*model.{{.ModelStructName}}, error)
	Select(whereFunc ...func(dao gen.Dao) gen.Dao) (*model.{{.ModelStructName}}, error)
	SelectList(whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.{{.ModelStructName}}, error)
	SelectPage(offset int, limit int, whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.{{.ModelStructName}}, int64, error)
}

type {{.ModelStructName}}Dao struct {
	{{.ModelStructName | LowerFirst }}Do query.I{{.ModelStructName}}Do
}

func New{{.ModelStructName}}Dao(ctx context.Context) {{.ModelStructName}}Dao {
	return {{.ModelStructName}}Dao{ {{.ModelStructName | LowerFirst }}Do: query.{{.ModelStructName}}.WithContext(ctx)}
}

{{ range .Fields }}
func (s *{{$modelStructName}}Dao) WhereBy{{.Name}}({{.ColumnName | CamelCase }} *{{.Type}}) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		if {{.ColumnName | CamelCase }} != nil {
			return dao.Where(query.{{$modelStructName}}.{{.Name}}.Eq(*{{.ColumnName | CamelCase }}))
		}
		return dao
    }
}
{{ end }}


func (s *{{.ModelStructName}}Dao) Create(m *model.{{.ModelStructName}}) (*model.{{.ModelStructName}}, error) {
	if err := s.{{.ModelStructName| LowerFirst }}Do.Create(m); err != nil {
		return nil, err
	}
	return s.Select(s.WhereByID(&m.ID))
}

func (s *{{.ModelStructName}}Dao) Select(whereFunc ...func(dao gen.Dao) gen.Dao) (*model.{{.ModelStructName}}, error) {
	return s.{{.ModelStructName| LowerFirst }}Do.Scopes(whereFunc...).First()
}

func (s *{{.ModelStructName}}Dao) SelectList(whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.{{.ModelStructName}}, error) {
	return s.{{.ModelStructName| LowerFirst }}Do.Scopes(whereFunc...).Find()
}

func (s *{{.ModelStructName}}Dao) SelectPage(offset int, limit int, whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.{{.ModelStructName}}, int64, error) {
	return s.{{.ModelStructName| LowerFirst }}Do.Scopes(whereFunc...).FindByPage(offset, limit)
}

func (s *{{.ModelStructName}}Dao) Update(m interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (*model.{{.ModelStructName}}, error) {
	toMap, err := structs.ToMap(m)
	if err != nil {
		return nil, err
	}

	if _, err := s.{{.ModelStructName| LowerFirst }}Do.Scopes(whereFunc...).Updates(toMap); err != nil {
		return nil, err
	}

	return s.Select(whereFunc...)
}

func (s *{{.ModelStructName}}Dao) Delete(whereFunc ...func(dao gen.Dao) gen.Dao) error {
	if _, err := s.{{.ModelStructName| LowerFirst }}Do.Scopes(whereFunc...).Delete(); err != nil {
		return err
	}
	return nil
}

func (s *{{.ModelStructName}}Dao) DeletePhysical(whereFunc ...func(dao gen.Dao) gen.Dao) error {
	if _, err := s.{{.ModelStructName| LowerFirst }}Do.Unscoped().Scopes(whereFunc...).Delete(); err != nil {
		return err
	}
	return nil

`
