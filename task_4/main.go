package main

import "fmt"

func mult_table(t int) {
	// Выводим строчку множетелей
	for i := 1; i <= t; i++ {
		fmt.Print("\t", i)
	}
	// Выводим строчки с множетелем и результатами произведения
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
