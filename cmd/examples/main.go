package main

import (
	"fmt"
	example "go-hello/internal"
)

func main() {
	ar, ok := example.Area(example.Square)

	if ok {
		fmt.Println(ar(2))
	} else {
		fmt.Println("Нет такой фигуры")
	}
}
