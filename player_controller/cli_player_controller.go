package player_controller

import (
	"fmt"
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

func NewCLIPlayerController(boardToInterface *board.Board, p1Char string, p2Char string, name1 string, name2 string) CLIPlayerController {
	return CLIPlayerController{
		b:               boardToInterface,
		p1Sign:          p1Char,
		p2Sign:          p2Char,
		thisPlayerName:  name1,
		otherPlayerName: name2,
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
	c.DisplayBoard()

	fmt.Printf("%s, choose where to place your sign [1-9]: ", c.thisPlayerName)

	var userMove int

	for {
		n, err := fmt.Scanf("%d", &userMove)

		if n == 1 {
			if userMove > 0 && userMove < 10 {
				break
			} else {
				fmt.Println("Choose a number between 1 and 9!")
			}
		} else {
			fmt.Printf("Not a valid integer! %s", err.Error())
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
	if thisPlayer {
		fmt.Printf("The winner is %s\n", c.thisPlayerName)
		fmt.Printf("%s, you won!! Congratulations!\n\n", c.thisPlayerName)
	} else {
		fmt.Printf("The winner is %s\n", c.otherPlayerName)
		fmt.Printf("%s, you lost! Better luck next time :(\n\n", c.thisPlayerName)
	}
}

func (c CLIPlayerController) DisplayDraw() {
	fmt.Println("The game ended in a draw!")
}

func (c CLIPlayerController) PlayAgain() bool {

	fmt.Printf("%s, do you want to play again?\n1: Yes\n2: No\nYour choice [1-2]: ", c.thisPlayerName)

	var answer int

	fmt.Scanf("%d", &answer)

	return answer == 1

}
