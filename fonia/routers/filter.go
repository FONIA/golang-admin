package routers

import (
	"github.com/beego/beego/v2/server/web"
	"fonia/middle"
)

func init(){
	//api控制
	web.InsertFilter("*", web.BeforeRouter, middle.FilterApi)



}
