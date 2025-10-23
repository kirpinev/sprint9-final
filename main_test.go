package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElementsWhenSizeIsZero(t *testing.T) {
	size := 0
	slice := generateRandomElements(size)

	assert.Equal(t, size, len(slice))
}

func TestGenerateRandomElementsWhenSizeIsNegative(t *testing.T) {
	size := -1
	slice := generateRandomElements(size)

	assert.Equal(t, 0, len(slice))
}

func TestGenerateRandomElementsWhenSizeIsNotZero(t *testing.T) {
	size := 10
	slice := generateRandomElements(size)

	assert.Equal(t, size, len(slice))
}

func TestMaximumWhenSizeIsZero(t *testing.T) {
	data := make([]int, 0)
	maxNum := maximum(data)

	assert.Equal(t, 0, maxNum)
}

func TestMaximumWhenSizeIsNotZeroWithPositiveNums(t *testing.T) {
	data := []int{1, 10, 9, 8, 4, 3, 2, 7, 6, 5}
	correctMaxNum := 10
	maxNum := maximum(data)

	assert.Equal(t, correctMaxNum, maxNum)
}

func TestMaximumWhenSizeIsNotZeroWithNegativeNums(t *testing.T) {
	data := []int{-1, -10, -9, -8, -4, -3, -2, -7, -6, -5}
	correctMaxNum := -1
	maxNum := maximum(data)

	assert.Equal(t, correctMaxNum, maxNum)
}

func TestMaxChunksWhenSizeIsZero(t *testing.T) {
	data := make([]int, 0)
	maxNum := maxChunks(data)

	assert.Equal(t, 0, maxNum)
}

func TestMaxChunksWhenSizeIsNotZeroWithPositiveNums(t *testing.T) {
	data := make([]int, SIZE)

	for i := 0; i < SIZE; i += 1 {
		data[i] = i + 1
	}

	correctMaxNum := SIZE
	maxNum := maxChunks(data)

	assert.Equal(t, correctMaxNum, maxNum)
}

func TestMaxChunksWhenSizeIsNotZeroWithNegativeNums(t *testing.T) {
	data := make([]int, SIZE)

	for i := 0; i < SIZE; i += 1 {
		data[i] = -(i + 1)
	}

	correctMaxNum := -1
	maxNum := maxChunks(data)

	assert.Equal(t, correctMaxNum, maxNum)
}
