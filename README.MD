**SIMPLE GO APP TO INTERACT WITH REDIS**

* to run test prompt `go run test ./...` (-v)
* to implement a new in-memory DB API just implement `storage.Storage`
* to run docker REDIS server `docker-compose up`
* to interact with REDIS use Create Read Update and Delete methods of `redis.Mapper`
