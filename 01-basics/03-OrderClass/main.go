package main

import "fmt"

func main() {
	price := int32(0)

	fmt.Scanln(&price)

	if price < 0 {
		fmt.Println("Ошибка. Сумма заказа не может быть отрицательной.")
		return
	}

	if price < 10_000 {
		fmt.Println("small")
	} else if price < 50_000 {
		fmt.Println("regular")
	} else {
		fmt.Println("large")
	}
}
