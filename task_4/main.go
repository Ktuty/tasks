package main

import "fmt"

func mult_table(t int) {
	for i := 1; i <= t; i++ {
		fmt.Print("\t", i)
	}

	for i := 1; i <= t; i++ {
		fmt.Print("\n", i)
		for j := 1; j <= t; j++ {
			fmt.Print("\t", i*j)
		}
	}
}

func main() {
	var value int

	fmt.Print("Введите число: ")
	fmt.Scan(&value)

	mult_table(value)
}
