package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	ctx, cancelFunc := context.WithCancel(context.Background())
	// Запуск горутины, которая останавливается при закрытии контекста
	go func() {
		defer wg.Done()
		<-ctx.Done() // Остановка горутины
		fmt.Println("goroutine stopped")
	}()
	cancelFunc()

	ch := make(chan struct{}) // init channel
	// Запуск горутины, которая останавливается, когда канал получает значение
	go func() {
		defer wg.Done()
		<-ch // Остановка горутины
		fmt.Println("goroutine stopped")
	}()
	ch <- struct{}{} // Отправка значения в канал
	// Запуск горутины, которая останавливается после получения значения канала через 1 секунду
	go func() {
		defer wg.Done()
		<-time.After(time.Second * 1)
		fmt.Println("goroutine stopped")
	}()
	wg.Wait() // Ожидание, пока горутины закнчат работу
}
