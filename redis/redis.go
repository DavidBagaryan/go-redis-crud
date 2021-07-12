package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"go_redis_crud"
	"log"
	"sync"
)

type mapper struct {
	rdbMutex *sync.RWMutex
	rdb      *redis.Client
}

func New(addr string, pwd string, db int) *mapper {
	return &mapper{
		rdb:      redisClient(addr, pwd, db),
		rdbMutex: new(sync.RWMutex),
	}
}

func (mapper *mapper) Close() {
	err := mapper.rdb.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// RedisClient func connects to Redis server
func redisClient(addr string, pwd string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
}

// Create a new element in dict
func (mapper *mapper) Create(kv go_redis_crud.KeyValue) (string, error) {
	_, err := mapper.find(kv.Key)
	if nil == err {
		return "", errors.New(fmt.Sprint("[create] already exists: ", prepare(kv.Key)))
	}

	key, value, err := mapper.set(kv)

	if nil == err {
		return fmt.Sprint("[create] k:", prepare(key), "val:", prepare(value)), nil
	}

	return "", err
}

// Read element from dict by key
// Possible to read structured key by passing a string sign, for instance `"{\"Key\": \"Val\"}"`
func (mapper *mapper) Read(key interface{}) (interface{}, error) {
	b, err := mapper.find(key)
	if err != nil {
		return struct{}{}, errors.New(fmt.Sprint("[read] not found: ", prepare(key)))
	}

	var value interface{}
	err = json.Unmarshal(b, &value)
	if nil != err {
		return struct{}{}, err
	}

	return fmt.Sprint("[read] key:", prepare(key), " val:", prepare(value)), nil
}

// Update element in dict by key
func (mapper *mapper) Update(kv go_redis_crud.KeyValue) (string, error) {
	_, err := mapper.find(kv.Key)
	if err != nil {
		return "", errors.New(fmt.Sprint("[update] not found: ", prepare(kv.Key)))
	}

	key, value, err := mapper.set(kv)

	if nil == err {
		return fmt.Sprint("[update] k:", prepare(key), "val:", prepare(value)), nil
	}

	return "", err
}

// Delete element from dict
func (mapper *mapper) Delete(key interface{}) (string, error) {
	k := prepare(key)

	_, err := mapper.find(key)
	if err != nil {
		return "", errors.New("[delete] not found: " + k)
	}

	mapper.rdbMutex.Lock()
	defer mapper.rdbMutex.Unlock()

	_, err = mapper.rdb.Del(k).Result()
	if nil != err {
		return "", err
	}

	return fmt.Sprint("[delete] key:", prepare(k)), nil
}

// set Uses redis SET function to creates OR updates an element
// set retrieves an err if exists and wrote stringify key/value
func (mapper *mapper) set(kv go_redis_crud.KeyValue) (string, string, error) {
	mapper.rdbMutex.Lock()
	defer mapper.rdbMutex.Unlock()

	k := prepare(kv.Key)

	value, err := json.Marshal(kv.Value)
	if nil != err {
		return "", "", err
	}

	_, err = mapper.rdb.Set(k, value, 0).Result()
	if nil != err {
		return "", "", err
	}

	return k, prepare(kv.Value), err
}

// find an element from redis dict
func (mapper *mapper) find(key interface{}) ([]byte, error) {
	mapper.rdbMutex.RLock()
	defer mapper.rdbMutex.RUnlock()

	return mapper.rdb.Get(prepare(key)).Bytes()
}

// prepare Key/Value preparation for CRUD or represent
// special case on string to NOT add extra ""
func prepare(value interface{}) string {
	switch value.(type) {
	case string:
		return fmt.Sprint(value)
	default:
		k, err := json.Marshal(value)
		if nil != err {
			log.Fatalln(err)
		}

		return string(k)
	}
}
