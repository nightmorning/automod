package automod

import "testing"

func TestNewDbBase(t *testing.T) {
	dbConf := &DbConf{
		Host:"127.0.0.1",
		User:"root",
		Pwd:"root",
		Port:"3306",
		Database:"lottery",
		Prefix:"ts_",
	}
	dbBase := NewDbBase(dbConf)
	file := &File{
		//Package:"model",
		//Path:"./model",
		//SavePath:"./models",
	}
	dbBase.Init(file)
}
