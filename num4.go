package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// воркер, который читает произвольные данные из канала
func worker(numWorker int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range ch {
		fmt.Printf("воркер %d: %d\n", numWorker, data)
	}
}

func main() {
	var num int
	fmt.Scan(&num)
	mainThread := make(chan int)
	var wg sync.WaitGroup

	// постоянная запись в главный поток
	go func(ch chan<- int) {
		for i := 1; ; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}(mainThread)

	// запуск воркеров
	for i := 1; i <= num; i++ {
		wg.Add(1)
		go worker(i, mainThread, &wg)
	}

	// создание канала для ожидания сигнала прерывания
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	close(mainThread)

	// ожидание завершение всех воркеров
	wg.Wait()
}
