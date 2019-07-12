# automod
生成表struct的工具

## 使用方法
```
  //数据库配置
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
		//Package:"model", //包名，不配置就是默认model
		//Path:"./model", //保存文件目录
		//SavePath:"./",  //gorm.txt的文件目录，常用的sql查询方法
	}
	dbBase.Init(file)
```
## 目的就是节省大家给表写strcut的时间，直接一键生成，哈哈哈哈
