package main

import "fmt"

func simple_value(start, end int) []int {
	var result []int

	for ; start <= end; start++ {
		if start%2 != 0 && start%3 != 0 && start%5 != 0 {
			result = append(result, start)
		}
		if start == 2 || start == 3 || start == 5 {
			result = append(result, start)
		}
	}

	return result
}

func main() {
	fmt.Println(simple_value(12, 12))
}
