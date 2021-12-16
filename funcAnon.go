package main

import (
	"fmt"
)

func main() {
	// SYNC. EXECUTION
	for i := 0; i < 5; i++ {
		func(i int) {
			msg := "I'm anon function -> "
			fmt.Println(msg, i)
		}(i)
	}

	// ASSIGN FUNC TO VAR.
	var f func() = func() {
		fmt.Println("hello")
	}
	f()

}
