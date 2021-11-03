package lib 

import (
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/core/config"
	"os"
	"errors"
	"strings"
)

//客户端IP
func getClientIP(ctx *context.Context) string {
	ip := ctx.Request.Header.Get("X-Forwarded-For")
	if strings.Contains(ip, "127.0.0.1") || ip == "" {
		ip = ctx.Request.Header.Get("X-real-ip")
	}
  
	if ip == "" {
		return "127.0.0.1"
	}
  
	return ip
}

//检测文件是否存在
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//创建文件夹
func CreateDir(path string) error {
	if !FileExists(path){
		return os.MkdirAll(path, 0777)
	}
	return nil
}

//读json
func GetINI(filename string, field string, t string) (interface{}, error) {
	path, _ := os.Getwd()
	appConfigPath := path+"/conf/"+filename
	if FileExists(appConfigPath){
		cfg, _ := config.NewConfig("ini", appConfigPath)
		switch t {
			case "int":
				return cfg.Int(field)
			case "string":
				return cfg.String(field)
		}
	}
	return nil,errors.New(field+" not exists")
}
