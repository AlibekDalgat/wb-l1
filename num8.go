package main

import (
	"fmt"
)

func setBit(n int64, pos uint, bitValue int) int64 {
	mask := int64(1 << pos) // Создаем маску с установленным битом в позиции pos

	if bitValue == 1 {
		n |= mask // Устанавливаем бит в 1 путем побитового ИЛИ с маской
	} else if bitValue == 0 {
		n &= ^mask // Устанавливаем бит в 0 путем побитового И с инверсией маски
	}

	return n
}

func main() {
	var num int64
	fmt.Scan(&num) // введение числа
	var i uint
	fmt.Scan(&i) // позиция бита (нумерация с нуля)
	var isOne bool
	fmt.Scan(&isOne) // значение, которое нужно установить (1 - true, 0 - false)

	mask := int64(1 << i) // маска с установленным битом в i
	if isOne {
		num |= mask // добавление бита с помощью побитового ИЛИ с маской
	} else {
		num &= ^mask // изъятия бита с помощью побитового И с инверсией маски
	}

	fmt.Println(num)
}
