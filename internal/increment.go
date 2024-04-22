package example

import "fmt"

const (
	one = iota*2 + 1
	three
	five
	seven
	nine
	eleven
)

func Increment() {
	fmt.Println(one, three, five, seven, nine, eleven)
}
