package main

import (
	"context"
	"fmt"
	"time"
)

func worker6_2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// горутина остановлена
			return
		default:
			fmt.Println("Работа горутины...")
			// работа горутины
		}
	}
}

func main() {
	ctxCh, cancel := context.WithCancel(context.Background())

	go worker6_2(ctxCh)

	// подача остановки горутины
	time.Sleep(1 * time.Second)
	cancel()

	// завершение программы
}
