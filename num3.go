package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	res := 0
	
	var wg sync.WaitGroup
	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			res += val * val
		}(val)
	}
	wg.Wait()
	fmt.Println(res)
}
