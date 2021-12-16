package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")

	defer fmt.Println("Defer 1")

	// PANIC EXE. AFTER DEFER
	panic("something bad happened")

	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)

	fmt.Println("End")
}
