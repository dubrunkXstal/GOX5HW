package main

import "fmt"

func main() {
	var minn, maxx, summ, order, count int32
	var mean float64

	minn = 2147483647
	maxx = -2147483648
	summ = 0

	fmt.Print("Количество заказов: ")
	fmt.Scanln(&count)

	for i := int32(0); i < count; i++ {
		fmt.Scan(&order)

		if order < minn {
			minn = order
		}
		if order > maxx {
			maxx = order
		}
		summ += order
	}

	if count == 0 {
		fmt.Println("Пустой ввод.")
	} else {
		mean = float64(summ) / float64(count)
		fmt.Printf("Минимум:  %d\n", minn)
		fmt.Printf("Максимум: %d\n", maxx)
		fmt.Printf("Сумма:    %d\n", summ)
		fmt.Printf("Среднее:  %f\n", mean)
	}
}
