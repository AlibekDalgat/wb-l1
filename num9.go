package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}

	input := make(chan int)
	output := make(chan int)

	// ввод чисел x в первый поток
	go func() {
		for _, val := range nums {
			input <- val
		}
		close(input)
	}()

	// чтение чисел x из первого потока и ввод чиел x*2 во второй поток
	go func() {
		for val := range input {
			output <- val * 2
		}
		close(output)
	}()

	// вывод содержимого второго потока
	for val := range output {
		fmt.Printf("%d ", val)
	}
}
