package main

import (
	"fmt"
)

func main() {
	// числа которые больше 2^20
	var a float64 = 3 << 20
	var b float64 = 2 << 20

	// Вывод результатов
	fmt.Printf("Умножение: %f\n", a*b)
	fmt.Printf("Деление: %f\n", a/b)
	fmt.Printf("Сложение: %f\n", a+b)
	fmt.Printf("Вычитание: %f\n", a-b)
}
