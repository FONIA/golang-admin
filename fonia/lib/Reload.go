/*
 * @Description: 全局参数被动更新
 * @Author: FONIA
 * @Email: 3069193298@qq.com
 * @Date: 2021-06-30 11:31:48
 * @LastEditTime: 2021-06-30 11:39:56
 */

package lib

func RELOAD_REDIS_API_FILTER(){
	apilist, err := REDIS.HGetAll(REDIS_API_FILTER).Result()
	if err == nil {
		Api_Filter.Data = apilist
	}
}

