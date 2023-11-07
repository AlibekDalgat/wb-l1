package main

import (
	"fmt"
	"sync"
)

func addToGroup(a, b float64, nums []float64, mu *sync.Mutex, wg *sync.WaitGroup, groups map[float64][]float64) {
	defer wg.Done()
	// прохождение по всем числам массива для просмотра того, подходит ли оно к данной группе
	for _, val := range nums {
		if val >= a && val < b {
			// блокировка и разблокировка памяти для добавления температуры в группу
			// блокировка нужна из-за того что наж маппой групп работают несколько горутин
			mu.Lock()
			groups[a] = append(groups[a], val)
			mu.Unlock()
		}
	}
}

func main() {
	nums := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[float64][]float64)
	var mu sync.Mutex
	var wg sync.WaitGroup
	// прохождение по каждой возможной группе в отдельной горутине
	for borderStart := -30.; borderStart < 40; borderStart += 10 {
		wg.Add(1)
		go addToGroup(borderStart, borderStart+10, nums, &mu, &wg, groups)
	}

	// ожидание завершения горутин
	wg.Wait()
	fmt.Println(groups)
}
