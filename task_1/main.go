package main

import "fmt"

func sclone(value int) string {
	if value < 0 {
		return "отрицательных компьютеров быть не может!"
	}

	point := value % 10
	hpoint := value % 100

	if hpoint == 11 || hpoint == 12 || hpoint == 13 || hpoint == 14 {
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

	return "Не могу обработать!"
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
