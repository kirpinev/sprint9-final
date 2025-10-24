package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}

	slice := make([]int, size)

	for i := 0; i < size; i += 1 {
		slice[i] = rand.Int()
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	maxNum := data[0]

	for i := 0; i < len(data); i += 1 {
		if data[i] > maxNum {
			maxNum = data[i]
		}
	}

	return maxNum
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	if len(data) < CHUNKS {
		return maximum(data)
	}

	var wg sync.WaitGroup

	maxNums := make([]int, CHUNKS)
	sliceSize := len(data) / CHUNKS

	for currentIndex := 0; currentIndex < CHUNKS; currentIndex += 1 {
		startIndex := currentIndex * sliceSize
		endIndex := startIndex + sliceSize

		if currentIndex == CHUNKS-1 {
			endIndex = len(data)
		}

		wg.Add(1)

		go func(data, maxNums []int, currentIndex int) {
			defer wg.Done()
			maxNums[currentIndex] = maximum(data)
		}(data[startIndex:endIndex], maxNums, currentIndex)
	}

	wg.Wait()

	return maximum(maxNums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	slice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	maxNum := maximum(slice)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxNum, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	maxNum = maxChunks(slice)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxNum, elapsed)
}
