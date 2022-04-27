package player_controller

import (
	"math/rand"
	"tic-tac-go/game/board"
	"time"
)

type RandomPlayerController struct{}

func NewRandomPLayerController() RandomPlayerController {
	rand.Seed(time.Now().UnixNano())
	return RandomPlayerController{}
}

func (r RandomPlayerController) GetMove() board.Spot {
	userMove := rand.Intn(10)

	return board.Spot{
		Row: uint8(userMove) / 3,
		Col: uint8(userMove) % 3,
	}
}

// Useless methods for a random player
func (r RandomPlayerController) DisplayBoard()              {}
func (r RandomPlayerController) DisplayError(err string)    {}
func (r RandomPlayerController) DisplayWin(thisPlayer bool) {}
func (r RandomPlayerController) DisplayDraw()               {}
func (r RandomPlayerController) PlayAgain() bool            { return true }
