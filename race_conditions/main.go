package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu sync.Mutex
)

func main() {
	numbers := []int{}
	// Two goroutines accessing the same shared resource -> numbers slice
	go appendOne(&numbers)
	go appendTwo(&numbers)
	time.Sleep(30 * time.Millisecond)
}

func appendTwo(numbers *[]int) {
	// using mutex to lock the resource until action on it is completed
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("numbers %v\n", numbers)
	*numbers = append(*numbers, []int{1, 2}...)
	fmt.Printf("numbers %v\n", numbers)
}

func appendOne(numbers *[]int) {
	mu.Lock()
	defer mu.Unlock()
	*numbers = append(*numbers, 12)
}
