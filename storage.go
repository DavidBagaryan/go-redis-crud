package go_redis_crud

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
