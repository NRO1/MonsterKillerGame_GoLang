package interaction

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

//Return user's choice
func PlayerChoice(specialAttackAvail bool) string {
	//A constant loop for getting the correct user
	for {
		playerChoice, _ := getPlayerInput()

		if playerChoice == "1" {
			return "ATTACK"
		} else if playerChoice == "2" {
			return "HEAL"
		} else if playerChoice == "3" && specialAttackAvail {
			return "SPECIAL_ATTACK"
		}
		fmt.Println("Invalid user input. Try again please")
	}
}

//Get and purge user input
func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")

	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	if runtime.GOOS == "windows" {
		userInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		userInput = strings.Replace(userInput, "\n", "", -1)
	}
	return userInput, err
}
