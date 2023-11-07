package main

import "fmt"

func main() {
	arr := []int{3, 3, 1, 5, 6, 8, 9, 4, 7}
	res := quicksort(arr)
	fmt.Println(res)
}

func quicksort(arr []int) []int {
	// завершаем сортировку когда
	if len(arr) < 2 {
		return arr
	}
	// выборопорного элемента
	pivot := arr[0]

	var less, greater []int
	// проходимся по данному срезу и добавляем меньшие опорного элемента элементы в срез меньших чисел less и
	// добавляем большие опорного элемента элементы в срез больших чисел greater
	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	// сортируем срез меньших и больших срезов и соединяем меньший срез, опорный элемент и больший срез
	res := append(quicksort(less), pivot)
	res = append(res, quicksort(greater)...)
	return res
}
