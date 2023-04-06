package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []int{2, 4, 6, 8, 10}
	sum := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(len(a))
	for _, val := range a {
		go func(v int) {
			defer wg.Done()
			mu.Lock() // mutex ON
			sum += v * v
			mu.Unlock() // mutex OFF
		}(val)
	}
	wg.Wait()
	fmt.Printf("The sum of squares: %d", sum)
}
