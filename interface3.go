package main

import (
	"fmt"
)

func main() {
	// var i interface{} = 0
	// var i interface{} = "0"
	var i interface{} = true

	// TYPE WILL BE CHECKED
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I dunno")
	}
}
