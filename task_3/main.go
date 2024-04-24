package main

import "fmt"

func simple_value(start, end int) []int {
	var result []int

	// если число не делится на 2, 3 и 5 => это простое число
	for ; start <= end; start++ {
		if start%2 != 0 && start%3 != 0 && start%5 != 0 {
			result = append(result, start)
		}
		// числа 2, 3 и 5 простые, чтобы их не пропустить делаем проверку
		if start == 2 || start == 3 || start == 5 {
			result = append(result, start)
		}
	}

	return result
}

func main() {
	fmt.Println(simple_value(11, 20))
}
