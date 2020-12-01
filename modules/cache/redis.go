package cache

import (
	"time"

	"github.com/gomodule/redigo/redis"
	. "go.xiet16.com/gopmsweb/conf"
)

var RedisClient *redis.Pool

func init() {
	//建立redis连接池
	RedisClient = &redis.Pool{
		//从配置文件获取maxIdle 以及maxActive, 取不到则取默认值
		MaxIdle:     16,                //最初的连接数量
		MaxActive:   0,                 //连接池最大连接数量，不确定可以用0
		IdleTimeout: 300 * time.Second, //连接关闭时间 300秒
		Dial: func() (redis.Conn, error) { //要连接的redis 数据库
			c, err := redis.Dial(Redis["type"], Redis["address"])
			if err != nil {
				return nil, err
			}

			if Redis["auth"] != "" {
				if _, err = c.Do("AUTH", Redis["auth"]); err != nil {
					_ = c.Close()
					return nil, err
				}
			}

			return c, nil
		},
	}
}
