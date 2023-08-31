package main

import "fmt"

func main() {
	/*
		Go has pointers.
			A pointer holds the memory address of a value.
		Unlike C, Go has no pointer arithmetic.
	*/

	a := 5
	// The & operator generates a pointer to its operand.
	b := &a //the memory address of the a

	fmt.Println("a:=", a, "b:=", b)
	fmt.Printf("%T\n", b) //we get *int (pointer to int)

	a = a + 1

	//User * to read value from memory address
	fmt.Println(*b)
	//fmt.Println(*&b) // same thing as line 13

	//Change value with pointer
	*b = 10
	fmt.Println(a)

}
