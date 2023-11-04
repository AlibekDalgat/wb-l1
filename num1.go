package main

import "fmt"

// создание структуры с обозначенными полями
type Human struct {
	age int
	fio string
}

// функция структуры
func (h Human) print() {
	fmt.Println(h.fio, h.age)
}

// создание структуры Action с встроенной структурой Human, что позволит пользоваться полями и функцями структуры Human
type Action struct {
	Human
}

// пример работы
func main() {
	action := Action{Human{age: 20, fio: "DalgatovAM"}}
	fmt.Println(action.fio)
	fmt.Println(action.age)
	action.print()
}
