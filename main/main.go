package main

import (
	"bufio"
	"fmt"
	"os"
	"tic-tac-go/game/game_controller"
	"tic-tac-go/game/player_controller"
)

func main() {

	stdin := bufio.NewReader(os.Stdin)

	// TODO add configuration for cli vs gui

	// number of players

	playerNum := 0

	for {

		fmt.Printf("How many human players? [1-2] ")
		_, err := fmt.Fscanf(stdin, "%d\n", &playerNum)

		if err != nil { // error
			stdin.ReadString('\n')
			fmt.Println("Number is not valid. Try again.")
		} else if playerNum < 1 || playerNum > 2 { // semantics
			// stdin.ReadString('\n')
			fmt.Println("Number of players is not in range [1-2]. Try again.")
		} else {
			break
		}

	}

	// Player One's name

	var player1Name string

	fmt.Printf("Insert Player One's name: ")
	fmt.Scanf("%s", &player1Name)

	fmt.Printf("Hello %s!\n", player1Name)

	// Player Two's name

	var player2Name string

	if playerNum == 2 {

		fmt.Printf("Insert Player Two's name: ")
		fmt.Scanf("%s", &player2Name)

		fmt.Printf("Hello %s!\n", player2Name)
	} else {
		//TODO Choose difficulty of CPU
		player2Name = "CPU"
	}

	// Player One's Sign

	var player1Sign string

	for {
		fmt.Printf("Insert Player One's sign: ")
		fmt.Scanf("%s", &player1Sign)

		if len(player1Sign) != 1 {
			fmt.Println("Please insert one character.")
		} else {
			break
		}
	}

	fmt.Printf("%s, your sign is %s!\n", player1Name, player1Sign)

	// Player Two's Sign

	var player2Sign string

	for {
		fmt.Printf("Insert Player Two's sign: ")
		fmt.Scanf("%s", &player2Sign)

		if len(player2Sign) != 1 {
			fmt.Println("Please insert one character.")
		} else {
			break
		}
	}

	fmt.Printf("%s, your sign is %s!\n", player2Name, player2Sign)

	// Instantiate game controller with configuration

	var p2ControllerType player_controller.PlayerType

	if playerNum == 2 {
		p2ControllerType = player_controller.CLIPLAYER
	} else {
		p2ControllerType = player_controller.RANDOMPLAYER
	}

	gameController := game_controller.NewGameController(player_controller.CLIPLAYER, p2ControllerType, player1Sign, player2Sign, player1Name, player2Name)

	gameController.ApplicationLoop()

	fmt.Println("It was fun playing with you! Bye!!")
}
