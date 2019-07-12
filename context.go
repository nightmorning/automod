package automod

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"strings"
)

func NewDbBase(conf *DbConf) *DbBase {

	dbBase := &DbBase{
		dbConf:conf,
	}

	db := dbBase.conn()
	dbBase.db = db

	return dbBase

}

func (dbBase *DbBase) conn() *sql.DB {
	var err error

	dbConf := dbBase.dbConf
	sourceName := dbConf.User + ":" + dbConf.Pwd + "@tcp(" + dbConf.Host + ":" + dbConf.Port + ")/" + dbConf.Database + "?charset=utf8"
	db, err := sql.Open("mysql", sourceName)
	//defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (dbBase *DbBase) Init(file *File)  {
	prefix := dbBase.dbConf.Prefix
	length := len(prefix)
	tables := dbBase.GetTables()
	//fmt.Println(tables)

	fields := make([][]string, 100)
	for _,table := range tables {
		if table == "" {
			continue
		}
		lengths := len(table)

		fields = dbBase.GetFields(table)
		table = table[length:lengths]
		//fmt.Println(table)
		//判断路径是否存在，不存在使用默认路径
		var path string
		list := make(map[string]struct{})
		if _, ok := list[file.Path]; !ok {
			path = "./model"
		}else{
			path = file.Path
		}

		filename := table + ".go"
		var text string
		if _, ok := list[file.Package]; !ok {
			text = "package model\n\n"
		}else{
			text = "package " + file.Package + "\n\n"
		}

		text += "type " + camelString(table) + " struct {\n"
		for _, field := range fields {
			//fileText,err := file.ReadAll("/sql/gorm.go")
			primaryKey := ""

			//fileContent := string(fileText)
			if field == nil {
				break
			}

			if len(field) == 0 {
				break
			}
			fmt.Println(field)
			field[0] = strings.Replace(field[0], "\r\n", "", 0)
			text += "    " + camelString(field[0]) + " " + getGoFieldType(field[1]) + " "
			fmt.Println(field[1])

			if field[3] == "PRI" {
				text += "`gorm:\"primary_key;"
				primaryKey = field[0]
			}

			if field[5] == "auto_increment" {
				text += "AUTO_INCREMENT\" "
			}

			if(primaryKey != ""){
				text += "json:\""+field[0]+"\"`\n"
			}else{
				text += "`json:\""+field[0]+"\"`\n"
			}


		}
		text += "}\n\n"

		var savePath string
		if _, ok := list[file.SavePath]; !ok {
			savePath = "./gorm.txt"
		}else{
			savePath = file.SavePath + "gorm.txt"
		}

		fileText, err := readAll(savePath)

		if err != nil {
			log.Fatal(err)
		}
		content := string(fileText)
		content = strings.Replace(content, "table", table, -1)
		content = strings.Replace(content, "Table", camelString(table), -1)
		text += content
		createModelFile(path, filename, "gorm", text)

	}
}

func (dbBase *DbBase) GetTables() []string {
	db := dbBase.db
	sqlQuery := "show tables"
	rows,err := db.Query(sqlQuery)

	columns, err := rows.Columns()

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	list := make([]string, 100)
	j := 0

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("log:", err)
			panic(err.Error())
		}

		var value string
		for _, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			list[j] = value
		}

		j++

	}

	return list
}

func (dbBase *DbBase) GetFields(table string) [][]string {
	db := dbBase.db
	list := make([][]string, 100)
	if table == "" {
		return list
	}
	querySql := "desc "+table
	rows,err := db.Query(querySql)

	columns, err := rows.Columns()

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	j := 0
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("log:", err)
			panic(err.Error())
		}

		var value string
		for i, col := range values {

			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			list[j] = append(list[j], value)
			list[j][i] = value
		}

		j++
	}

	return list
}

func getFieldType(filedType string) string {

	filedTypeSlice := strings.Split(filedType, "(")

	ok := strings.Contains(filedType, "unsigned")

	fmt.Println(ok)

	if len(filedTypeSlice) > 0 {

		if ok {
			filedType = "u" + filedTypeSlice[0]
		} else {
			filedType = filedTypeSlice[0]
		}

	}
	fmt.Println()

	return filedType
}

func getGoFieldType(filedType string) string {

	filedType = getFieldType(filedType)

	var res string
	
	switch filedType {
	
	case "int":
		res = "int"
		break

	case "uint":
		res = "uint"
		break

	case "tinyint":
		res = "int8"
		break

	case "utinyint":
		res = "uint8"
		break

	case "smallint":
		res = "int16"
		break

	case "usmallint":
		res = "uint16"
		break

	case "bigint":
		res = "int64"
		break

	case "ubigint":
		res = "uint64"
		break

	case "varchar":
		res = "string"
		break

	case "char":
		res = "string"
		break

	case "text":
		res = "string"
		break

	case "longtext":
		res = "string"
		break

	case "datetime":
		res = "time.Time"
		break

	case "decimal":
		res = "float32"
		break

	case "bool":
		res = "bool"
		break

	default:
		res = "int"
		break
	}

	return res
}
