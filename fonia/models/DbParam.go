package models

import (
	"time"
)
//json标签意义是定义此结构体解析为json或序列化输出json时value字段对应的key值,如不想此字段被解析可将标签设为`json:"-"`	
//orm的column标签意义是将orm查询结果解析到此结构体时每个结构体字段对应的数据表字段名

const (
	PREX = "FoN_"
	Api = PREX+"api"
	
)

type Api struct {
	Id           	 int       `json:"id" 			orm:"column(id)"`
	Path         	 string    `json:"path" 		orm:"column(path)"`
	Url         	 string    `json:"url" 			orm:"column(url)"`
	Status       	 int       `json:"status" 		orm:"column(status)"`
	ErrMsg      	 string    `json:"errmsg" 		orm:"column(errmsg)"`
	CreateTime       time.Time `json:"create_time" 	orm:"column(create_time)"`
	ModifyTime       time.Time `json:"modify_time" 	orm:"column(modify_time)"`
	RM				 int	   `json:"rm" 			orm:"column(rm)"`
}
