package main

import (
	"fmt"
	"myapp/game"
)

func main() {
	displayChannel := make(chan string)
	roundChannel := make(chan int)

	game := game.Game{
		DisplayChannel: displayChannel,
		RoundChannel:   roundChannel,
		Round: game.Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}

	go game.Rounds()
	game.ClearScreen()
	game.PrintIntro()

	for {
		game.RoundChannel <- 1
		<-game.RoundChannel // syntax means "wait for something to happen"

		if game.Round.RoundNumber > 3 {
			break
		}

		if !game.PlayRound() {
			game.RoundChannel <- -1
			<-game.RoundChannel
		}
	}

	fmt.Println("Final score")
	fmt.Println("-----------")
}
