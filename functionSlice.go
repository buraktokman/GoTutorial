package main

import (
	"fmt"
)

func main() {
	s := sum("the sum is", 1, 2, 3, 4, 5)
	fmt.Println(s)
}

func sum(msg string, values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	// fmt.Println(msg, result)
	return result
}

func sum2(msg string, values ...int) (result int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return
}
