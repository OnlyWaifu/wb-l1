package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(len(a))

	for _, val := range a {
		// Запуск горутины с входным параметром int, чтобы избежать использования общей памяти
		go func(v int) {
			defer wg.Done() // Уменьшение счетчика wg.Wait()
			fmt.Println(v * v)
		}(val)
	}

	wg.Wait() // Ожидание, пока горутины закончат свою работу
}
