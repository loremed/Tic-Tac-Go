package player_controller

import (
	"reflect"
	"testing"
	"tic-tac-go/game/board"
)

func TestMinimaxPlayerController_GetMove(t *testing.T) {
	type fields struct {
		game       *board.Game
		thisPlayer board.Sign
	}
	tests := []struct {
		name   string
		fields fields
		want   board.Spot
	}{
		{
			name: "first test",
			fields: fields{
				game: &board.Game{
					Board: board.Board{
						{board.EMPTY, board.EMPTY, board.EMPTY},
						{board.PLAYER_ONE, board.PLAYER_ONE, board.EMPTY},
						{board.EMPTY, board.EMPTY, board.EMPTY},
					},
					MoveCount:  2,
					Conditions: [8]int8{0, 2, 0, 1, 1, 1, 1, 1},
				},
				thisPlayer: board.PLAYER_TWO,
			},
			want: board.Spot{
				Row: 1,
				Col: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := MinimaxPlayerController{
				game:       tt.fields.game,
				thisPlayer: tt.fields.thisPlayer,
			}
			if got := r.GetMove(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinimaxPlayerController.GetMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
