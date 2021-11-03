package main

import (
	_ "fonia/routers"
	_ "fonia/initial"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {


	//路由前过滤
//	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	//自定义错误
//	beego.ErrorHandler("404", err.page_not_found)

	beego.Run()
}
