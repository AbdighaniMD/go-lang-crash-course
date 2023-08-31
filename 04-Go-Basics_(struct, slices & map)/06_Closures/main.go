package main

import "fmt"

//  Functions are values too. They can be passed around just like other values.
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Function closures: outer function
func message() func() string {

	// variable defined
	text := "Scaler"

	// return a nested anonymous function without a name
	return func() string {
		text = "Hi, " + text
		return text
	}
}

func main() {
	// note Function values may be used as function arguments and return values.
	hypot := func(x, y float64) float64 {
		return x * y
	}
	fmt.Println(hypot(5, 5))
	fmt.Println(compute(hypot))

	//  call the outer function
	mssg := message()

	// call the inner function
	fmt.Println(mssg())

}
