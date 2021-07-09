**SIMPLE GO APP TO INTERACT WITH REDIS**

* to run test prompt `go run test ./...` (-v)
* to implement a new in-memory DB API just implement `storage.Storage`
* to run docker REDIS server `docker-compose up`
* to interact with REDIS use Create Read Update and Delete methods of `redis.Mapper`


some examples

```go
package main

import (
	"fmt"
	"go_redis_crud"
	"go_redis_crud/redis"
)

func main() {
	mapper := redis.New("localhost:6379", "", 0)
	defer func() {
		go_redis_crud.CheckErr(mapper.RDB.Close())
	}()

	bobbaKeyValue := go_redis_crud.KeyValue{
		Key:   struct{ Bobba string }{Bobba: "Fett"},
		Value: struct{ Jedy string }{Jedy: "are dead 😵"},
	}

	go_redis_crud.CheckErr(mapper.Create(bobbaKeyValue))

	jangoKeyValue := go_redis_crud.KeyValue{
		Key:   struct{ Jango string }{Jango: "Fett"},
		Value: struct{ Syth string }{Syth: "POWER UNLIMITED! 🍪"},
	}

	go_redis_crud.CheckErr(mapper.Create(jangoKeyValue))

	bobbaKey := mapper.Read(struct{ Bobba string }{Bobba: "Fett"})
	fmt.Println(bobbaKey)
}

```