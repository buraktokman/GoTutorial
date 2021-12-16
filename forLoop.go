package main

import (
	"fmt"
)

func main() {
	// INCREMENT i AND j
	fmt.Println("\n---- i and j ----")
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	fmt.Println("\n---- modulo ----")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}

	fmt.Println("\n---- z ----")
	z := 0
	for ; z < 5; z++ {
		fmt.Println(z)
	}
	fmt.Println(z)

	fmt.Println("\n---- k ----")
	k := 0
	for k < 5 {
		fmt.Println(k)
		k++
	}

	fmt.Println("\n---- h ----")
	h := 1
	for {
		fmt.Println(h)
		if h%9 == 0 {
			// continue
			break
		}
		h++
		// fmt.Println(h)
	}

	fmt.Println("\n---- loop label  ----")
outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d -- %d\n", i, j)
			if i*j >= 3 {
				break outerLoop
			}
		}
	}

	fmt.Println("\n---- collection ----")
	s := []int{1, 2, 3}
	for k1, v := range s {
		fmt.Println(k1, v)
	}

	fmt.Println("\n---- range with map ----")
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("\n---- range with map string ----")
	s2 := "hello go!"
	for k, v := range s2 {
		fmt.Println(k, string(v))
	}
	// USE UNDERSCORE IF YOU DON'T NEED TO USE A VAR.
	for _, v := range s2 {
		fmt.Println(string(v))
	}

}
