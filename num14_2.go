package main

import (
	"fmt"
	"reflect"
)

func main() {
	var val interface{}

	// использование пакета reflect чтобы получить тип переменной
	val = 1
	fmt.Println(reflect.TypeOf(val).String())

	val = "string"
	fmt.Println(reflect.TypeOf(val).String())

	val = true
	fmt.Println(reflect.TypeOf(val).String())

	val = make(chan int)
	fmt.Println(reflect.TypeOf(val).String())
}
