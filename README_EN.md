# gen-table

English | [简体中文](README.md)

Generate model, query, and dao codes based on gorm-gen

## 1. Commands

```shell
Usage of ./gen-table:
  -createTimeField string
        auto create time field name
  -daoFile string
        specify dao filename (default "dao")
  -db string
        input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html] (default "mysql")
  -deleteField string
        delete time field name
  -dsn string
        consult[https://gorm.io/docs/connecting_to_the_database.html]
  -isgendao
        generate curd func dao
  -outPath string
        specify a directory for output (default "./dal")
  -tables string
        enter the required data table or leave it blank
  -updateTimeField string
        auto update time field name

```

## 2. Example

```shell
./gen-table \
	-dsn "example/test.db?_busy_timeout=5000" \
	-updateTimeField "update_at" \
	-createTimeField "create_at" \
	-tables "site" \
	-db "sqlite" \
	-outPath "example/dal" \
	-daoFile "repo" \
	-isgendao
```
