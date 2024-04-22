package main

import (
	"fmt"
	example "go-hello/internal"
)

func main() {
	fmt.Println(example.Global)
	example.Defer()
	fmt.Println(example.Global)
}
