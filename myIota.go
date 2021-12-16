package main

import (
	"fmt"
)

const (
	// _ = iota + 5 // iota starts from 5 -> offset +5
	catSpecialist   = iota // 0
	dogSpecialist          // 1
	snakeSpecialist        // 2
)

func main() {
	fmt.Println("\nIota")

	var specialistType int = catSpecialist

	fmt.Printf("%v\n", specialistType == catSpecialist)

}

/* SUMMARY
- Immutable, but can be shadowed
- Replaced by the compilter at compile time

*/
