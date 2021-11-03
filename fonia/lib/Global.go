package lib

import (
    "github.com/go-redis/redis"
)

var (
	REDIS  *redis.Client
)

//api 全局动态变量
var Api_Filter ApiFilter

const (
	//api redis键值
	REDIS_API_FILTER = "fonia_api_filter"
	//aes密钥
	AES_KEY = "eZkr4iy7UsMnuzbqIgTE9Q=="
)


