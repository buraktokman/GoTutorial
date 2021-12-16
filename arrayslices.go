package main

import (
	"fmt"
)

func main() {
	fmt.Println("\nArrays")

	grade1 := 97
	grade2 := 85
	grade3 := 52
	fmt.Printf("Grades: %v, %v, %v\n", grade1, grade2, grade3)

	grades := [3]int{97, 85, 52}
	fmt.Printf("Grades: %v, %v, %v\n", grades[0], grades[1], grades[2])
	// fmt.Printf("%v -> type: %T\n", c3, c3)

	// SIZE UNDEF.
	// grades1 := [...]int{97, 85, 52}
	// fmt.Printf("Grades: %v, %v, %v\n", grades[0], grades[1], grades[2])

	var students [3]string
	fmt.Printf("Students: %v", students)

	students[0] = "Lisa"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of students: %v\n", len(students))

	fmt.Println("Matrix")
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}

	// PRINT MATRIX
	for i := 0; i < len(identityMatrix); i++ {
		for j := 0; j < len(identityMatrix[i]); j++ {
			fmt.Printf("%v ", identityMatrix[i][j])
		}
		fmt.Println()
	}

	// COPY OF ARRAY
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// POINTER
	fmt.Println("\nPointer")
	b2 := &a
	b2[1] = 5
	fmt.Println(a)
	fmt.Println(b2)

	// -----------
	// Slices
	fmt.Println("\nSlices")

	//b3 := &a
	b3 := a
	b3[1] = 5
	fmt.Println(a)
	fmt.Println(b3)
	fmt.Printf("Length: %v\n", len(b3))
	fmt.Printf("Capacity: %v\n", cap(b3))

	fmt.Printf("\nSlice")
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b4 := a1[:]  // slice of all elements
	c := a1[3:]  // slice from 4th element to end
	d := a1[:6]  // slice first 6 elements
	e := a1[3:6] // slice the 4th, 5th, and 6th elements
	fmt.Println(a1)
	fmt.Println(b4)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Printf("\nMake")
	// a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// OR
	a2 := make([]int, 3, 100) // slice with lenghth=10 and capacity=100
	//a2 := make([]int, 10) // slice with length=10 and capacity=100=
	fmt.Println(a2)
	fmt.Printf("Length: %v\n", len(a2))
	fmt.Printf("Capacity: %v\n", cap(a2))

	// APPEND SINGLE
	a2 = append(a2, 11)
	fmt.Println(a2)
	fmt.Printf("Length: %v\n", len(a2))
	fmt.Printf("Capacity: %v\n", cap(a2))

	// APPEND MULTIPLE
	a2 = append(a2, 12, 13, 14, 15)
	fmt.Println(a2)
	fmt.Printf("Length: %v\n", len(a2))
	fmt.Printf("Capacity: %v\n", cap(a2))

	// -----------
	// POP
	fmt.Println("\nPop")
	a3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b5 := a3[:len(a3)-1]
	fmt.Println(b5)

	b6 := append(a3[:2], a3[3:]...)
	fmt.Println(b6)

}

/* SUMMARY

 */
