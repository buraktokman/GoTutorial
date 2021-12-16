package main

import (
	"fmt"
)

func main() {
	var ms *myStruct
	// ms = &myStruct{foo: 42}
	fmt.Println(ms)

	// ms = new(myStruct)
	// fmt.Println(ms)

	(*ms).foo = 42
	fmt.Println((*ms).foo)

	ms.foo = 42
	fmt.Println(ms.foo)
}

type myStruct struct {
	foo int
}
