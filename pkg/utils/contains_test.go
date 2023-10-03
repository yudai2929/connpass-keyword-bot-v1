package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsInt_Success(t *testing.T) {
	slice := []int{1, 2, 3}
	item := 1
	result := Contains(slice, item)
	assert.True(t, result)
}

func TestContainsInt_Fail(t *testing.T) {
	slice := []int{1, 2, 3}
	item := 4
	result := Contains(slice, item)
	assert.False(t, result)
}

func TestContainsString_Success(t *testing.T) {
	slice := []string{"a", "b", "c"}
	item := "a"
	result := Contains(slice, item)
	assert.True(t, result)
}

func TestContainsString_Fail(t *testing.T) {
	slice := []string{"a", "b", "c"}
	item := "d"
	result := Contains(slice, item)
	assert.False(t, result)
}

func TestContainsFloat_Success(t *testing.T) {
	slice := []float64{1.1, 2.2, 3.3}
	item := 1.1
	result := Contains(slice, item)
	assert.True(t, result)
}

func TestContainsFloat_Fail(t *testing.T) {
	slice := []float64{1.1, 2.2, 3.3}
	item := 4.4
	result := Contains(slice, item)
	assert.False(t, result)
}

func TestContainsBool_Success(t *testing.T) {
	slice := []bool{true, false}
	item := true
	result := Contains(slice, item)
	assert.True(t, result)
}

func TestContainsBool_Fail(t *testing.T) {
	slice := []bool{true, false}
	item := true
	result := Contains(slice, item)
	assert.True(t, result)
}

func TestContainsStruct_Success(t *testing.T) {
	type TestStruct struct {
		ID   int
		Name string
	}
	slice := []TestStruct{
		{ID: 1, Name: "a"},
		{ID: 2, Name: "b"},
		{ID: 3, Name: "c"},
	}
	item := TestStruct{ID: 1, Name: "a"}
	result := Contains(slice, item)
	assert.True(t, result)
}

func TestContainsStruct_Fail(t *testing.T) {
	type TestStruct struct {
		ID   int
		Name string
	}
	slice := []TestStruct{
		{ID: 1, Name: "a"},
		{ID: 2, Name: "b"},
		{ID: 3, Name: "c"},
	}
	item := TestStruct{ID: 4, Name: "d"}
	result := Contains(slice, item)
	assert.False(t, result)
}
