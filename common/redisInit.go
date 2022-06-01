package common

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var C redis.Conn

func RedisInit() {
	var err error
	C, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("redis conn success")

}

func GetRE() redis.Conn {
	return C
}
func CloseReBase() {
	C.Close()
}
