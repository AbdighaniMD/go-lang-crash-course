package main

import "fmt"

func main() {
	// An array in go are fixed size
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("Index:0 =", a[0], "Index:1 =", a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// 2d array
	var twoD [2][3]int
	for i := 0; i < 2; i++ { // row
		for j := 0; j < 3; j++ { //column
			twoD[i][j] = i + j
			// fmt.Println(i, j, i+j)
		}
	}
	fmt.Println("2D ", twoD)

	// The range form of the for loop iterates over a arrays, slice or map. For each loop
	letters := []string{"A", "B", "C", "D"}
	for index, value := range letters {
		fmt.Println("Index: ", index, "Value: ", value)
	}

}
