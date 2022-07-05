package main

import (
	"fmt"

	"github.com/NRO1/MonsterKillerGame_GoLang/interaction"
)

var currentRound = 0

func main() {
	startGame()

	winner := "" //"Player" || "Monster" || ""

	//this loop will run while there is no winner, and end when we have a winner
	for winner == "" {
		winner = executeRound()
	}

	endGame()

}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowActions(isSpecialRound)
	userChoice := interaction.PlayerChoice(isSpecialRound)

	fmt.Println(userChoice)
	return ""
}

func endGame() {}
