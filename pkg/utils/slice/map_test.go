package slice_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/utils/slice"
)

func TestMap_IntToString(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []string{"1", "2", "3", "4", "5"}

	result := slice.Map(input, func(i int) string {
		return strconv.Itoa(i)
	})

	assert.Equal(t, expected, result)
}

func TestMap_StringToLength(t *testing.T) {
	input := []string{"apple", "banana", "cherry"}
	expected := []int{5, 6, 6}

	result := slice.Map(input, func(s string) int {
		return len(s)
	})

	assert.Equal(t, expected, result)
}

func TestMap_EmptySlice(t *testing.T) {
	input := []int{}
	expected := []string{}

	result := slice.Map(input, func(i int) string {
		return strconv.Itoa(i)
	})

	assert.Equal(t, expected, result)
}
