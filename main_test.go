package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	type testCase struct {
		name     string
		input    int
		expected int
	}

	testCases := []testCase{
		{
			name:     "zero number",
			input:    0,
			expected: 0,
		},
		{
			name:     "negative number",
			input:    -1,
			expected: 0,
		},
		{
			name:     "size is not zero",
			input:    10,
			expected: 10,
		},
	}

	for _, tt := range testCases {
		actual := generateRandomElements(tt.input)
		assert.Equal(t, tt.expected, len(actual), tt.name)
	}
}

func TestMaximum(t *testing.T) {
	type testCase struct {
		name     string
		input    []int
		expected int
	}

	testCases := []testCase{
		{
			name:     "zero number",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "positive number",
			input:    []int{1, 10, 9, 8, 4, 3, 2, 7, 6, 5},
			expected: 10,
		},
	}

	for _, tt := range testCases {
		actual := maximum(tt.input)
		assert.Equal(t, tt.expected, actual, tt.name)
	}
}

func TestMaxChunksWhenSizeIsZero(t *testing.T) {
	type testCase struct {
		name     string
		input    []int
		expected int
	}

	data := make([]int, SIZE)

	for i := 0; i < SIZE; i += 1 {
		data[i] = i + 1
	}

	testCases := []testCase{
		{
			name:     "zero number",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "positive number",
			input:    data,
			expected: SIZE,
		},
	}

	for _, tt := range testCases {
		actual := maxChunks(tt.input)
		assert.Equal(t, tt.expected, actual, tt.name)
	}
}
