package main

import (
	"fmt"
)

func main() {
	switch i := 2 + 3; i {
	// switch 2 {
	case 1:
		fmt.Println("one")
	case 2, 3, 4:
		fmt.Println("two, 2, 3, 4")
	default:
		fmt.Println("another number")
	}

	// SWITCH
	i2 := 10
	switch {
	case i2 <= 10:
		fmt.Println("i2 <= 10")
		fallthrough // asagidaki case'i de calistir
	case i2 <= 20:
		fmt.Println("i2 <= 20")
	default:
		fmt.Println("i2 > 20")
	}

	// SWITCH TYPE
	// var i interface{} = 1
	var i interface{} = [3]int{}
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	case [3]int:
		fmt.Println("i is an array of int")
	default:
		fmt.Println("i is another type")
	}
}
