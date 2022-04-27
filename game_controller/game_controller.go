package game_controller

import (
	"tic-tac-go/game/board"
	"tic-tac-go/game/player_controller"
)

type GameController struct {
	game    *board.Game
	player1 player_controller.PlayerController
	player2 player_controller.PlayerController
}

func NewGameController(playerType1 player_controller.PlayerType, playerType2 player_controller.PlayerType, player1Char string, player2char string) *GameController {
	game := board.NewGame()

	return &GameController{
		game:    game,
		player1: player_controller.NewPlayerController(playerType1, game.GetBoard(), player1Char, player2char),
		player2: player_controller.NewPlayerController(playerType2, game.GetBoard(), player1Char, player2char),
	}
}

func (c *GameController) ApplicationLoop() {

	for c.SingleGameLoop() {
	}

}

func (c *GameController) SingleGameLoop() (playAgain bool) {
	// main game loop
	for {

		// Player1 turn

		gameIsOver, playAgain := c.PlayTurn(c.player1, c.player2, board.PLAYER_ONE)

		if gameIsOver {
			return playAgain
		}

		// Player2 turn

		gameIsOver, playAgain = c.PlayTurn(c.player2, c.player1, board.PLAYER_TWO)

		if gameIsOver {
			return playAgain
		}
	}
}

func (c *GameController) manageWin(winner board.Sign) {
	switch winner {
	case board.PLAYER_ONE:
		c.player1.DisplayWin("Player One", true)
		c.player2.DisplayWin("Player One", false)
	case board.PLAYER_TWO:
		c.player1.DisplayWin("Player Two", false)
		c.player2.DisplayWin("Player Two", true)
	}
}

func (c *GameController) PlayTurn(thisPlayer player_controller.PlayerController,
	otherPlayer player_controller.PlayerController,
	thisPlayerSign board.Sign) (gameIsOver bool, anotherGame bool) {

	thisPlayer.DisplayBoard()

	var winner board.Sign
	var err error

	// player1 move loop
	for {
		player1Spot := thisPlayer.GetMove()

		// Try placing player1's sign on the board

		winner, err = c.game.SetSpot(thisPlayerSign, player1Spot)

		// Error management after player1's move

		if err != nil {
			thisPlayer.DisplayError(err.Error())
		} else { // succesful move
			break
		}
	}

	if winner != board.EMPTY {
		c.manageWin(board.PLAYER_ONE)
		return true, thisPlayer.PlayAgain() && otherPlayer.PlayAgain()
	} else if c.game.IsFull() {
		c.player1.DisplayDraw()
		c.player2.DisplayDraw()
		return true, thisPlayer.PlayAgain() && otherPlayer.PlayAgain()
	}

	return false, true
}
