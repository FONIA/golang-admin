package lib

// import (
// 	"sync"
// )

//出口结构
type Msg struct {
	ErrCode
	Data interface{} `json:"data,omitempty"`
}

// 定义错误码 0~65536
type ErrCode struct {
    Code    uint16 `json:"code"`
    Msg 	string `json:"msg"`
}

// api放行状态缓存
type ApiFilter struct {
    Data  map[string]string
}