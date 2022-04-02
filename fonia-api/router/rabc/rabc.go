package rabc

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {

	web.Get("/b",func(ctx *context.Context){
     ctx.Output.Body([]byte("hello world2"))
})

}