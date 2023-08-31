package main

import "fmt"

func main() {
	//Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		//i = i + 1
		i++
	}

	//Short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	/*
		FizzBuzz:
		Write a program that takes two numbers and prints the numbers.
		But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”.
		For numbers which are multiples of both three and five print “FizzBuzz”.
	*/
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 && i%5 == 0 { //
			fmt.Println("“FizzBuzz”")
		}
	}
}
