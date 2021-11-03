//api返回结果
package lib

import (
	"github.com/beego/beego/v2/server/web"
	"encoding/json"
	"io"
	"compress/gzip"
	"strings"
)

type ApiBase struct {
	web.Controller
}

func (this *ApiBase) Reponse(httpStatus int, data string) {
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	if strings.Contains(strings.ToLower(this.Ctx.Request.Header.Get("Accept-Encoding")), "gzip") { //gzip压缩
		this.Ctx.ResponseWriter.Header().Set("Content-Encoding", "gzip")
		this.Ctx.ResponseWriter.WriteHeader(httpStatus)
		w := gzip.NewWriter(this.Ctx.ResponseWriter)
		defer w.Close()
		var bt []byte = []byte(data)
		w.Write(bt)
		w.Flush()
	} else {
		io.WriteString(this.Ctx.ResponseWriter, data)
	}
	this.StopRun()
}


//成功返回
func (this *ApiBase) Success(httpStatus int, data interface{}) {
	reps, _ := json.Marshal(data)
	this.Reponse(httpStatus, string(reps))
}

func (this *ApiBase) Error(httpStatus int, err ErrCode) {
	
	reps, _ := json.Marshal(err)
	this.Reponse(httpStatus, string(reps))
}