package main

import "fmt"

func remove(nums []int, ind int) []int {
	// проверка на out of range
	if ind < len(nums) {
		// возвращаем слайс, соединённый из слайса до индекса удаления и из слайса после индекса удаления
		return append(nums[:ind], nums[ind+1:]...)
	} else {
		return nil
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	nums = remove(nums, 2)

	fmt.Println(nums)
}
