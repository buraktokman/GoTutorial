package main

import (
	"fmt"
)

func main() {
	fmt.Println("\nMaps")
	// statePopulations := make(map[string]int, 10)
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        26956958,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12741080,
		"Ohio":         11570808,
	}
	// m := map[[3]int]string{}
	// fmt.Println(statePopulations, m)
	fmt.Println(statePopulations)

	// fmt.Println(statePopulations["Ohio"])

	fmt.Println("Change Georgia")
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations["Georgia"])

	// DELETE
	fmt.Println("Delete")
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations["Georgia"])
	fmt.Println(statePopulations)

	fmt.Println("Key exists?  ")
	pop, ok := statePopulations["Ohio"]
	fmt.Println(pop, ok)

	// MANIPULATING MAP AFFECTS ORIGINAL
	fmt.Println("\nManipulating map")
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)

	// IF
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println("Florida ->", pop)
	} else {
		fmt.Println("Florida is not in the map")
	}

}
