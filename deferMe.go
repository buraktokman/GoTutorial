package main

import (
	"fmt"
)

func main() {
	// START
	fmt.Println("start")

	// COME BACK TO HERE (exec. bottom to top)
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	// END
	fmt.Println("end")
}
