package storage

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"os"
	"strconv"
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
	Key   interface{}
	Value interface{}
}

func Client() *redis.Client {
	addr := getEnv("addr", fmt.Sprint(HOST, ":", PORT))
	pwd := getEnv("pwd", PWD)
	db := getEnv("db", strconv.Itoa(0))
	dbInt, _ := strconv.Atoi(db)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       dbInt,
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
