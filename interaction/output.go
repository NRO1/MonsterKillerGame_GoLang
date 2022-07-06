package interaction

import (
	"fmt"
	"os"
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

//Greeting the player in the begining of the game
func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
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
	fmt.Println("GAME OVER!")
	fmt.Println("<------------------>")
	fmt.Printf("%v is the WINNER \n", winner)
}

func WriteLogFile(rounds *[]RoundData) {
	file, err := os.Create("GameLog.txt")

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
