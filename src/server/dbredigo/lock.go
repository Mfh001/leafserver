package dbredigo

import (
	"fmt"
	"github.com/name5566/leaf/log"
	"server/util"
	"time"
)

type Lock struct {
	Key     string
	Token   string
	Timeout int
}

func (lock *Lock) GetKey() string {
	return fmt.Sprintf("redislock:%s", lock.Key)
}

func (lock *Lock) GenerateToken() string {
	lock.Token = fmt.Sprintf("%s%x", util.RandomString(8), time.Now().UnixNano())
	return lock.Token
}

func (lock *Lock) GetToken() string {
	return lock.Token
}

func (lock *Lock) Lock() (bool, error) {
	return SetNX(lock.GetKey(), lock.GenerateToken(), lock.Timeout)
}

//func (lock *Lock) AddTimeout(ex_time int64) (ok bool, err error) {
//	ttl_time, err := redis.Int64(lock.conn.Do("TTL", lock.key()))
//	if err != nil {
//		log.Fatal("redis get failed:", err)
//	}
//	if ttl_time > 0 {
//		_, err := redis.String(lock.conn.Do("SET", lock.key(), lock.token, "EX", int(ttl_time+ex_time)))
//		if err == redis.ErrNil {
//			return false, nil
//		}
//		if err != nil {
//			return false, err
//		}
//	}
//	return false, nil
//}

func TryLock(key string, timeout int) (*Lock, bool, error) {
	lock := &Lock{
		Key:     key,
		Timeout: timeout,
	}
	ok, err := lock.Lock()
	return lock, ok, err
}

func (lock *Lock) UnLock() bool {
	val, err := GetString(lock.GetKey())
	if err != nil {
		log.Error("UnLock:" + lock.GetKey())
		return false
	}
	if val != lock.GetToken() {
		log.Error("UnLock:" + lock.GetKey())
		return false
	}
	ok, err := Delete(lock.GetKey())
	if err != nil {
		log.Error("UnLock:" + lock.GetKey())
		return false
	}
	return ok
}
