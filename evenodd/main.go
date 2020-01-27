package main

import (
	"fmt"
)

var (
	numSlice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

func main() {
	for i := range numSlice {
		if i % 2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}