package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}

	input := make(chan int)
	output := make(chan int)

	go func() {
		for _, val := range nums {
			input <- val
		}
		close(input)
	}()

	go func() {
		for val := range input {
			output <- val * 2
		}
		close(output)
	}()

	for val := range output {
		fmt.Printf("%d ", val)
	}
}
