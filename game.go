package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Define constan choices
const (
	Rock     = iota //0
	Paper           //1
	Scissors        //2

)

//fungsi const

func choices(choice int) string {

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

func main() {
	//Scan generator number
	rand.Seed(time.Now().UnixNano())
	//display pilihan
	fmt.Println("Let's Play Rock, Paper, Scissors!")
	fmt.Println("Enter your choice: 0 for Rock, 1 for Paper, 2 for Scissors:")

	var playerChoice int
	fmt.Scanf("%d", &playerChoice)

	//input player
	if playerChoice < Rock || playerChoice > Scissors {
		fmt.Println("Invalid choice, please enter 0, 1, or 2.")
		return
	}

	//Random Generator
	computerChoice := rand.Intn(3)

	//Display 2 pilihan
	fmt.Printf("You are: %s\n", choices(playerChoice))
	fmt.Printf("The computer: %s\n", choices(computerChoice))

	//operasi
	if playerChoice == computerChoice {
		fmt.Println("it's a tie")
	} else if (playerChoice == Rock && computerChoice == Scissors) ||
		(playerChoice == Scissors && computerChoice == Paper) ||
		(playerChoice == Paper && computerChoice == Rock) {
		fmt.Println("You win")
	} else {
		fmt.Println("You Lose :(")
	}

}
