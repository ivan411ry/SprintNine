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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = r.Int()
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	m := data[0]
	for _, v := range data[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	n := len(data)
	if n == 0 {
		return 0
	}
	chunkSize := (n + CHUNKS - 1) / CHUNKS
	chunkCount := (n + chunkSize - 1) / chunkSize
	result := make([]int, chunkCount)
	var wg sync.WaitGroup
	wg.Add(chunkCount)
	for i := 0; i < chunkCount; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > n {
			end = n
		}
		if start >= end {
			continue
		}

		part := data[start:end]
		idx := i
		go func(idx int, part []int) {
			defer wg.Done()
			result[idx] = maximum(part)
		}(idx, part)
	}
	wg.Wait()
	return maximum(result)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
