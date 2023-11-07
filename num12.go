package main

import "fmt"

func main() {
	strSlice := []string{"cat", "cat", "dog", "cat", "tree"}
	// множество куда будут добавляться строки
	set := make(map[string]struct{})

	// прохождение по последовательности и добавление в множество
	for _, str := range strSlice {
		set[str] = struct{}{}
	}

	fmt.Println(set)
}
