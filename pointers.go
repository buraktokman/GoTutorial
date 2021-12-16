package main

import (
	"fmt"
)

func main() {
	// a := 45
	var a int = 45
	// POINT b to a
	var b *int = &a
	// b := a
	fmt.Printf("a-address: %v \t a-value: %v\n", &a, a)
	fmt.Printf("b-address: %v \t b-value: %v\n", b, *b)

	fmt.Println("\nChange a to 27")

	a = 27
	fmt.Printf("a-address: %v \t a-value: %v\n", &a, a)
	fmt.Printf("b-address: %v \t b-value: %v\n", b, *b)

	fmt.Println("\nChange b to 14")

	*b = 14
	fmt.Printf("a-address: %v \t a-value: %v\n", &a, a)
	fmt.Printf("b-address: %v \t b-value: %v\n", b, *b)
}
