package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:     3,
		actorName:  "Joe Petwee",
		companions: []string{"John", "Paul", "Ringo"},
	}
	fmt.Println(aDoctor)

	fmt.Println(aDoctor.number)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])

	// INLINE STRUCT
	aDoctor2 := struct{ name string }{name: "Joe Petwee"}
	anotherDoctor := &aDoctor2
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor2)
	fmt.Println(anotherDoctor)
}
