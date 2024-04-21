package main

import (
	"fmt"
	"sort"
)

func slice() {
	var s []int

	for i := 0; i <= 100; i++ {
		s = append(s, i)
	}

	s = append(s[:10], s[91:]...)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))

	fmt.Println(s, len(s))
}
