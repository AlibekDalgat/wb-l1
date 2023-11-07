package main

import "fmt"

func main() {
	a := 3
	b := 7

	// метод с помощью математических операций. можно воспользоваться механизмом языка: a, b = b, a
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}
