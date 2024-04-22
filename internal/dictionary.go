package example

import "fmt"

func Dictionary() {
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

func Find(arr []int, k int) [][]int {
	pairs := make([][]int, 0)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == k {
				pairs = append(pairs, []int{i, j})
			}
		}
	}

	return pairs
}

func RemoveDuplicates(input []string) []string {
	unique := make(map[string]bool)
	result := make([]string, 0)

	for _, v := range input {
		if !unique[v] {
			unique[v] = true
			result = append(result, v)
		}
	}

	return result
}
