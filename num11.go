package main

import "fmt"

func main() {
	// создание неупорядоченных множеств
	a := map[int]struct{}{4: {}, 5: {}, 3: {}, 1: {}, 2: {}}
	b := map[int]struct{}{7: {}, 4: {}, 8: {}, 6: {}, 5: {}}
	// множество пересечения
	res := make(map[int]struct{})

	// прохождение по первому множеству и анализ того что число существует во втором множестве
	for num := range a {
		if _, ok := b[num]; ok {
			res[num] = struct{}{}
		}
	}

	fmt.Println(res)
}
