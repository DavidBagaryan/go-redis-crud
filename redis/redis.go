package redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"go_redis_crud"
)

type Mapper struct {
	RDB *redis.Client
}

func (mapper *Mapper) Create(kv go_redis_crud.KeyValue) error {
	k := prepareKey(kv.Key)

	value, err := json.Marshal(kv.Value)
	go_redis_crud.CheckErr(err)

	_, err = mapper.RDB.Set(k, value, 0).Result()
	go_redis_crud.CheckErr(err)

	if nil == err {
		fmt.Println("[add] k:", k, "val:", string(value))
	}

	return err
}

func (mapper *Mapper) Read(key interface{}) *interface{} {
	k := prepareKey(key)

	b, err := mapper.RDB.Get(k).Bytes()
	go_redis_crud.CheckErr(err)

	var value interface{}
	err = json.Unmarshal(b, &value)

	if nil == err {
		fmt.Println("[read] key:", key, "val:", value)
	}

	return &value
}

func (mapper *Mapper) Update(kv go_redis_crud.KeyValue) error {
	k := prepareKey(kv.Key)

	value, err := json.Marshal(kv.Value)
	go_redis_crud.CheckErr(err)

	_, err = mapper.RDB.Set(k, value, 0).Result()
	go_redis_crud.CheckErr(err)

	if nil == err {
		fmt.Println("[update] k:", k, "val:", string(value))
	}

	return err
}

func (mapper *Mapper) Delete(key interface{}) error {
	k := prepareKey(key)

	_, err := mapper.RDB.Del(k).Result()
	go_redis_crud.CheckErr(err)

	if nil == err {
		fmt.Println("[delete] key:", k)
	}

	return err
}

func prepareKey(k interface{}) string {
	key, err := json.Marshal(k)
	go_redis_crud.CheckErr(err)

	return string(key)
}
