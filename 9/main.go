package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	a := [10]int{}

	// Отправка целых чисел в канал ch1
	go func() {
		for i := range a {
			ch1 <- i
		}
		close(ch1)
	}()

	// Отправка целых чисел в канал ch2, их удвоение
	go func() {
		for val := range ch1 {
			ch2 <- val * 2
		}
		close(ch2)
	}()

	for val := range ch2 {
		fmt.Println(val)
	}
}
