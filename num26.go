package main

import (
	"fmt"
)

func check(str string) bool {
	// мапа соддержащая встреченные буквы
	syms := make(map[rune]struct{})
	for _, r := range str {
		// если символ уже встречался возвращает false
		if _, ok := syms[r]; ok {
			return false
		}
		// елси нет, то указывает в мапе что такой символ встречался
		syms[r] = struct{}{}
	}
	// если не встречали повторно символ возвращает true
	return true
}

func main() {
	fmt.Println(check("abcd"))
	fmt.Println(check("abCdefAaf"))
	fmt.Println(check("aabcd"))
}
