package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	var n int
	ch := make(chan string)
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&n)
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	runWorkers(ctx, ch, n)

	r := bufio.NewScanner(os.Stdin)
	// Последовательное чтение данных из stdin ввода и отправка их в канала ch
	for {
		select {
		// Если получен os.Interrupt
		case <-sig:
			cancel()                // Отмена воркеров с контекстом
			time.Sleep(time.Second) // Ожидание конца работыв воркеров
			return
		default:
			if r.Scan() {
				ch <- r.Text() // Отправка данных в канал
			}
		}
	}
}

// runWorkers запускает N горутин, которые считывают значения из канала chan и выводят их
func runWorkers(ctx context.Context, ch chan string, n int) {
	for i := 0; i < n; i++ {
		// Запуск горутины
		go func(ind int) {
			// Горутина слушает канал ch и распечатывает из него значения
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("finish worker %d\n", ind)
					return
				case val := <-ch:
					fmt.Printf("value %s from worker %d\n", val, ind)
				}
			}
		}(i)
	}
}
