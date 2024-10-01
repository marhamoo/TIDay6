package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Define constans for choices
const ( 
	Rock = iota // 0
	Paper		// 1
	Scissors	// 2
)

// Helper function to convert choice to string
func choiceToString(choice int) string{
	switch choice {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	default:
		return "Unknown"
	}
}

func main(){
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Ask the player to choose between Rock, Paper, Scissors
	fmt.Println("Let`s play Rock, Paper, Scissors!")
	fmt.Println("Enter your choice: 0 for Rock, 1 for Paper, 2 for Scissors:")

	var playerChoice int
	fmt.Scanf("%d", &playerChoice)

	// Validate player`s input
	if playerChoice < Rock || playerChoice > Scissors {
		fmt.Println("Invalid choice, please enter 0, 1, or 2.")
		return
	}

	// Randomly generate computer`s choice
	computerChoice := rand.Intn(3)

	// Display both choices
	fmt.Printf("You chose: %s\n", choiceToString(playerChoice))
	fmt.Printf("Computer chose: %s\n", choiceToString(computerChoice))

	// Determine the winner
	if playerChoice == computerChoice {
		fmt.Println("It`s a tie!")
	} else if (playerChoice == Rock && computerChoice == Scissors) ||
		(playerChoice == Paper && computerChoice == Rock) ||
		(playerChoice == Scissors && computerChoice == Paper) {
			fmt.Println("You win!")
		} else {
			fmt.Println("Computer wins!")
		}
}

