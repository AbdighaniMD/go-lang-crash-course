package main

import "fmt"

func main() {

	color := "blue"

	if color == "red" {
		fmt.Println("color is red")
	} else {
		fmt.Println("It's different color than red.")
	}

	x := 15
	y := 10
	//You can add or ||, and && etc.
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if y < x {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//Switch cases
	switch color {
	case "red":
		fmt.Println("The color is red")
	case "blue":
		fmt.Println("The color is blue")
	default:
		fmt.Println("It's different color than red or blue.")
	}

	// Short Statements
	if n := 15; n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	} else {
		fmt.Printf("%d is odd\n", n)
	}

	/*
		Defer:
		The deferred call's arguments are evaluated immediately,
			but the function call is not executed until the surrounding function returns.
	*/
	defer fmt.Println("defer Statement")

	fmt.Println("Abdighani used: ")

	/*
		Stacking defers:
		Deferred function calls are pushed onto a stack.
		When a function returns, its deferred calls are executed in last-in-first-out order.
	*/

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
