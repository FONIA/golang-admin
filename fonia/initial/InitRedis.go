package initial

import (
    "github.com/go-redis/redis"
    "github.com/beego/beego/v2/server/web"
    "fonia/lib"
    "time"
    "github.com/beego/beego/v2/core/logs"
)

func InitRedis(){
    
    RedisNetwork, _ := web.AppConfig.String("RedisNetwork")
    RedisAddr, _ := web.AppConfig.String("RedisAddr")
    RedisPassword, _ := web.AppConfig.String("RedisPassword")
    RedisDB, _ := web.AppConfig.Int("RedisDB")
    RedisPoolSize, _ := web.AppConfig.Int("RedisPoolSize")
    RedisMinIdleConns, _ := web.AppConfig.Int("RedisMinIdleConns")
    
    RedisDialTimeout, _ := web.AppConfig.Int("RedisDialTimeout")
	RedisReadTimeout, _ := web.AppConfig.Int("RedisReadTimeout")
	RedisWriteTimeout, _ := web.AppConfig.Int("RedisWriteTimeout")
	RedisPoolTimeout, _ := web.AppConfig.Int("RedisPoolTimeout")
	RedisIdleCheckFrequency, _ := web.AppConfig.Int("RedisIdleCheckFrequency")
	RedisIdleTimeout, _ := web.AppConfig.Int("RedisIdleTimeout")
    RedisMaxConnAge, _ := web.AppConfig.Int("RedisMaxConnAge")
    RedisMaxRetries, _ := web.AppConfig.Int("RedisMaxRetries")
	RedisMinRetryBackoff, _ := web.AppConfig.Int("RedisMinRetryBackoff")
	RedisMaxRetryBackoff, _ := web.AppConfig.Int("RedisMaxRetryBackoff")
    
    client := redis.NewClient(&redis.Options{
		Network:        RedisNetwork,                  
		Addr:           RedisAddr, 
		Password:       RedisPassword,
		DB:             RedisDB,
		PoolSize:       RedisPoolSize,
		MinIdleConns:   RedisMinIdleConns,
 
		DialTimeout:    time.Duration(RedisDialTimeout) * time.Second,
		ReadTimeout:    time.Duration(RedisReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(RedisWriteTimeout) * time.Second,
		PoolTimeout:    time.Duration(RedisPoolTimeout) * time.Second,
		IdleCheckFrequency: time.Duration(RedisIdleCheckFrequency) * time.Second,
		IdleTimeout:    time.Duration(RedisIdleTimeout) * time.Minute,
		MaxConnAge:     time.Duration(RedisMaxConnAge) * time.Second,
		MaxRetries:     RedisMaxRetries,
		MinRetryBackoff: time.Duration(RedisMinRetryBackoff) * time.Millisecond,
		MaxRetryBackoff: time.Duration(RedisMaxRetryBackoff) * time.Millisecond,
    })

    _, err := client.Ping().Result()
	if err != nil {
        logs.Error("redis init error")
	} else {
        logs.Info("redis init success")
		lib.REDIS = client
		apilist, err := client.HGetAll(lib.REDIS_API_FILTER).Result()
		if err == nil {
			lib.Api_Filter.Data = apilist
		}
	}

}

