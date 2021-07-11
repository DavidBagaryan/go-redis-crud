package redis

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go_redis_crud"
	"testing"
)

type MockMapper struct {
	mock.Mock
}

func (mm *MockMapper) Create(kv go_redis_crud.KeyValue) error {
	args := mm.Called(kv)

	return args.Error(0)
}

func (mm *MockMapper) Read(k interface{}) (interface{}, error) {
	args := mm.Called(k)

	return args.Get(0), nil
}

func (mm *MockMapper) Update(kv go_redis_crud.KeyValue) error {
	args := mm.Called(kv)

	return args.Error(0)
}

func (mm *MockMapper) Delete(k interface{}) error {
	args := mm.Called(k)

	return args.Error(0)
}

var testKeyValue = go_redis_crud.KeyValue{Key: "test key 123", Value: "test value 654"}
var expectedKeyOrValue = struct {
	Super interface{}
}{
	Super: struct {
		Val int
	}{Val: 13579},
}

func TestMapper_Create(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Create", testKeyValue).Return(nil)
	assert.Nil(t, mm.Create(testKeyValue))
	mm.AssertExpectations(t)
}

func TestMapper_CreateErr(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Create", testKeyValue).Return(errors.New("test error"))
	assert.NotNil(t, mm.Create(testKeyValue))
	mm.AssertExpectations(t)
}

func TestMapper_Read(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Read", expectedKeyOrValue).Return(expectedKeyOrValue)
	read, err := mm.Read(expectedKeyOrValue)
	assert.Equalf(t, read, expectedKeyOrValue, "not equal on read")
	assert.Nil(t, err)
	mm.AssertExpectations(t)
}

func TestMapper_Update(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Update", testKeyValue).Return(nil)
	assert.Nil(t, mm.Update(testKeyValue))
	mm.AssertExpectations(t)
}

func TestMapper_UpdateErr(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Update", testKeyValue).Return(errors.New("test err"))
	assert.NotNil(t, mm.Update(testKeyValue))
	mm.AssertExpectations(t)
}

func TestMapper_Delete(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Delete", testKeyValue).Return(nil)
	assert.Nil(t, mm.Delete(testKeyValue))
	mm.AssertExpectations(t)
}
func TestMapper_DeleteErr(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Delete", testKeyValue).Return(errors.New("test err"))
	assert.NotNil(t, mm.Delete(testKeyValue))
	mm.AssertExpectations(t)
}

func TestMapper_PrepareKey(t *testing.T) {
	var emptyInterface interface{}

	testTable := []struct {
		TestValue   interface{}
		Expectation string
	}{
		{
			TestValue:   struct{ SupSpecProp interface{} }{SupSpecProp: struct{ SpecProp int }{SpecProp: 9876543210}},
			Expectation: "{\"SupSpecProp\":{\"SpecProp\":9876543210}}",
		},
		{
			TestValue:   emptyInterface,
			Expectation: "null",
		},
		{
			TestValue:   nil,
			Expectation: "null",
		},
		{
			TestValue:   true,
			Expectation: "true",
		},
		{
			TestValue:   "test_value",
			Expectation: "test_value",
		},
		{
			TestValue:   123,
			Expectation: "123",
		},
		{
			TestValue:   123.654,
			Expectation: "123.654",
		},
	}

	for _, testCase := range testTable {
		prepared := prepare(testCase.TestValue)
		assert.IsType(t, "", prepared, "invalid type, expected: string")
		assert.Equal(t, testCase.Expectation, prepared, "invalid prepared key value")
	}
}

func TestRedisClient(t *testing.T) {
	fakeAddr, fakePwd, fakeDB := "fakehost:666", "a_fake_pwd", 9876543210
	rdb := redisClient(fakeAddr, fakePwd, fakeDB)
	opts := rdb.Options()
	assert.Equal(t, opts.Addr, fakeAddr, "wrong host")
	assert.Equal(t, opts.Password, fakePwd, "wrong password")
	assert.Equal(t, opts.DB, fakeDB, "wrong db")
}
