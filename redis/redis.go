package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"go_redis_crud"
	"sync"
)

type mapper struct {
	RDBMutex *sync.RWMutex
	RDB               *redis.Client
}

func New(addr string, pwd string, db int) *mapper {
	return &mapper{
		RDB:      redisClient(addr, pwd, db),
		RDBMutex:      new(sync.RWMutex),
	}
}

// RedisClient func connects to Redis server
func redisClient(addr string, pwd string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})

	return rdb
}

// Create a new element in dict
func (mapper *mapper) Create(kv go_redis_crud.KeyValue) error {
	found, _ := mapper.find(kv.Key)
	if found != nil {
		return errors.New(fmt.Sprint("[create] already exists: ", kv.Key))
	}

	err, key, value := mapper.set(kv)

	if nil == err {
		fmt.Println("[create] k:", key, "val:", value)
	}

	return err
}

// Read element from dict by key
// Possible to read structured key by passing a string sign, for instance `"{\"Key\": \"Val\"}"`
func (mapper *mapper) Read(key interface{}) (interface{}, error) {
	b, err := mapper.find(key)
	if err != nil {
		return struct{}{}, errors.New(fmt.Sprint("[read] not found: ", key))
	}

	var value interface{}
	go_redis_crud.CheckErr(json.Unmarshal(b, &value))

	fmt.Println("[read] key:", key, "val:", value)

	return value, nil
}

// Update element in dict by key
func (mapper *mapper) Update(kv go_redis_crud.KeyValue) error {
	_, err := mapper.find(kv.Key)
	if err != nil {
		return errors.New(fmt.Sprint("[update] not found: ", kv.Key))
	}

	err, key, value := mapper.set(kv)

	if nil == err {
		fmt.Println("[update] k:", key, "val:", value)
	}

	return err
}

// Delete element from dict
func (mapper *mapper) Delete(key interface{}) error {
	k := prepare(key)

	_, err := mapper.find(key)
	if err != nil {
		return errors.New("[delete] not found: " + k)
	}

	mapper.RDBMutex.Lock()
	defer mapper.RDBMutex.Unlock()

	_, err = mapper.RDB.Del(k).Result()
	go_redis_crud.CheckErr(err)

	if nil == err {
		fmt.Println("[delete] key:", k)
	}

	return err
}

// set Uses redis SET function to creates OR updates an element
// set retrieves an err if exists and wrote stringify key/value
func (mapper *mapper) set(kv go_redis_crud.KeyValue) (error, string, string) {
	mapper.RDBMutex.Lock()
	defer mapper.RDBMutex.Unlock()

	k := prepare(kv.Key)

	value, err := json.Marshal(kv.Value)
	go_redis_crud.CheckErr(err)

	_, err = mapper.RDB.Set(k, value, 0).Result()
	go_redis_crud.CheckErr(err)

	return err, k, prepare(kv.Value)
}

// find an element from redis dict
func (mapper *mapper) find(key interface{}) ([]byte, error) {
	mapper.RDBMutex.RLock()
	defer mapper.RDBMutex.RUnlock()

	return mapper.RDB.Get(prepare(key)).Bytes()
}

// prepare Key/Value preparation for CRUD or represent
// special case on string to NOT add extra ""
func prepare(val interface{}) string {
	switch val.(type) {
	case string:
		return fmt.Sprint(val)
	default:
		k, err := json.Marshal(val)
		go_redis_crud.CheckErr(err)

		return string(k)
	}
}
