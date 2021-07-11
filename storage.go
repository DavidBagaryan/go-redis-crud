package go_redis_crud

import (
	"log"
)

// Storage base abstraction to implement any other in-memory DB CRUD integration
type Storage interface {
	Create(keyValue KeyValue) error
	Read(key interface{}) (interface{}, error)
	Update(keyValue KeyValue) error
	Delete(key interface{}) error
}

// KeyValue a struct to interact with redis
type KeyValue struct {
	Key   interface{}
	Value interface{}
}

// CheckErr func prints an err to log output
func CheckErr(err error) {
	if nil != err {
		log.Print(err)
	}
}
