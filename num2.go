package main

import (
	"fmt"
	"time"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	// прохождение по массиву
	for _, val := range arr {
		// вычисление квадрата элемента массива и вывод происходит в отдельной горутине
		go func(val int) {
			fmt.Println(val * val)
		}(val) // подача переменной val, как элемента массива в функцию, работающеей в горутине
	}
	// дожидаемся вывода всех квадратов
	time.Sleep(1 * time.Second)
}
