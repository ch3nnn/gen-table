# gen-table

[English](README_EN.md) | 简体中文

基于 gorm-gen 生成 model、query、dao 代码

## 一、相关命令

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

## 二、示例

注意: 需要在 go.mod 文件同级目录下执行

```shell
./gen-table \
	-dsn "example/test.db?_busy_timeout=5000" \
	-tables "site" \
	-db "sqlite" \
	-outPath "example/dal" \
	-isgendao
```
