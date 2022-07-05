package main

func main() {
	startGame()

	winner := "" //"Player" || "Monster" || ""

	//this loop will run while there is no winner, and end when we have a winner
	for winner == "" {
		winner = executeRound()
	}

	endGame()

}

func startGame() {}

func executeRound() string {}

func endGame() {}
