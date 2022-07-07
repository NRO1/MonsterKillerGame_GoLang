package interaction

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

//Collecting game data for display
type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	MonsterAttackDmg int
	PlayerHealVal    int
	PlayerHealth     int
	MonsterHealth    int
}

//Greeting the player in the begining of the game, using 3rd party pack=age of ASCII art
func PrintGreeting() {
	ascFigure := figure.NewFigure("MONSTER SLAYER", "", true)
	ascFigure.Print()
	fmt.Println("Starting a new game...")
}

//Show available options for the player
func ShowActions(specialAttackAvail bool) {
	fmt.Println("<------------------>")
	fmt.Println("Choose your action:")
	fmt.Println("<------------------>")
	fmt.Println("(1) Attack")
	fmt.Println("(2) Heal")

	if specialAttackAvail {
		fmt.Println("(3) Special Attack")
	}
}

//Poniter to the struct for memory efficiency
func ShowRoundStats(RoundData *RoundData) {
	if RoundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage \n", RoundData.PlayerAttackDmg)
	} else if RoundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a special attack against the monster for %v damage \n", RoundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed for %v \n", RoundData.PlayerHealVal)
	}

	fmt.Printf("Monster attacked player for %v damage \n", RoundData.MonsterAttackDmg)
	fmt.Printf("Current player health: %v\n", RoundData.PlayerHealth)
	fmt.Printf("Current monster health: %v\n", RoundData.MonsterHealth)
}

func DeclareWinner(winner string) {
	fmt.Println("<------------------>")
	ascFigure := figure.NewColorFigure("Game over", "", "green", true)
	ascFigure.Print()
	fmt.Println("<------------------>")
	fmt.Printf("%v is the WINNER \n", winner)
}

func WriteLogFile(rounds *[]RoundData) {
	exPath, err := os.Executable() //creating an executable path, so we could build this app to be a standalone app
	if err != nil {
		fmt.Println("Writing log file failed. Exiting")
	}

	exPath = filepath.Dir(exPath) //Absolute path to where the executable file is

	file, err := os.Create(exPath + "/GameLog.txt")

	if err != nil {
		fmt.Println("Log file could not be created. Exiting")
		return
	}

	for index, value := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                value.Action,
			"Player attack damage":  fmt.Sprint(value.PlayerAttackDmg),
			"Player heal value":     fmt.Sprint(value.PlayerHealVal),
			"Monster attack damage": fmt.Sprint(value.MonsterAttackDmg),
			"Player health":         fmt.Sprint(value.PlayerHealth),
			"Monster health":        fmt.Sprint(value.MonsterHealth),
		}

		logLine := fmt.Sprintln(logEntry)
		_, err := file.WriteString(logLine)

		if err != nil {
			fmt.Println("Writing to log file failed. Exiting")
			continue
		}
	}

	file.Close()
	fmt.Println("Wrote data to log")
}
