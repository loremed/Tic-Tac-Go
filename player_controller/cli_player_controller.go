package player_controller

import (
	"fmt"
	"tic-tac-go/game/board"
)

// Implements PlayerController interface. Specific for a CLI bound player.

type CLIPlayerController struct {
	b  *board.Board
	p1 string
	p2 string
}

func NewCLIPlayerController(boardToInterface *board.Board, p1Char string, p2Char string) CLIPlayerController {
	return CLIPlayerController{
		b:  boardToInterface,
		p1: p1Char,
		p2: p2Char,
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
		return c.p1
	} else if s == board.PLAYER_TWO {
		return c.p2
	} else {
		return " "
	}
}

func (c CLIPlayerController) GetMove() board.Spot {
	c.DisplayBoard()
	fmt.Print("Choose where to place your sign: ")

	var userMove int

	for {
		n, err := fmt.Scanf("%d", &userMove)

		if n == 1 {
			break
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
	fmt.Printf("Error: %s\n", err)
}

func (c CLIPlayerController) DisplayWin(winner string, thisPlayer bool) {
	fmt.Printf("The winner is %s\n", winner)
	if thisPlayer {
		fmt.Println("You won!! Congratulations!")
	} else {
		fmt.Println("You lost! Better luck next time :(")
	}
}

func (c CLIPlayerController) DisplayDraw() {
	fmt.Println("The gme ended in a draw!")
}
