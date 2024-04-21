package main

import "fmt"

func dictionary() {
	groceries := map[string]int{"хлеб": 50, "молоко": 100, "масло": 200, "колбаса": 500, "соль": 20, "огурцы": 200, "сыр": 600, "ветчина": 700, "буженина": 900, "помидоры": 250, "рыба": 300, "хамон": 1500}
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	total := 0

	for k, v := range groceries {
		if v > 300 {
			fmt.Printf("%s: %d\n", k, v)
		}
	}

	for _, v := range order {
		total += groceries[v]
	}

	fmt.Printf("Заказ: %v, cумма: %d\n", order, total)
}
