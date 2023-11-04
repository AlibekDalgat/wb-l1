package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)
	thread := make(chan int)

	// горутина с заполнением канала данными
	go func(ch chan<- int) {
		for i := 1; ; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}(thread)

	// горутиан с чтением канала
	go func(ch <-chan int) {
		for data := range ch {
			fmt.Printf("%d ", data)
		}
	}(thread)

	// ожидание n количества секунд перед закрытием канала
	time.Sleep(time.Duration(n) * time.Second)
	close(thread)
}
