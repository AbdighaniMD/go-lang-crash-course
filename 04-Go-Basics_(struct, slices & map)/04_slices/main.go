package main

import (
	"fmt"
)

func main() {
	//slice is dynamically-sized: that  does not store any data, it just describes a section of an underlying array.
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Array:", primes)

	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	/// primes [low : high]
	var a []int = primes[1:4]
	fmt.Println("slice A:", a) // [3, 5, 7]

	// note Slices are like references to arrays
	var b []int = primes[3:]
	fmt.Println("slice B:", b) // [7, 11, 13]

	//Changing the elements of a slice modifies the corresponding elements of its underlying array.
	b[0] = 111

	//Other slices that share the same underlying array will see those changes.
	fmt.Println("Both slices: after changed", a, b)
	fmt.Println("Array: after changed", primes)

	//  Slice literals
	//This is an array literal:
	array_literal := [3]bool{true, true, false}
	fmt.Println("Array literal build:", array_literal)

	// this creates the same array as above, then builds a slice that references it:
	Slice_literal := []bool{true, true, false}
	fmt.Println("Slice literal build:", Slice_literal)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// note The zero value of a slice is 0 === nil.

	//Creating a slice with make
	M := make([]int, 5)
	fmt.Println("M ", "Len =", len(M)) // len(a)=5
	F := make([]int, 0, 5)
	fmt.Println("F:", "Len =", len(F), "cap", cap(F)) // len(b)=0, cap(b)=5
	// Extend its length.
	F = F[:cap(F)]
	fmt.Println("F:", "Len =", len(F), "cap", cap(F)) // len(b)=5, cap(b)=5
	// Drop its first one values.
	F = F[1:]
	fmt.Println("F:", "Len =", len(F), "cap", cap(F)) // len(b)=4, cap(b)=4

	// Appending to a slice
	var G []int
	fmt.Println(G)
	// append works on nil slices.
	G = append(G, 0, 3)
	fmt.Println(G)

}
