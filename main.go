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
	if size <= 0 {
		return []int{}
	}

	s := make([]int, size)

	for i := range s {
		s[i] = rand.Int()
	}

	return s
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	} else if len(data) == 1 {
		return data[0]
	}
	maxNum := 0
	for i := range data {
		if data[i] > maxNum {
			maxNum = data[i]
		}
	}
	return maxNum
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	} else if len(data) == 1 {
		return data[0]
	}

	var wg sync.WaitGroup
	sizeOfSlice := SIZE / CHUNKS
	maxInSlices := make([]int, CHUNKS)

	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		initalIndex := i * sizeOfSlice

		go func(slice []int) {
			defer wg.Done()
			maxNum := 0

			for j := range slice {
				if slice[j] > maxNum {
					maxNum = slice[j]
				}
			}
			maxInSlices[i] = maxNum
		}(data[initalIndex : initalIndex+sizeOfSlice])
	}
	wg.Wait()

	totalMax := 0
	for i := range maxInSlices {
		if maxInSlices[i] > totalMax {
			totalMax = maxInSlices[i]
		}
	}
	return totalMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := (time.Since(start)).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = (time.Since(start)).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
