package middle

import (
	"github.com/beego/beego/v2/server/web/context"
	"strings"
	"encoding/json"
	"fonia/lib"
)

//过滤api
func FilterApi(this *context.Context) {
	//过滤UA
	CheckClientUa(this)
	//检测api通行状态
	CheckApiAuthByMem(this)

}

//redis 过滤api 
func CheckApiAuthByMem(this *context.Context) {
	flag := false
	var err string = ""
	fapi := this.Input.Header("FoniaApi")
	url := strings.Replace(this.Input.URL(), "/", ".", -1)
	api := strings.Trim(strings.TrimSpace(url), ".")
	//去除favicon.ico
	if api == "favicon.ico" {
		return 
	}
	//api存在则refuse key：v1.user value：api暂停 不存在正常
	
	msg, ok := lib.Api_Filter.Data[api] 
	if fapi == api && !ok {
		flag = true
	}else{
		err = msg
	}
	
	//拦截
	if !flag {
		jsonData := &lib.Msg{lib.ErrApi, nil}
		if err !="" {
			errmsg := lib.ErrApi
			errmsg.Msg = err
			jsonData = &lib.Msg{errmsg, nil}
		}

		returnJSON, _ := json.Marshal(jsonData)
		//this.ResponseWriter.Write(returnJSON)
		this.Output.Body(returnJSON)
	}
}

//检查UA
func CheckClientUa(this *context.Context){
	UA := this.Input.UserAgent()
	PassUa := lib.Ua()
	flag := false
	for _, v := range PassUa {
		if strings.Contains(UA, v) {
			//UA正常
			flag = true
			break;
		}
	}

	//非法UA出口
	if !flag {
		jsonData := &lib.Msg{lib.ErrUa, nil}

		returnJSON, _ := json.Marshal(jsonData)
		//this.ResponseWriter.Write(returnJSON)
		this.Output.Body(returnJSON)
	}
}