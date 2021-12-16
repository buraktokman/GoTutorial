package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("always true")
	}

	number := 50
	guess := 30

	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("Guess must be between 1 and 100!")
	} else if guess > 100 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}

	// 1 < guess < 100
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("too low")
		}
		if guess > number {
			fmt.Println("too high")
		}
		if guess == number {
			fmt.Println("correct")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}
	fmt.Println(!true)
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
