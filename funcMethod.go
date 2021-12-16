package main

import (
	"fmt"
)

type greeter struct {
	greeter string
	name    string
}

// type counter int { }

func main() {
	g := greeter{
		greeting: "hello",
		name:     "world",
	}
	g.greet()
	fmt.Println("new name is: ", g.name)
}

func (g *greeter) greet() {
	fmt.Printf("%s, %s\n", g.greeter, g.name)
	g.name = ""
}
