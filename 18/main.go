package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Counter — это счетчик с защитой от параллелизма
type Counter struct {
	counter int64
}

// Inc увеличивает значение счетчика с защитой от параллелизма
func (c *Counter) Inc() {
	atomic.AddInt64(&c.counter, 1)
}

// Reset устанавливает значение счетчика на 0 с учетом параллелизма
func (c *Counter) Reset() {
	atomic.StoreInt64(&c.counter, 0)
}

func (c Counter) Value() int64 {
	return atomic.LoadInt64(&c.counter)
}

func main() {
	c := Counter{}
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(c.Value())
}
