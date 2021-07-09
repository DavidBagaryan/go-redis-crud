package storage

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

const HOST = "localhost"
const PORT = 6379
const PWD = ""
const DB = 0

type Storage interface {
	Create(keyValue KeyValue) error
	Read(key interface{}) *interface{}
	Update(keyValue KeyValue) error
	Delete(key interface{}) error
}

type KeyValue struct {
	Key interface{}
	Value interface{}
}

func Client() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprint(HOST, ":", PORT),
		Password: PWD,
		DB:       DB,
	})

	return rdb
}

func CheckErr(err error) {
	if nil != err {
		log.Print(err)
	}
}
