package main

import "fmt"

func addition(a, b int) int {
	return a + b
}

func substraction(a, b int) int {
	return a - b
}

func multiplication(a, b int) int {
	return a * b
}

func division(a, b int) int {
	return a / b
}

func main() {
	//Go User Input
	var option int
	var a, b int

	for {
		fmt.Println("enter 1 -> Addition, 2 -> Substraction, 3 -> Multiplication, 4 -> division, 5 -> exits")
		// "&option" is the memory locations of option .
		fmt.Scanln(&option)

		if option == 5 {
			break
		}
		fmt.Println("enter value one numbers")
		fmt.Scanln(&a)
		fmt.Println("enter value two numbers")
		fmt.Scanln(&b)

		switch option {
		case 1:
			result := addition(a, b)
			fmt.Println("Your numbers are:", a, "and", b, "with additon = ", result)
		case 2:
			result := substraction(a, b)
			fmt.Println("Your numbers are:", a, "and", b, "with substraction = ", result)
		case 3:
			result := multiplication(a, b)
			fmt.Println("Your numbers are:", a, "and", b, "with multiplication = ", result)
		case 4:
			result := division(a, b)
			fmt.Println("Your numbers are:", a, "and", b, "with division = ", result)
		default:
			fmt.Println(option, " is not a mathematics operation in the list. Choice again or exit.")
		}
	}
}
