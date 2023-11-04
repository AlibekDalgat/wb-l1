package main

import (
	"fmt"
	"sync"
)

func add(mbappe map[string]int, mu *sync.Mutex, key string, value int) {
	// захват мютекса перед записью
	mu.Lock()
	// разблокировка поссле записи
	defer mu.Unlock()

	mbappe[key] = value
}

func main() {
	myMap := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	// запуск несколько горутин которые одновременно заполняютмапу в одном и том же ключе
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			add(myMap, &mu, "key", i)
		}()
	}

	// ожидание завершения горутин
	wg.Wait()

	// вывод содержания карты
	for key, value := range myMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
