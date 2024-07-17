package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"time"
	"vue3-bashItem/pkg/settings"
)

var RedisPool *redis.Pool

func Setup() {
	RedisPool = &redis.Pool{
		MaxIdle:     settings.RedisSetting.MaxIdle,     // 最大空闲连接数
		MaxActive:   settings.RedisSetting.MaxActive,   // 一个pool所能分配的最大的连接数目, 当设置成0的时候，该pool连接数没有限制
		IdleTimeout: settings.RedisSetting.IdleTimeout, // 空闲连接超时时间，超过超时时间的空闲连接会被关闭, 设0永不关闭
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", settings.RedisSetting.Host)
			if err != nil {
				return nil, err
			}

			//
			if settings.RedisSetting.Password != "" {
				if _, err = c.Do("AUTH", settings.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err = c.Do("SELECT", settings.RedisSetting.DB); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// Set 无过期时间
func Set(key string, data interface{}) error {
	conn := RedisPool.Get()
	if err := conn.Err(); err != nil {
		return err
	}
	defer conn.Close()
	if _, typeInfo := data.(string); typeInfo != true { //类型不是string 成立
		value, err := json.Marshal(data)
		if err != nil {
			return err
		}
		_, err = conn.Do("SET", key, value)
		return err
	} else {
		_, dataErr := conn.Do("SET", key, data)
		return dataErr
	}

}

// SetEx 有过期时间  redisGo.SetAndEx("aa","a1111","20") //20秒超时 -1是永久存在
func SetAndEx(key string, data interface{}, Ex string) error {
	conn := RedisPool.Get()
	if err := conn.Err(); err != nil {
		return err
	}
	defer conn.Close()
	if _, typeInfo := data.(string); typeInfo != true { //类型不是string 成立
		value, err := json.Marshal(data)
		if err != nil {
			return err
		}
		_, dataErr := conn.Do("SET", key, value, "EX", Ex)
		return dataErr
	} else {
		_, dataErr := conn.Do("SET", key, data, "EX", Ex)
		return dataErr
	}
}

func Get(key string) ([]byte, error) {
	conn := RedisPool.Get()
	if err := conn.Err(); err != nil {
		return nil, err
	}
	defer conn.Close()
	return redis.Bytes(conn.Do("GET", key))
}

func GetInt(key string) (int, error) {
	conn := RedisPool.Get()
	if err := conn.Err(); err != nil {
		return 0, err
	}
	defer conn.Close()
	return redis.Int(conn.Do("GET", key))
}

func GetString(key string) (string, error) {
	conn := RedisPool.Get()
	if err := conn.Err(); err != nil {
		return "", err
	}
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

func Delete(key string) (bool, error) {
	conn := RedisPool.Get()
	if err := conn.Err(); err != nil {
		return false, err
	}
	defer conn.Close()
	return redis.Bool(conn.Do("DEL", key))
}

func Incr(key string) (int, error) {
	conn := RedisPool.Get()
	if err := conn.Err(); err != nil {
		return 0, err
	}
	defer conn.Close()
	return redis.Int(conn.Do("INCR", key))
}

func Expire(key string, expire int) (bool, error) {
	conn := RedisPool.Get()
	if err := conn.Err(); err != nil {
		return false, err
	}
	defer conn.Close()
	return redis.Bool(conn.Do("EXPIRE", key, expire))
}

func GetTTL(key string) (int, error) {
	conn := RedisPool.Get()
	if err := conn.Err(); err != nil {
		return -2, err
	}
	defer conn.Close()
	return redis.Int(conn.Do("TTL", key))
}
