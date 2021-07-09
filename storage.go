package go_redis_crud

import (
	"github.com/go-redis/redis"
	"log"
	"os"
)

type Storage interface {
	Create(keyValue KeyValue) error
	Read(key interface{}) *interface{}
	Update(keyValue KeyValue) error
	Delete(key interface{}) error
}

type KeyValue struct {
	Key   interface{}
	Value interface{}
}

func Client(addr string, pwd string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})

	return rdb
}

func CheckErr(err error) {
	if nil != err {
		log.Print(err)
	}
}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}

	return value
}
