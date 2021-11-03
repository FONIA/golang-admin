package initial

import (
	"encoding/json"
	"strings"
	"github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/core/logs"
	"fonia/lib"
	"os"
)


type multifile struct {
    Filename string
	Maxlines int 
	Maxsize int 
    Maxdays int
	Separate []string
}


func InitLog(){
	log_path, err := web.AppConfig.String("LogPath")
	if err != nil {
		log_path = "/log/"
	}

	log_type, err2 := web.AppConfig.String("LogType")
	if err2 != nil {
		log_type = "multifile"
	}

	log_name, err3 := web.AppConfig.String("LogName")
	if err3 != nil {
		log_name = "http.log"
	}

	log_line, err4 := web.AppConfig.Int("LogMaxlines")
	if err4 != nil {
		log_line = 100000
	}

	log_size, err5 := web.AppConfig.Int("LogMaxsize")
	if err5 != nil {
		log_size = 1<<25
	}

	log_day, err5 := web.AppConfig.Int("LogMaxdays")
	if err5 != nil {
		log_day = 7
	}
	//异步
	logs.Async()
	//根目录
	rootPath, _ := os.Getwd()

	//创建文件
	//tm := time.Now().Format("2006-01-02")
	filename := rootPath+log_path+log_name
	if !lib.FileExists(rootPath+log_path) {
		lib.CreateDir(rootPath+log_path)
	}

	switch log_type {
		case "multifile":
			log_sep, err6 := web.AppConfig.String("LogSeparate")
			if err6 != nil {
				log_sep = "error"
			}
			sep_arr := strings.Fields(log_sep)
			logcfg := multifile{
				Filename: filename,
				Maxlines: log_line,
				Maxsize: log_size,
				Maxdays: log_day,
				Separate: sep_arr,
			}

			cfg, _ := json.Marshal(logcfg)
			logs.SetLogger(logs.AdapterMultiFile, string(cfg))
		default:
			
	}

	//覆盖全局变量
	//web.BeeLogger.Async() 
}