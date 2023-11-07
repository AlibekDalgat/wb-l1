package main

import (
	"fmt"
	"sync"
)

// структура состоит из переменной которую инкриментируем и функции инкриментации
type Conter struct {
	k int
}

func (c *Conter) incrementor(mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	// пять увеличиваем счётчик
	for i := 0; i < 5; i++ {
		// блокируем мьютекс перед увеличеним и разблокируем после
		mu.Lock()
		c.k++
		mu.Unlock()
	}
}

func main() {
	var c Conter
	var mu sync.Mutex
	var wg sync.WaitGroup
	// запуск трёх горутин которые инкриментриуют один и тот же счётчик
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go c.incrementor(&mu, &wg)
	}
	// ожидание завершения всех горутин
	wg.Wait()
	fmt.Println(c.k)
}
