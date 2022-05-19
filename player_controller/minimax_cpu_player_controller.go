package player_controller

import (
	"math"
	"tic-tac-go/game/board"
)

type MinimaxPlayerController struct {
	game       *board.Game
	thisPlayer board.Sign
}

func NewMinimaxPLayerController(gameToInterface *board.Game, thisPlayerPar board.Sign) MinimaxPlayerController {
	return MinimaxPlayerController{
		game:       gameToInterface,
		thisPlayer: thisPlayerPar,
	}
}

func (r MinimaxPlayerController) GetMove() board.Spot {

	var bestScore int = int(math.Inf(-1))
	var bestMove board.Spot = board.Spot{
		Row: uint8(math.Inf(-1)),
		Col: uint8(math.Inf(-1)),
	}

	for _, spot := range r.game.GetPossibleMoves() {
		r.game.SetSpot(r.thisPlayer, spot)
		score := r.Minimax(r.game, 1, r.thisPlayer.OtherPlayer())
		r.game.UndoMove(spot)

		if score > bestScore {
			bestScore = score
			bestMove = spot
		}
	}

	return bestMove
}

func (r MinimaxPlayerController) Minimax(game *board.Game, depth int, movingPlayer board.Sign) int {

	if game.IsOver() {
		return r.ScoreGame(*game, depth)
	}

	scores := []int{}

	for _, spot := range game.GetPossibleMoves() {
		game.SetSpot(movingPlayer, spot)
		scores = append(scores, r.Minimax(game, depth+1, movingPlayer.OtherPlayer()))
		game.UndoMove(spot)
	}

	min_index, max_index := findMinAndMaxIndex(scores)

	if movingPlayer == r.thisPlayer {
		return scores[max_index]
	} else {
		return scores[min_index]
	}
}

func (r MinimaxPlayerController) ScoreGame(game board.Game, depth int) int {
	winner := game.Winner()

	switch winner {
	case r.thisPlayer:
		return 10 - depth

	case r.thisPlayer.OtherPlayer():
		return depth - 10

	case board.EMPTY:
		return 0
	}

	return 0
}

// Useless methods for a Minimax player
func (r MinimaxPlayerController) DisplayBoard()              {}
func (r MinimaxPlayerController) DisplayError(err string)    {}
func (r MinimaxPlayerController) DisplayWin(thisPlayer bool) {}
func (r MinimaxPlayerController) DisplayDraw()               {}
func (r MinimaxPlayerController) PlayAgain() bool            { return true }

func findMinAndMaxIndex(a []int) (minIndex int, maxIndex int) {
	minIndex = 0
	maxIndex = 0
	minValue := a[0]
	maxValue := a[0]
	for i, value := range a {
		if value < minValue {
			minIndex = i
		}
		if value > maxValue {
			maxIndex = i
		}
	}
	return minIndex, maxIndex
}
