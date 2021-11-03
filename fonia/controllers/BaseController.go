//基础控制父
package controllers

import (
	"github.com/beego/beego/v2/server/web"

)

type BaseController struct {
	web.Controller
	//Member                *models.Member
}
