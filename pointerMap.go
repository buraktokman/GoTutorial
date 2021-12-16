package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n------ Map ------")

	// ARRAY
	a := map[string]string{"foo": "bar", "bar": "buz"}
	b := a
	fmt.Println(a, b)

	a["foo"] = "quz"
	fmt.Println(a, b)
}
