package persistence

import (
	"github.com/Mindyu/blog_system/config"
	"github.com/gomodule/redigo/redis"
	"time"
)

// redis 连接池
var pool *redis.Pool
//根据配置初始化打开redis连接
func init() {
	conf := config.Config().RedisCfg
	pool = &redis.Pool{
		MaxIdle:     20,        // 最大空闲连接数
		MaxActive:   30,        // 一个pool所能分配的最大的连接数目
		IdleTimeout: 60 * time.Second,      // 空闲连接超时时间，超过超时时间的空闲连接会被关闭
		Dial: func() (redis.Conn, error) {  // Dial()方法返回一个连接，从在需要创建连接到的时候调用
			c, err := redis.Dial("tcp", conf.Host+":"+conf.Port,
				redis.DialConnectTimeout(10*time.Second),
				redis.DialReadTimeout(10*time.Second),
				redis.DialWriteTimeout(10*time.Second))
			if err != nil {
				return nil, err
			}
			// 线上环境redis配置密码, 则需要加上这句AUTH
			//_,err = c.Do("AUTH","123456")
			return c, err
		},
		//testOnBorrow 向资源池借用连接时是否做连接有效性检测(ping)，无效连接会被移除 默认值 false 业务量很大时候建议设置为false(多一次ping的开销)。
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
func GetRedisPool() *redis.Pool {
	return pool
}

// 获取redis全局实例
func GetR() redis.Conn {
	return pool.Get()
}
