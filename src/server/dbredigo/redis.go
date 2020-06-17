package dbredigo

import (
	"github.com/gomodule/redigo/redis"
	"github.com/name5566/leaf/log"
	"server/conf"
	"time"
)

var RedisConn *redis.Pool

// Setup Initialize the Redis instance
func Start() {
	RedisConn = &redis.Pool{
		MaxIdle:     conf.Server.RedisMaxIdle,
		MaxActive:   conf.Server.RedisMaxActive,
		IdleTimeout: conf.Server.RedisIdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conf.Server.RedisHost)
			if err != nil {
				return nil, err
			}
			if conf.Server.RedisPassword != "" {
				if _, err := c.Do("AUTH", conf.Server.RedisPassword); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	log.Release("create redis pool finished")
	return
}

// Set a key/value
func Set(key string, value interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	//value, err := json.Marshal(data)
	//if err != nil {
	//	return err
	//}

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Get get a key
func Get(key string) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}

func HGet(key string, field string) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("HGET", key, field))
	if err != nil {
		return "", err
	}

	return reply, nil
}

func HSet(key string, field string, value string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := redis.Int(conn.Do("HSET", key, field, value))

	if err != nil {
		return false, err
	}

	return true, nil
}

// SetNX a key/value
func SetNX(key string, value interface{}, time int) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	//value, err := json.Marshal(data)
	//if err != nil {
	//	return false, err
	//}
	res, err := redis.String(conn.Do("SET", key, value, "EX", time, "NX"))
	if err != nil || res != "OK" {
		return false, err
	}
	return true, nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func GetString(key string) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}

func HMGet(key string, fields ...string) ([]interface{}, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Values(conn.Do("HMGET", redis.Args{}.Add(key).AddFlat(fields)...))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func HMSet(key string, fields map[string]interface{}) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	res, err := redis.String(conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(fields)...))

	if err != nil || res != "OK" {
		return false, err
	}

	return true, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetKeys(key string) []string {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", key+"*"))
	if err != nil {
		return []string{}
	}

	return keys
}

// Incr a key
func Incr(key string) (int, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	incr, err := redis.Int(conn.Do("INCR", key))
	if err != nil {
		return incr, err
	}
	if incr == 0 {
		incr, err = redis.Int(conn.Do("INCR", key))
		if err != nil {
			return incr, err
		}
	}
	return incr, nil
}

func SetBit(key string, offset uint, value int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := conn.Do("SETBIT", key, offset, value)
	if err != nil {
		return err
	}

	return nil
}

func GetBit(key string, offset uint) (int, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("GETBIT", key, offset))
	if err != nil {
		return 0, err
	}

	return reply, nil
}

// Delete delete a kye
func GetTTL(key string) (int, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Int(conn.Do("TTL", key))
}

//
func LPush(key string, value string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := redis.Strings(conn.Do("LPUSH", key, value))
	if err != nil {
		return err
	}

	return nil
}

//获取列表指定范围内的元素
func LRange(key string, start, stop int) ([]string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	res, err := redis.Strings(conn.Do("LRANGE", key, start, stop))
	if err != nil {
		return nil, err
	}

	return res, nil
}