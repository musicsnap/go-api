package storage

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"go-api/config"
	"strconv"
	"time"
)

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisConn *redis.Pool

func InitRedis() {
	conf := config.GetConfig()
	redisMap := conf.GetStringMap("redis")

	RedisConn = &redis.Pool{
		MaxIdle:     3,
		MaxActive:   10,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			addr := redisMap["host"].(string) + ":" + strconv.Itoa(redisMap["port"].(int))
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if redisMap["password"].(string) != "" {
				if _, err := c.Do("AUTH", redisMap["password"]); err != nil {
					c.Close()
					return nil, err
				}
			}
			database := strconv.Itoa(redisMap["database"].(int))
			if database != "" {
				if _, err := c.Do("SELECT", database); err != nil {
					c.Close()
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
}

func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

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
