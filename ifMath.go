package main

import (
	"fmt"
	"math"
)

func main() {
	myNum := 0.1
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("same")
	} else {
		fmt.Println("different")
	}

	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("same")
	} else {
		fmt.Println("different")
	}
}
