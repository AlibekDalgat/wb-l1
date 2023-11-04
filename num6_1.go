package main

import (
	"time"
)

func worker6_1(stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			// горутина остановлена
			return
		default:
			// работа горутины
		}
	}
}

func main() {
	stopCh := make(chan struct{})

	// запуск воркера
	go worker6_1(stopCh)

	// подача остановки горутины
	time.Sleep(1 * time.Millisecond)
	stopCh <- struct{}{}

	// завершение программы
}
