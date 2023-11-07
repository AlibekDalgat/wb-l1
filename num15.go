package main

import (
	"fmt"
	"math/rand"
)

func createHugeString(size int) string {
	alf := "qwertyuiopasdfghjklzxcvbnm123456789"
	buffer := make([]byte, size)
	for i := 0; i < size; i++ {
		buffer[i] = alf[rand.Intn(35)]
	}
	return string(buffer)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// создадим слайс рун куда скопируем первые сто элементов огромной строки
	// тогда justString перестанет ссылаться на содержимое перенной v и память v освободиться
	chars := make([]rune, 100)
	copy(chars, []rune(v[:100]))
	justString = string(chars)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
