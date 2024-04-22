package example

import "fmt"

var Global = 5

func Defer() {
	defer func(checkout int) {
		Global = checkout
	}(Global)

	Global = 42
	fmt.Println(Global)
}
