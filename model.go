package automod

import "database/sql"

type DbConf struct {
	Host string
	User string
	Pwd string
	Port string
	Database string
	Prefix string
}

type DbBase struct {
	dbConf *DbConf
	db *sql.DB
	//tables *[]Table
}

type Table struct {
	field string
	fType string
	isnull string
	defValue string
	primaryKey string
	autoInc string
	comment string
}

type File struct {
	Package string
	Path string
	SavePath string
}
