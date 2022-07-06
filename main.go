package main

import (
	"github.com/NRO1/MonsterKillerGame_GoLang/actions"
	"github.com/NRO1/MonsterKillerGame_GoLang/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()

	winner := "" //"Player" || "Monster" || ""

	//this loop will run while there is no winner, and end when we have a winner
	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowActions(isSpecialRound)
	userChoice := interaction.PlayerChoice(isSpecialRound)

	var playerAttackDmg int
	var playerHealVal int
	var monsterAttackDmg int

	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealVal = actions.Heal()
	} else {
		playerAttackDmg = actions.AttackMonster(true)
	}

	monsterAttackDmg = actions.AttackPlayer()

	plHealth, moHealth := actions.GetHealthAmounts()

	roundData := interaction.RoundData{
		Action:           userChoice,
		PlayerHealth:     plHealth,
		MonsterHealth:    moHealth,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealVal:    playerHealVal,
		MonsterAttackDmg: monsterAttackDmg,
	}

	interaction.ShowRoundStats(&roundData)

	gameRounds = append(gameRounds, roundData)

	if plHealth <= 0 {
		return "Monster"
	} else if moHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
