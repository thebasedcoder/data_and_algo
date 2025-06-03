package main

import "fmt"

func linear_search(arr []int, val int) bool {
	for k := range arr {
		if arr[k] == val {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{4, 2, 5, 8, 6, 56, 53, 58, 54, 1, 35, 78, 12}

	ok := linear_search(arr, 999)
	fmt.Println(ok)
}
