package main

import "fmt"

func generation() {
	var year int

	fmt.Println("Введите ваш год рождения: ")
	fmt.Scan(&year)

	switch {
	case year >= 1946 && year <= 1964:
		fmt.Println("Привет, бумер!")
	case year >= 1965 && year <= 1980:
		fmt.Println("Привет, представитель X!")
	case year >= 1981 && year <= 1996:
		fmt.Println("Привет, миллениал!")
	case year >= 1997 && year <= 2012:
		fmt.Println("Привет, зумер!")
	case year >= 2013:
		fmt.Println("Привет, альфа!")
	default:
		fmt.Println("Вы ввели некорректный год рождения.")
	}
}
