package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

// note you can omit the type from all but the last
func getSum(num1, num2 int) int {
	return num1 + num2
}

// multi return functions
func swap(x, y string) (string, string) {
	return y, x
}

/*
Naked Return functions:
if you name your return values, you don't have to return them
the funciton will automatically return these named return values
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(greeting("AbdighaniMD"))
	fmt.Println(getSum(2, 6))
	fmt.Println(swap("Hello", "World"))
	fmt.Println(split(17))
}
