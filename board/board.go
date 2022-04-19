package board

import (
	"errors"
)

// Types used

type Sign int8

type Spot struct {
	row uint8
	col uint8
}

type Board [3][3]Sign

// Constants

const EMPTY = Sign(0)
const X = Sign(1)
const O = Sign(-1)

const _ROW = 0
const _COL = 3
const _DIAG = 6
const _ANTIDIAG = 7

// Variables

var board Board
var moveCount uint8
var conditions [8]int8

// Init

func init() {
	board = [3][3]Sign{
		{EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY},
	}

	moveCount = 0

	conditions = [8]int8{0, 0, 0, 0, 0, 0, 0, 0}
}

// Gets a spot from the board
func GetSpot(spot Spot) Sign {
	return board[spot.row][spot.col]
}

// Places an X or an O in the spot asked, if it is free

func SetSpot(sign Sign, spot Spot) (Sign, error) {

	// Exit if sign is not valid

	if sign != X && sign != O {
		return EMPTY, errors.New("invalid sign")
	}

	// Exit if spot is not EMPTY

	if GetSpot(spot) != EMPTY {
		return EMPTY, errors.New("Spot is already taken")
	}

	board[spot.row][spot.col] = sign

	// update conditions array and moveCount

	moveCount++

	conditions[spot.row+_ROW] += int8(sign)
	conditions[spot.col+_COL] += int8(sign)

	if spot.row == spot.col {
		conditions[_DIAG] += int8(sign)
	}

	if spot.row+spot.col == 2 {
		conditions[_ANTIDIAG] += int8(sign)
	}

	// Check if board is winning after current move

	isWinning := IsWinning()

	resultSign := EMPTY
	if isWinning {
		resultSign = sign
	}

	return resultSign, nil
}

// Checks if board is winning using the conditions array

func IsWinning() bool {
	for _, value := range conditions {
		if value == 3 || value == -3 {
			return true
		}
	}
	return false
}

// Checks if there is place on the board

func IsFull() bool {
	return moveCount == 9
}

func GetBoard() Board {
	result := board
	return result
}

func String(sign Sign) string {
	if sign == EMPTY {
		return " "
	} else if sign == X {
		return "X"
	} else {
		return "O"
	}
}
