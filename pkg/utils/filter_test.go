package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/utils"
)

func TestFilter_Int(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4}

	result := utils.Filter(input, func(i int) bool {
		return i%2 == 0
	})

	assert.Equal(t, expected, result)
}

func TestFilter_String(t *testing.T) {
	input := []string{"apple", "banana", "cherry"}
	expected := []string{"apple"}

	result := utils.Filter(input, func(s string) bool {
		return len(s) < 6
	})

	assert.Equal(t, expected, result)
}

func TestFilter_EmptySlice(t *testing.T) {
	input := []int{}
	expected := []int{}

	result := utils.Filter(input, func(i int) bool {
		return i%2 == 0
	})

	assert.Equal(t, expected, result)
}
