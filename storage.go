package go_redis_crud

// Storage base abstraction to implement any other in-memory DB CRUD integration
type Storage interface {
	Create(keyValue KeyValue) (string, error)
	Read(key interface{}) (interface{}, error)
	Update(keyValue KeyValue) (string, error)
	Delete(key interface{}) (string, error)
}

// KeyValue a struct to interact with redis
type KeyValue struct {
	Key   interface{}
	Value interface{}
}
