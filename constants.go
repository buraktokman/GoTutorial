package main

import (
	"fmt"
)

const a int16 = 27

const (
	a3 = iota
	b3 = iota
	c3 = iota
)
const (
	a4 = iota
)

func main() {
	fmt.Println("\nConstant")
	const myConst = 42
	// const myConst float64 = math.Sin(1.57)
	fmt.Printf("%v -> type: %T\n", myConst, myConst)

	// GLOBAL
	fmt.Printf("%v -> type: %T\n", a, a)

	// LOCAL
	const a int = 13
	fmt.Printf("%v -> type: %T\n", a, a)

	// const b int16 = 27
	const b = 27
	fmt.Printf("%v -> type: %T\n", a+b, a+b)

	fmt.Println("\nIota")
	fmt.Printf("%v -> type: %T\n", a3, a3)
	fmt.Printf("%v -> type: %T\n", b3, b3)
	fmt.Printf("%v -> type: %T\n", c3, c3)

}

/* SUMMARY

 */
