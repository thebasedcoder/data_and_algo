package main

import (
	"fmt"
)

func binary_search(arr []int, target int) bool {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2
		value := arr[mid]
		if value == target {
			return true
		} else if value > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 56, 78, 456, 459, 489, 45896, 125698}

	ok := binary_search(arr, 11)

	fmt.Println(ok)
}
