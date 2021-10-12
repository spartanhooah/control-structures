package game

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

type Game struct {
	DisplayChannel chan string
	RoundChannel   chan int
	Round          Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (game *Game) Rounds() {
	// use select to process input in channels
	// print to screen
	// track the round number

	for {
		select {
		case round := <- game.RoundChannel:
			game.Round.RoundNumber += round
			game.RoundChannel <- 1
		case message := <- game.DisplayChannel:
			fmt.Println(message)
		}
	}
}

// clearScreen clears the screen
func (game *Game) ClearScreen() {
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

func (game *Game) PrintIntro() {
	fmt.Println("Rock, Paper, Scissors")
	fmt.Println("---------------------")
	fmt.Println("Game is played for three rounds, best of three wins")
}

func (game *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())

	playerValue := -1
	fmt.Println("\nRound", game.Round.RoundNumber)
	fmt.Println("--------")
	fmt.Print("Please enter rock, paper, or scissors -> ")

	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(2)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	game.DisplayChannel <- fmt.Sprintf("\nPlayer chose %s", strings.ToUpper(playerChoice))

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
		game.DisplayChannel <- "It's a draw"
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				game.computerWins()
			} else {
				game.playerWins()
			}
		case PAPER:
			if computerValue == SCISSORS {
				game.computerWins()
			} else {
				game.playerWins()
			}
		case SCISSORS:
			if computerValue == ROCK {
				game.computerWins()
			} else {
				game.playerWins()
			}
		default:
			game.DisplayChannel <- "Invalid choice!"
			return false
		}
	}

	return true
}

func (game *Game) computerWins() {
	game.Round.ComputerScore++
	game.DisplayChannel <- "Computer wins!"
}

func (game *Game) playerWins() {
	game.Round.PlayerScore++
	game.DisplayChannel <- "Player wins!"
}
