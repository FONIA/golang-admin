package initial

func init() {
	//注册mysql服务
	InitSql()
	//注册缓存服务
	InitRedis()
	//配置日志
	InitLog()
	//注册错误异常
	InitErr()
}
