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
		case round := <-game.RoundChannel:
			game.Round.RoundNumber += round
			game.RoundChannel <- 1
		case message := <-game.DisplayChannel:
			fmt.Println(message)
			game.DisplayChannel <- ""
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
	game.DisplayChannel <- `
Rock, Paper, Scissors
---------------------
Game is played for three rounds, best of three wins`
	<-game.DisplayChannel
	game.DisplayChannel <- ""
	<-game.DisplayChannel
}

func (game *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())

	playerValue := -1
	game.DisplayChannel <- fmt.Sprintf(
		`Round %d
--------`,
		game.Round.RoundNumber)
	<-game.DisplayChannel
	fmt.Print("Please enter rock, paper, or scissors -> ")

	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(2)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	game.DisplayChannel <- ""
	<-game.DisplayChannel
	game.DisplayChannel <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))
	<-game.DisplayChannel

	switch computerValue {
	case ROCK:
		game.DisplayChannel <- "Computer chose ROCK"
		<-game.DisplayChannel
	case PAPER:
		game.DisplayChannel <- "Computer chose PAPER"
		<-game.DisplayChannel
	case SCISSORS:
		game.DisplayChannel <- "Computer chose SCISSORS"
		<-game.DisplayChannel
	default:
	}

	game.DisplayChannel <- ""
	<-game.DisplayChannel
	if playerValue == computerValue {
		game.DisplayChannel <- "It's a draw"
		<-game.DisplayChannel
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
			<-game.DisplayChannel
			return false
		}
	}

	return true
}

func (game *Game) computerWins() {
	game.Round.ComputerScore++
	game.DisplayChannel <- "Computer wins!"
	<-game.DisplayChannel
}

func (game *Game) playerWins() {
	game.Round.PlayerScore++
	game.DisplayChannel <- "Player wins!"
	<-game.DisplayChannel
}

func (game *Game) PrintSummary() {
	game.DisplayChannel <- fmt.Sprintf(`
Final score
-----------
Player: %d/3, Computer: %d/3`, game.Round.PlayerScore, game.Round.ComputerScore)
<-game.DisplayChannel

	if game.Round.PlayerScore > game.Round.ComputerScore {
		game.DisplayChannel <- "Player wins the game!"
		<-game.DisplayChannel
	} else {
		game.DisplayChannel <- "Computer wins the game!"
		<-game.DisplayChannel
	}
}
