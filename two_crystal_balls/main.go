package main

import (
	"fmt"
	"math"
)

func Two_Crystal_Balls(breaks []bool) int {
	var jumpAmount = int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := 0
	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}
	i -= jumpAmount

	for j := 0; j <= jumpAmount && i < len(breaks); j++ {
		if breaks[i] {
			return i
		}
		i++
	}
	return -1
}

func main() {
	slice := []bool{false, false, true}
	output := Two_Crystal_Balls(slice)
	fmt.Println(output)
}
