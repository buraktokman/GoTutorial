package main

import (
	"fmt"
)

var (
	actorName string = "Tom Cruise"
)

var (
	counter int = 0
)

func main() {

	fmt.Println("\nBoolean")

	var myBool bool
	fmt.Printf("%v -> type: %T\n", myBool, myBool)

	var myBool2 bool = true
	fmt.Printf("%v -> type: %T\n", myBool2, myBool2)

	// EQUALS TO RESULT
	fmt.Println("\nResult")
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v -> type: %T\n", n, n)
	fmt.Printf("%v -> type: %T\n", m, m)

	// UINT
	fmt.Println("\nUnsigned INT")
	var n1 uint16 = 42
	fmt.Printf("%v -> type: %T\n", n1, n1)

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// TYPE CONVERSION
	fmt.Println("\nType Convesion")
	var b2 int8 = 3
	fmt.Println(a + int(b2))

	// BIT OPERATOR
	fmt.Println("\nBit Operator")
	at := 10              // 1010
	bt := 3               // 0011
	fmt.Println(at & bt)  // 0010
	fmt.Println(at | bt)  // 1011
	fmt.Println(at ^ bt)  // 1001
	fmt.Println(at &^ bt) // 1000

	// BIT SHIFTING
	fmt.Println("\nBit Shifting")
	as := 8              // 1000
	fmt.Println(as << 3) //
	fmt.Println(as >> 3) //

	// FLOATING POINT
	fmt.Println("\nFloating Point")
	af := 10.2
	bf := 3.7
	fmt.Println(af + bf)
	fmt.Println(af - bf)
	fmt.Println(af * bf)
	fmt.Println(af / bf)

	// COMPLEX NUMBERS
	fmt.Println("\nComplex Numbers")
	var n3 complex64 = 1 + 2i
	fmt.Println("%v -> type: %T", n3, n3)
	fmt.Println("%v -> type: %T", real(n3), real(n3))
	fmt.Println("%v -> type: %T", imag(n3), imag(n3))

	// STRING
	fmt.Println("\nString")
	s := "this is a str"
	//s2[2] = "u"                                 // THIS WON'T WORK!
	fmt.Printf("%v -> type: %T\n", s[2], s[2]) // print byte

	// s2 := "this is also a string"
	// b := []byte(s2)
	// fmt.Println("%v -> type: %T", b, b)

	// RUNE
	var r rune = 'a'
	fmt.Println("%v -> type: %T", r, r)

}

/* SUMMARY
Strings
- Immutable
- Concatenated w/ '+' operator
*/
