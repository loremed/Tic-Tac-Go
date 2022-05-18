package player_controller

import (
	"tic-tac-go/game/board"
)

// Typedef for PlayerType

type PlayerType uint8

// Definition of player controllers

const CLIPLAYER = PlayerType(0)
const RANDOMPLAYER = PlayerType(1)
const MINIMAXPLAYER = PlayerType(2)

// Intergace for a Generic Player Controller, implementable by CLI, GUI or CPU

type PlayerController interface {
	GetMove() board.Spot
	DisplayBoard()
	DisplayError(err string)
	DisplayWin(thisPlayer bool)
	DisplayDraw()
	PlayAgain() bool
}

func NewPlayerController(playerType PlayerType, boardToInterface *board.Game, p1Char string, p2Char string, p1Name string, p2Name string, thisPlayer board.Sign) PlayerController {
	switch playerType {
	case CLIPLAYER:
		return NewCLIPlayerController(boardToInterface.GetBoard(), p1Char, p2Char, p1Name, p2Name)
	case RANDOMPLAYER:
		return NewRandomPLayerController()
	case MINIMAXPLAYER:
		return NewMinimaxPLayerController(boardToInterface, thisPlayer)
	default:
		return NewRandomPLayerController()
	}
}
