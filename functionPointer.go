package main

import (
	"fmt"
)

func main() {
	// var count int = 5
	// for i := 0; i < count; i++ {
	// 	sayMessage("Hello Go!", i)
	// }
	greeting := "hello"
	name := "world"
	sayGreeting(&greeting, &name)
	fmt.Println("name: ", name)
}

func sayGreeting(greeting, name *string) {
	fmt.Printf("greeting: %v -- name: %v\n", *greeting, name)
	*name = "ted"
	fmt.Println("name (sayGreeting): ", *name)

}
