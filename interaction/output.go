package interaction

import "fmt"

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
