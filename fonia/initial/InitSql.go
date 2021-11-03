package initial

import (
	"github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/client/orm"
    _ "github.com/go-sql-driver/mysql"
	"fmt"
)

func InitSql(){
	user, _ := web.AppConfig.String("mysqluser")
	passwd, _ := web.AppConfig.String("mysqlpass")
	host, _ := web.AppConfig.String("mysqlurls")
	port, err := web.AppConfig.Int("mysqlport")
	dbname, _ := web.AppConfig.String("mysqldb")
	MaxOpenConns, err2 := web.AppConfig.Int("MaxOpenConns")
	MaxIdleConns, err3 := web.AppConfig.Int("MaxIdleConns")
	Timeout, err4 := web.AppConfig.Int("Timeout")
	ReadTimeout, err5 := web.AppConfig.Int("ReadTimeout")
	
	if nil != err {
		port = 3306
	}
	if nil != err2{
		MaxOpenConns = 1000
	}
	if nil != err3{
		MaxIdleConns = 100
	}
	if nil != err4{
		Timeout = 10
	}
	if nil != err5{
		ReadTimeout = 10
	}

	runmode, _ := web.AppConfig.String("RunMode")
	if runmode == "dev" {
		orm.Debug = true
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%ds&readTimeout=%ds&charset=utf8", user, passwd, host, port, dbname, Timeout, ReadTimeout))
	orm.SetMaxIdleConns("default", MaxIdleConns)
	orm.SetMaxOpenConns("default", MaxOpenConns)


}