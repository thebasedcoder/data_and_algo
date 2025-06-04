package main

import "fmt"

func bubble_sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}

func main() {
	slice := []int{89, 5, 469, 25, 598, 46, 123, 0, 569875, 1, 136, 465, 2}

	sorted := bubble_sort(slice)

	fmt.Println(sorted)
}
