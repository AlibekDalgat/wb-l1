package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func reverseString(input string) string {
	// создаём переменную типа strings.Builder для удобной запииси рун
	output := strings.Builder{}
	for len(input) > 0 {
		// декодируюм последнюю руну и записываем в билдер
		r, size := utf8.DecodeLastRuneInString(input)
		output.WriteRune(r)
		input = input[:len(input)-size]
	}
	// возвращаем содержимое билдера
	return output.String()
}

func reverseString2(input string) string {
	// вариант с использованием слайса рун
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	input := "главрыба"
	res := reverseString(input)
	fmt.Println(res)

	res = reverseString2(input)
	fmt.Println(res)

}
