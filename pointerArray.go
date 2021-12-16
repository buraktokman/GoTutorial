package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n------ Array ------")

	// ARRAY
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(a, b)

	a[1] = 42
	fmt.Println(a, b)

	fmt.Println("\n------ Slice ------")

	// SLICE
	a1 := []int{1, 2, 3}
	b1 := a1
	fmt.Println(a1, b1)

	a[1] = 42
	fmt.Println(a1, b1)
}
