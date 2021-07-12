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

func (mm *MockMapper) Create(kv go_redis_crud.KeyValue) (string, error) {
	args := mm.Called(kv)

	return args.String(0), args.Error(1)
}

func (mm *MockMapper) Read(k interface{}) (interface{}, error) {
	args := mm.Called(k)

	return args.Get(0), args.Error(1)
}

func (mm *MockMapper) Update(kv go_redis_crud.KeyValue) (string, error) {
	args := mm.Called(kv)

	return args.String(0), args.Error(1)
}

func (mm *MockMapper) Delete(k interface{}) (string, error) {
	args := mm.Called(k)

	return args.String(0), args.Error(1)
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
	mm.On("Create", testKeyValue).Return("created", nil)
	res, err := mm.Create(testKeyValue)
	mm.AssertExpectations(t)
	assert.Equal(t, "created", res)
	assert.Nil(t, err)
}

func TestMapper_CreateErr(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Create", testKeyValue).Return("", errors.New("test error"))
	res, err := mm.Create(testKeyValue)
	assert.Empty(t, res)
	assert.NotNil(t, err)
	mm.AssertExpectations(t)
}

func TestMapper_Read(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Read", expectedKeyOrValue).Return(expectedKeyOrValue, nil)
	res, err := mm.Read(expectedKeyOrValue)
	assert.Equal(t, res, expectedKeyOrValue)
	assert.Nil(t, err)
	mm.AssertExpectations(t)
}

func TestMapper_ReadFail(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Read", expectedKeyOrValue).Return("", errors.New("not found"))
	res, err := mm.Read(expectedKeyOrValue)
	assert.Empty(t, res)
	assert.NotNil(t, err)
	mm.AssertExpectations(t)
}

func TestMapper_Update(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Update", testKeyValue).Return("updated", nil)
	res, err := mm.Update(testKeyValue)
	assert.Equal(t, "updated", res)
	assert.Nil(t, err)
	mm.AssertExpectations(t)
}

func TestMapper_UpdateErr(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Update", testKeyValue).Return("", errors.New("test err"))
	res, err := mm.Update(testKeyValue)
	assert.Empty(t, res)
	assert.NotNil(t, err)
	mm.AssertExpectations(t)
}

func TestMapper_Delete(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Delete", testKeyValue).Return("deleted", nil)
	res, err := mm.Delete(testKeyValue)
	assert.Equal(t, "deleted", res)
	assert.Nil(t, err)
	mm.AssertExpectations(t)
}

func TestMapper_DeleteErr(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Delete", testKeyValue).Return("", errors.New("test err"))
	res, err := mm.Delete(testKeyValue)
	assert.Empty(t, res)
	assert.NotNil(t, err)
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
