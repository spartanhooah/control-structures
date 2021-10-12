package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	clearScreen()

	for i := 1; i < 4; i++ {
		fmt.Println("\nRound", i)
		fmt.Println("--------")
		fmt.Print("Please enter rock, paper, or scissors -> ")
		playerChoice := ""
		playerValue := -1
		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)
		computerValue := rand.Intn(2)

		if playerChoice == "rock" {
			playerValue = ROCK
		} else if playerChoice == "paper" {
			playerValue = PAPER
		} else if playerChoice == "scissors" {
			playerValue = SCISSORS
		}

		switch computerValue {
		case ROCK:
			fmt.Println("Computer chose ROCK")
		case PAPER:
			fmt.Println("Computer chose PAPER")
		case SCISSORS:
			fmt.Println("Computer chose SCISSORS")
		default:
		}

		if playerValue == computerValue {
			fmt.Println("It's a draw")
			i--
		} else {
			switch playerValue {
			case ROCK:
				if computerValue == PAPER {
					fmt.Println("Computer wins")
				} else {
					fmt.Println("Player wins")
				}
			case PAPER:
				if computerValue == SCISSORS {
					fmt.Println("Computer wins")
				} else {
					fmt.Println("Player wins")
				}
			case SCISSORS:
				if computerValue == ROCK {
					fmt.Println("Computer wins")
				} else {
					fmt.Println("Player wins")
				}
			default:
				fmt.Println("Invalid choice")
				i--
			}
		}
	}
}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
