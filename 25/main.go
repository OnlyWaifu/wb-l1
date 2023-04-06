package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sleep(time.Second * 3)
	fmt.Println(time.Since(start))
}

func sleep(n time.Duration) {
	<-time.After(n)
}
