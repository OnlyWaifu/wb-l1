package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var n int
	ch := make(chan int)
	a := [10000]int{}
	fmt.Print("Введите время ожидания в секундах: ")
	fmt.Scan(&n)

	// Отключение по окончании времени
	time.AfterFunc(time.Duration(n)*time.Second, func() {
		fmt.Println("time out")
		os.Exit(0)
	})

	// Функция для последовательной отправки данных в канал
	go func() {
		for i := range a {
			ch <- i
		}
		close(ch)
	}()

	// Чтение всех значений из канала
	for val := range ch {
		fmt.Printf("value from chan: %d\n", val)
	}
}
