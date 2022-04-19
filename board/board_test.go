package board

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpot(t *testing.T) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			assert.Equal(t, EMPTY, GetSpot(Spot{uint8(i), uint8(j)}), "Spot should be EMPTY")
		}
	}
}

func TestSetSpot(t *testing.T) {
	spot := Spot{0, 0}

	sign := X

	winnerSign, errorResult := SetSpot(sign, spot)

	assert.Equal(t, EMPTY, winnerSign, "there should be no winner")
	assert.Equal(t, sign, GetSpot(spot), "spot should be assigned correctly")
	assert.Nil(t, errorResult, "Error should be nil")

	winnerSign2, errorResult2 := SetSpot(sign, spot)

	assert.NotNil(t, errorResult2, "Error should not be nil when setting same spot")
	assert.Equal(t, winnerSign2, EMPTY, "there should be no winner when error")

}

func TestIsWinning(t *testing.T) {
	sign := O
	SetSpot(sign, Spot{2, 0})
	SetSpot(sign, Spot{1, 1})
	winnerSign, errorResult := SetSpot(sign, Spot{0, 2})

	result := IsWinning()

	assert.Equal(t, true, result, "board should be winning")
	assert.Equal(t, sign, winnerSign, "X should be the winner")
	assert.Nil(t, errorResult, "there should be no error")
}

func TestIsFull(t *testing.T) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			SetSpot(X, Spot{uint8(i), uint8(j)})
		}
	}

	assert.Equal(t, true, IsFull(), "board should be full")
}

func TestGetBoard(t *testing.T) {
	tests := []struct {
		name string
		want Board
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBoard(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
