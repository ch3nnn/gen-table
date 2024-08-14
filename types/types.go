/**
 * @Author: chentong
 * @Date: 2024/07/31 上午10:45
 */

package types

import (
	"html/template"

	"gorm.io/gen/field"
)

// SourceCode source code
type SourceCode int

const (
	// Struct ...
	Struct SourceCode = iota
	// Table ...
	Table
	// Object ...
	Object
)

type Param struct {
	PkgPath   string
	Package   string
	Name      string
	Type      string
	IsArray   bool
	IsPointer bool
}

type Method struct {
	Receiver   Param
	MethodName string
	Doc        string
	Params     []Param
	Result     []Param
	Body       string
}

type Field struct {
	Name             string
	Type             string
	ColumnName       string
	ColumnComment    string
	MultilineComment bool
	Tag              field.Tag
	GORMTag          field.GormTag
	CustomGenType    string
	Relation         *field.Relation
}
type QueryStructMeta struct {
	Command         template.HTML
	ImportPaths     []template.HTML
	ModelImportPath template.HTML
	QueryImportPath template.HTML
	DaoFileName     string

	Generated       bool   // whether to generate db model
	FileName        string // generated file name
	S               string // the first letter(lower case)of simple Name (receiver)
	QueryStructName string // internal query struct name
	ModelStructName string // origin/model struct name
	TableName       string // table name in db server
	TableComment    string // table comment in db server
	StructInfo      Param
	Fields          []*Field
	Source          SourceCode
	ImportPkgPaths  []string
	ModelMethods    []*Method // user custom method bind to db base struct
}
