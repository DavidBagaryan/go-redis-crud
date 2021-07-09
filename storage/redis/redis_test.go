package redis

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"key-value-storage/storage"
	"testing"
)

type MockMapper struct {
	mock.Mock
}

func (mm *MockMapper) Create(kv storage.KeyValue) error {
	args := mm.Called(kv)
	return args.Error(0)
}

func (mm *MockMapper) Read(k interface{}) *interface{} {
	args := mm.Called(k)
	value := args.Get(0)
	return &value
}

func (mm *MockMapper) Update(kv storage.KeyValue) error {
	args := mm.Called(kv)
	return args.Error(0)
}

func (mm *MockMapper) Delete(k interface{}) error {
	args := mm.Called(k)
	return args.Error(0)
}

var testKeyValue = storage.KeyValue{Key: "test key 123", Value: "test value 654"}
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
	mm.Create(testKeyValue)
	mm.AssertExpectations(t)

	mm.On("Create", testKeyValue).Return(errors.New("test error"))
	mm.Create(testKeyValue)
	mm.AssertExpectations(t)
}

func TestMapper_Read(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Read", expectedKeyOrValue).Return(expectedKeyOrValue)
	mm.Read(expectedKeyOrValue)
	mm.AssertExpectations(t)
}

func TestMapper_Update(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Update", testKeyValue).Return(nil)
	mm.Update(testKeyValue)
	mm.AssertExpectations(t)

	mm.On("Update", testKeyValue).Return(errors.New("test err"))
	mm.Update(testKeyValue)
	mm.AssertExpectations(t)
}

func TestMapper_Delete(t *testing.T) {
	mm := new(MockMapper)
	mm.On("Delete", testKeyValue).Return(nil)
	mm.Delete(testKeyValue)
	mm.AssertExpectations(t)

	mm.On("Delete", testKeyValue).Return(errors.New("test err"))
	mm.Delete(testKeyValue)
	mm.AssertExpectations(t)
}

func TestMapper_PrepareKey(t *testing.T) {
	k := struct {
		SupSpecProp interface{}
	}{SupSpecProp: struct {
		SpecProp int
	}{SpecProp: 9876543210}}

	newKey := prepareKey(k)
	assert.IsType(t, "", newKey, "invalid type, expected: string")
	assert.Equal(t, "{\"SupSpecProp\":{\"SpecProp\":9876543210}}", newKey, "invalid prepared key value")
}
