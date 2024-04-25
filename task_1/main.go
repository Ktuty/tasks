package main

import "fmt"

func sclone(value int) string {
	point := value % 10

	if value == 11 || value == 12 || value == 13 || value == 14 {
		return fmt.Sprintf("%d компьютеров", value)
	}
	switch point {
	case 0, 5, 6, 7, 8, 9:
		return fmt.Sprintf("%d компьютеров", value)
	case 1:
		return fmt.Sprintf("%d компьютер", value)
	case 2, 3, 4:
		return fmt.Sprintf("%d компьютера", value)
	}

	return "отрицательных компьютеров быть не может!"
}

func main() {
	var value int
	for value != -1 {
		_, err := fmt.Scan(&value)
		if err != nil {
			fmt.Printf("Некорректное число %v\n", err)
			break
		}
		fmt.Println(sclone(value))
	}
}
