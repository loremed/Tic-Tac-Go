package player_controller

import (
	"bufio"
	"fmt"
	"os"
	"tic-tac-go/game/board"
)

// Implements PlayerController interface. Specific for a CLI bound player.

type CLIPlayerController struct {
	b               *board.Board
	p1Sign          string
	p2Sign          string
	thisPlayerName  string
	otherPlayerName string
}

func NewCLIPlayerController(boardToInterface *board.Board, p1Char string, p2Char string, thisPlayerNameArg string, otherPlayerNameArg string) CLIPlayerController {
	return CLIPlayerController{
		b:               boardToInterface,
		p1Sign:          p1Char,
		p2Sign:          p2Char,
		thisPlayerName:  thisPlayerNameArg,
		otherPlayerName: otherPlayerNameArg,
	}
}

func (c CLIPlayerController) DisplayBoard() {
	fmt.Printf("\n %s | %s | %s \n", c.toString(c.b[0][0]), c.toString(c.b[0][1]), c.toString(c.b[0][2]))
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n", c.toString(c.b[1][0]), c.toString(c.b[1][1]), c.toString(c.b[1][2]))
	fmt.Println("-----------")
	fmt.Printf(" %s | %s | %s \n\n", c.toString(c.b[2][0]), c.toString(c.b[2][1]), c.toString(c.b[2][2]))

}

func (c CLIPlayerController) toString(s board.Sign) string {
	if s == board.PLAYER_ONE {
		return c.p1Sign
	} else if s == board.PLAYER_TWO {
		return c.p2Sign
	} else {
		return " "
	}
}

func (c CLIPlayerController) GetMove() board.Spot {
	// c.DisplayBoard()

	var userMove int
	stdin := bufio.NewReader(os.Stdin)

	for {

		fmt.Printf("%s, choose where to place your sign [1-9]: ", c.thisPlayerName)
		_, err := fmt.Fscanf(stdin, "%d\n", &userMove)

		if err != nil { // error
			stdin.ReadString('\n')
			fmt.Println("Number is not valid. Try again.")
		} else if userMove < 1 || userMove > 9 { // semantics
			// stdin.ReadString('\n')
			fmt.Println("Move number is not in range [1-9]. Try again.")
		} else {
			break
		}

	}

	userMove--

	return board.Spot{
		Row: uint8(userMove) / 3,
		Col: uint8(userMove) % 3,
	}

}

func (c CLIPlayerController) DisplayError(err string) {
	fmt.Printf("%s, error: %s\n", c.thisPlayerName, err)
}

func (c CLIPlayerController) DisplayWin(thisPlayer bool) {

	c.DisplayBoard()

	if thisPlayer {
		fmt.Printf("\nThe winner is %s\n", c.thisPlayerName)
		fmt.Printf("%s, you won!! Congratulations!\n\n", c.thisPlayerName)
	} else {
		fmt.Printf("\nThe winner is %s\n", c.otherPlayerName)
		fmt.Printf("%s, you lost! Better luck next time :(\n\n", c.thisPlayerName)
	}
}

func (c CLIPlayerController) DisplayDraw() {
	c.DisplayBoard()

	fmt.Println("\nThe game ended in a draw!")
}

func (c CLIPlayerController) PlayAgain() bool {

	var answer int
	stdin := bufio.NewReader(os.Stdin)

	for {

		fmt.Printf("%s, do you want to play again?\n1: Yes\n2: No\nYour choice [1-2]: ", c.thisPlayerName)
		_, err := fmt.Fscanf(stdin, "%d\n", &answer)

		if err != nil { // error
			stdin.ReadString('\n')
			fmt.Println("Number is not valid. Try again.")
		} else if answer < 1 || answer > 2 { // semantics
			// stdin.ReadString('\n')
			fmt.Println("Option number is not in range [1-2]. Try again.")
		} else {
			break
		}

	}

	return answer == 1

}
