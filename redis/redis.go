package redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"go_redis_crud"
	"sync"
)

type mapper struct {
	RDBMutex *sync.Mutex
	RDB      *redis.Client
}

func New(addr string, pwd string, db int) *mapper {
	return &mapper{
		RDB:      go_redis_crud.Client(addr, pwd, db),
		RDBMutex: new(sync.Mutex),
	}
}

func (mapper *mapper) Create(kv go_redis_crud.KeyValue) error {
	mapper.RDBMutex.Lock()
	defer mapper.RDBMutex.Unlock()

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

func (mapper *mapper) Read(key interface{}) interface{} {
	mapper.RDBMutex.Lock()
	defer mapper.RDBMutex.Unlock()

	k := prepareKey(key)

	b, err := mapper.RDB.Get(k).Bytes()
	go_redis_crud.CheckErr(err)

	var value interface{}
	go_redis_crud.CheckErr(json.Unmarshal(b, &value))

	if nil == err {
		fmt.Println("[read] key:", key, "val:", value)
	}

	return value
}

func (mapper *mapper) Update(kv go_redis_crud.KeyValue) error {
	mapper.RDBMutex.Lock()
	defer mapper.RDBMutex.Unlock()

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

func (mapper *mapper) Delete(key interface{}) error {
	mapper.RDBMutex.Lock()
	defer mapper.RDBMutex.Unlock()

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
