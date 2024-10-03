package main

import (
	"fmt"
)

func selectionSort(arr []int) []int {
	sorted := []int{}

	for len(arr) > 0 {
		lowestInt := arr[0]
		for _, val := range arr {
			if val < lowestInt {
				lowestInt = val
			}
		}
		sorted = append(sorted, lowestInt)

		for i, val := range arr {
			if val == lowestInt {
				arr = append(arr[:i], arr[i+1:]...)
				break
			}
		}
	}

	return sorted
}

func main() {
	fmt.Println(selectionSort([]int{5, 7, 3, 9}))
}
