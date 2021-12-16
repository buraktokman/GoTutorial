package main

import (
	"fmt"
)

func main() {
	d, err := divide(5.0, 0.0)
	// PROCESS IF ERROR
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		// panic("error")
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
