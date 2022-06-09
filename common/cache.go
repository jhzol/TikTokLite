package common

import (
	"TikTokLite/log"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/garyburd/redigo/redis"
)

var (
	redisClient  *redis.Pool
	ErrMissCache = errors.New("miss Cache")
)

func RedisInit() {
	network := viper.GetString("redis.network")
	address := viper.GetString("redis.address")
	port := viper.GetString("redis.port")
	//auth := viper.GetString("redis.auth")
	host := fmt.Sprintf("%s:%s", address, port)
	redisClient = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, host,
				//redis.DialPassword(auth),
				redis.DialDatabase(2),
			)
			if err != nil {
				log.Error("conn redis failed,", err)
				return nil, err
			}

			return c, err
		},
	}
	redisClient.Get().Do("flushdb")
	log.Info("redis conn success")

}

func GetRedis() redis.Conn {
	return redisClient.Get()
}
func CloseRedis() {
	redisClient.Close()
}

/////////////////////////String类型接口////////////////////////////////////////

func CacheSet(key string, data interface{}, time int) error {
	conn := redisClient.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value, "EX", time)
	if err != nil {
		return err
	}
	return nil
}

func CacheGet(key string) ([]byte, error) {
	conn := redisClient.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	if len(reply) == 0 {
		return nil, ErrMissCache
	}

	return reply, nil
}

///////////////////////////List类型接口////////////////////////////////////////

func CacheLPush(key string, value interface{}) error {
	return listPush("LPUSH", key, value)
}

func CacheRPush(key string, value interface{}) error {
	return listPush("RPUSH", key, value)
}

func CacheLPop(key string) ([]byte, error) {
	return listPop("LPOP", key)
}

func CacheRPop(key string) ([]byte, error) {
	return listPop("RPOP", key)
}

func CacheLGetAll(key string) ([][]byte, error) {
	conn := redisClient.Get()
	defer conn.Close()

	data, err := redis.ByteSlices(conn.Do("LRANGE", key, "0", "-1"))
	if err != nil {
		return [][]byte{}, err
	}
	return data, nil
}

func listPush(op, key string, data interface{}) error {
	conn := redisClient.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do(op, key, value)
	if err != nil {
		return err
	}
	return nil
}

func listPop(op, key string) ([]byte, error) {
	conn := redisClient.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do(op, key))
	if err != nil {
		return reply, err
	}

	return reply, nil
}

/////////////////////////Hash类型接口///////////////////////////////////////////

func CacheHSet(key, mkey string, value ...interface{}) error {
	conn := redisClient.Get()
	defer conn.Close()

	for _, d := range value {
		data, err := json.Marshal(d)
		if err != nil {
			return nil
		}

		//data, err := json.Marshal(value)
		if err != nil {
			return nil
		}
		_, err = conn.Do("HSET", key, mkey, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func CacheHGet(key, mkey string) ([]byte, error) {
	conn := redisClient.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("HGET", key, mkey))

	//fmt.Printf("data:%v", data)
	if err != nil {
		return []byte{}, err
	}
	if len(data) == 0 {
		return []byte{}, ErrMissCache
	}
	return data, nil
}

func CacheDelHash(key, mkey string) error {

	conn := redisClient.Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", key, mkey)
	if err != nil {
		return err
	}
	return nil
}

func CacheDelHash2(key, mkey, comment_id string) error {

	conn := redisClient.Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", key, mkey, comment_id)
	if err != nil {
		return err
	}
	return nil
}
