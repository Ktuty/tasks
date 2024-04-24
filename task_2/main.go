package main

import "fmt"

func common_divisor(arr []int) []int {
	var min int
	Div := make([]int, 0)
	comonDiv := make(map[int]int)

	if len(arr) == 0 {
		return Div
	}

	min = arr[0]
	for _, i := range arr[1:] {
		if i < min {
			min = i
		}
	}

	for i := 2; i <= min; i++ {
		if min%i == 0 {
			Div = append(Div, i)
		}
	}

	for _, div := range Div {
		for _, val := range arr {
			if val%div == 0 {
				comonDiv[div]++
			}
		}
	}

	var result []int
	for _, val := range Div {
		if comonDiv[val] == len(arr) {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	arr := []int{12, 24, 48}
	fmt.Println(common_divisor(arr))
}
