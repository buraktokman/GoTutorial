package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered in f", err)
		}
	}()
	panic("PANIC")

	fmt.Println("End")
}
