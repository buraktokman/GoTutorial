package main

import (
	"fmt"
	"strconv"
)

var MYAGE int = 19

var myAge5 int = 25

var (
	actorName string = "Tom Cruise"
	companion string = "Sally"
	season    int    = 11
)

var (
	counter int = 0
)

func main() {
	var myAge int
	myAge = 25

	var myAge2 int = 25

	myAge3 := 25
	myAge4 := 25.

	var HttpRequest string = "https://www.youtube.com/"

	fmt.Println("I'm relocating to a new planet!")
	fmt.Println(HttpRequest)

	fmt.Println("myAge", myAge)
	fmt.Println("myAge2", myAge2)

	fmt.Printf("%v -> type: %T\n", myAge3, myAge3)
	fmt.Printf("%v -> type: %T\n\n", myAge4, myAge4)

	// CHANGE VAR TYPE
	var i int = 23
	fmt.Printf("%v -> type: %T\n", i, i)
	var j float32 = float32(i)
	fmt.Printf("%v -> type: %T\n", j, j)

	// INT TO STRING
	var k int = 23
	var j1 string = strconv.Itoa(k)
	fmt.Printf("%v -> type: %T\n", j1, j1)

}

/* SUMMARY
- Can't redeclare variables, but can shadow them
- All variables must be used
- Visibility
- Naming
	- pascal or camelCase)
	- Capitalize acronyms (HTTP, URL)
- Type conversions
	- destionationType(variable)
	- Use 'strconv' package for strings
*/
