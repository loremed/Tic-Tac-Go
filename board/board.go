package board

import (
	"errors"
)

// Types used

type Sign int8

type Spot struct {
	Row uint8
	Col uint8
}

type Board [3][3]Sign

// Uppercase (exported) for Minimax tests
type Game struct {
	Board      Board
	MoveCount  uint8
	Conditions [8]int8
}

// Constants

const EMPTY = Sign(0)
const PLAYER_ONE = Sign(1)
const PLAYER_TWO = Sign(-1)

const _ROW = 0
const _COL = 3
const _DIAG = 6
const _ANTIDIAG = 7

// Game Contructor

func NewGame() *Game {
	return &Game{
		Board: Board{
			{EMPTY, EMPTY, EMPTY},
			{EMPTY, EMPTY, EMPTY},
			{EMPTY, EMPTY, EMPTY},
		},
		MoveCount:  0,
		Conditions: [8]int8{0, 0, 0, 0, 0, 0, 0, 0},
	}
}

// Gets a spot from the board
func (g *Game) GetSpot(spot Spot) Sign {
	return g.Board[spot.Row][spot.Col]
}

// Places an X or an O in the spot asked, if it is free

func (g *Game) SetSpot(sign Sign, spot Spot) (Sign, error) {

	// Exit if sign is not valid

	if sign != PLAYER_ONE && sign != PLAYER_TWO && sign != EMPTY {
		// fmt.Println("DEBUG:invalid sign")
		return EMPTY, errors.New("invalid sign")
	}

	// Exit if spot is not EMPTY

	if g.GetSpot(spot) != EMPTY && sign != EMPTY {
		//fmt.Println("DEBUG:spot taken")
		return EMPTY, errors.New("spot is already taken")
	}

	g.Board[spot.Row][spot.Col] = sign

	// update Conditions array and MoveCount

	g.MoveCount++
	g.Conditions[spot.Row+_ROW] += int8(sign)
	g.Conditions[spot.Col+_COL] += int8(sign)

	if spot.Row == spot.Col {
		g.Conditions[_DIAG] += int8(sign)
	}

	if spot.Row+spot.Col == 2 {
		g.Conditions[_ANTIDIAG] += int8(sign)
	}

	// Check if board is winning after current move

	resultSign := EMPTY
	if g.IsWinning() {
		resultSign = sign
	}

	return resultSign, nil
}

func (g *Game) UndoMove(spot Spot) {

	oldSign := g.Board[spot.Row][spot.Col]

	// update Conditions array and MoveCount

	g.MoveCount--
	g.Conditions[spot.Row+_ROW] -= int8(oldSign)
	g.Conditions[spot.Col+_COL] -= int8(oldSign)

	if spot.Row == spot.Col {
		g.Conditions[_DIAG] -= int8(oldSign)
	}

	if spot.Row+spot.Col == 2 {
		g.Conditions[_ANTIDIAG] -= int8(oldSign)
	}

	g.Board[spot.Row][spot.Col] = EMPTY
}

// Checks if board is winning using the Conditions array

func (g *Game) IsWinning() bool {
	return g.Winner() != EMPTY
}

// Checks if there is place on the board

func (g *Game) IsFull() bool {
	return g.MoveCount == 9
}

func (g *Game) GetBoard() *Board {
	return &g.Board
}

func (g *Game) Winner() Sign {

	for _, value := range g.Conditions {
		if value == 3 {
			return PLAYER_ONE
		}
		if value == -3 {
			return PLAYER_TWO
		}
	}
	return EMPTY
}

func (g *Game) GetPossibleMoves() (spots []Spot) {
	moves := []Spot{}
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if g.Board[row][col] == EMPTY {
				moves = append(moves, Spot{
					Row: uint8(row),
					Col: uint8(col),
				})
			}
		}
	}

	/* for _, j := range moves {
		fmt.Printf("%d %d; ", j.Row, j.Col)
	} */

	return moves
}

func (g *Game) IsOver() bool {
	// fmt.Printf("%t %t\n", g.IsFull(), g.IsWinning())
	return g.IsFull() || g.IsWinning()
}

// Probably unused

func (g *Game) ResetGame() {
	g.Board = [3][3]Sign{
		{EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY},
	}

	g.MoveCount = 0

	g.Conditions = [8]int8{0, 0, 0, 0, 0, 0, 0, 0}
}

func (s Sign) OtherPlayer() Sign {
	return Sign(int8(s) * -1)
}
