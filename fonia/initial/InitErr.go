package initial

import (
	"github.com/beego/beego/v2/server/web"
	. "fonia/lib"
)

func InitErr(){
	//注册自定义错误页面
	web.ErrorController(&ErrController{})
}

type ErrController struct {
	web.Controller
}

//404异常出口
func (c *ErrController) Error404() {

	c.Data["json"] = &Msg{Err404, nil}
	c.ServeJSON()
}

//500异常
func (c *ErrController) Error500() {

	c.Data["json"] = &Msg{Err500, nil}
	c.ServeJSON()
}

//403资源不可用
func (c *ErrController) Error403() {

	c.Data["json"] = &Msg{Err403, nil}
	c.ServeJSON()
}

//401认证
func (c *ErrController) Error401() {

	c.Data["json"] = &Msg{Err401, nil}
	c.ServeJSON()
}

//系统配置出错
func (c *ErrController) ErrorCfg() {

	c.Data["json"] = &Msg{ErrCfg, nil}
	c.ServeJSON()
}