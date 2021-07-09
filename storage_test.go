package go_redis_crud

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

const keyWrongType = "key wrong type, expected: "
const valueWrongType = "value wrong type, expected: "

const keyNotEqual = "key not equal"
const valueNotEqual = "value not equal"

func TestModel_Int(t *testing.T) {
	kv := KeyValue{Key: 123, Value: 123}
	expectedType := "int"
	expectedVal := 123

	assert.IsType(t, 0, kv.Key, keyWrongType+expectedType)
	assert.Equalf(t, expectedVal, kv.Key, keyNotEqual)

	assert.IsType(t, 0, kv.Value, valueWrongType+expectedType)
	assert.Equalf(t, expectedVal, kv.Value, valueNotEqual)
}

func TestModel_Float(t *testing.T) {
	kv := KeyValue{Key: 123.321, Value: 123.321}
	expectedType := "float"
	expectedVal := 123.321

	assert.IsType(t, 0.0, kv.Key, keyWrongType+expectedType)
	assert.Equalf(t, expectedVal, kv.Key, keyNotEqual)

	assert.IsType(t, 0.0, kv.Value, valueWrongType+expectedType)
	assert.Equalf(t, expectedVal, kv.Value, valueWrongType)
}

func TestModel_Str(t *testing.T) {
	kv := KeyValue{Key: "test_super_test", Value: "test_super_test"}
	expectedType := "string"
	expectedVal := "test_super_test"

	assert.IsType(t, "", kv.Key, keyWrongType+expectedType)
	assert.Equalf(t, expectedVal, kv.Key, keyNotEqual)

	assert.IsType(t, "", kv.Value, valueWrongType+expectedType)
	assert.Equalf(t, expectedVal, kv.Value, valueNotEqual)
}

func TestModel_Bool(t *testing.T) {
	kv := KeyValue{Key: false, Value: false}
	expectedType := "boolean"

	assert.IsType(t, true, kv.Key, keyWrongType+expectedType)
	assert.Equal(t, false, kv.Key, keyNotEqual)

	assert.IsType(t, true, kv.Value, valueWrongType+expectedType)
	assert.Equalf(t, false, kv.Value, valueNotEqual)
}

func TestModel_Nil(t *testing.T) {
	kv := KeyValue{}
	expectedType := "nil"

	assert.IsType(t, nil, kv.Key, keyWrongType+expectedType)
	assert.Equalf(t, nil, kv.Key, keyNotEqual)

	assert.IsType(t, nil, kv.Value, valueWrongType+expectedType)
	assert.Equalf(t, nil, kv.Value, valueNotEqual)

	kv.Key = nil
	kv.Value = nil

	assert.IsType(t, nil, kv.Key, keyWrongType+expectedType)
	assert.Equalf(t, nil, kv.Key, keyNotEqual)

	assert.IsType(t, nil, kv.Value, valueWrongType+expectedType)
	assert.Equalf(t, nil, kv.Value, valueNotEqual)
}

func TestModel_Struct(t *testing.T) {
	key := struct {
		SuperKeyProp interface{}
	}{SuperKeyProp: struct {
		KeyProp string
	}{KeyProp: "test super"}}

	value := struct {
		ValueProp int
	}{ValueProp: 123}

	kv := KeyValue{Key: key, Value: value}
	expectedType := "struct"

	assert.IsType(t, struct{ SuperKeyProp interface{} }{}, kv.Key, keyWrongType+expectedType)
	assert.Equalf(t, key, kv.Key, keyNotEqual)

	assert.IsType(t, struct{ ValueProp int }{}, kv.Value, valueWrongType+expectedType)
	assert.Equalf(t, value, kv.Value, valueNotEqual)
}

func TestClient(t *testing.T) {
	fakeAddr, fakePwd, fakeDB := "fakehost:666", "a_fake_pwd", 9876543210
	rdb := Client(fakeAddr, fakePwd, fakeDB)
	opts := rdb.Options()
	assert.Equal(t, opts.Addr, fakeAddr, "wrong host")
	assert.Equal(t, opts.Password, fakePwd, "wrong password")
	assert.Equal(t, opts.DB, fakeDB, "wrong db")
}

func TestCheckErr(t *testing.T) {
	CheckErr(errors.New("test err msg")) // todo check stdout
}

func TestGetEnv(t *testing.T) {
	r := getEnv("wrong_key", "fallback")
	assert.Equalf(t, "fallback", r, "internal getEnv retrieves wrong value")
}
