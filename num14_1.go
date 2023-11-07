package main

import "fmt"

// проверка типа с помощью оператора switch
func whoIs(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("channel int")
	}
}

func main() {
	var val interface{}
	val = 1
	whoIs(val)

	val = "string"
	whoIs(val)

	val = true
	whoIs(val)

	val = make(chan int)
	whoIs(val)
}
