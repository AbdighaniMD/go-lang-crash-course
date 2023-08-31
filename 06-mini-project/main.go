package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var options = []string{"Rock", "Paper", "Scissors"}
	rand.Seed(time.Now().UnixNano())

	userScore := 0
	computerScore := 0

	fmt.Println("Welcome to Rock-Paper-Scissors Game")

	for {
		userOption := ""
		fmt.Println("Enter your choice: ")
		fmt.Scanln(&userOption)

		isValidChoice := false

		var option string

		for _, option = range options {
			if userOption == option {
				isValidChoice = true
				break
			}

		}

		if !isValidChoice {
			fmt.Println("Enter 'Rock', 'Paper', 'Scissors'")
			continue
		}

		computerOption := options[rand.Intn(len(options))]

		// let's datermine the winner
		if userOption == computerOption {
			fmt.Println("its a tie")
		} else if userOption == "Rock" && computerOption == "Scissors" || userOption == "Paper" && computerOption == "Rock" || userOption == "Scissors" && computerOption == "Paper" {
			fmt.Println("You Win!")
			userScore++
		} else {
			fmt.Println("Computer Wins!")
			computerScore++
		}

		fmt.Println("User Score is:", userScore)
		fmt.Println("Computer Score is:", computerScore)
		fmt.Println(" ")

		var ask_user string

		fmt.Println("Continue the game, 'Y/N' !")
		fmt.Scanln(&ask_user)

		if ask_user == "N" || ask_user == "n" {
			break
		} else {
			continue
		}

	}

}
