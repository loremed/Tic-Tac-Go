package board

import (
	"testing"
)

func TestGameStruct_GetSpot(t *testing.T) {
	type fields struct {
		board      Board
		moveCount  uint8
		conditions [8]int8
	}
	type args struct {
		spot Spot
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Sign
	}{
		{
			name: "0;0 PLAYER_ONE",
			fields: fields{
				board: Board{
					{PLAYER_ONE, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
				},
				moveCount:  0,
				conditions: [8]int8{0, 0, 0, 0, 0, 0, 0, 0},
			},
			args: args{
				spot: Spot{Row: 0, Col: 0},
			},
			want: PLAYER_ONE,
		},
		{
			name: "0;0 PLAYER_TWO",
			fields: fields{
				board: Board{
					{PLAYER_TWO, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
				},
				moveCount:  0,
				conditions: [8]int8{0, 0, 0, 0, 0, 0, 0, 0},
			},
			args: args{
				spot: Spot{Row: 0, Col: 0},
			},
			want: PLAYER_TWO,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board:      tt.fields.board,
				moveCount:  tt.fields.moveCount,
				conditions: tt.fields.conditions,
			}
			if got := g.GetSpot(tt.args.spot); got != tt.want {
				t.Errorf("GameStruct.GetSpot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameStruct_SetSpot(t *testing.T) {
	type fields struct {
		board      Board
		moveCount  uint8
		conditions [8]int8
	}
	type args struct {
		sign Sign
		spot Spot
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Sign
		wantErr bool
	}{
		{
			name: "0;0 PLAYER_ONE",
			fields: fields{
				board: Board{
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
				},
				moveCount:  0,
				conditions: [8]int8{0, 0, 0, 0, 0, 0, 0, 0},
			},
			args: args{
				sign: PLAYER_ONE,
				spot: Spot{Row: 0, Col: 0},
			},
			want:    EMPTY,
			wantErr: false,
		},
		{
			name: "0;0 ERROR",
			fields: fields{
				board: Board{
					{PLAYER_TWO, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
				},
				moveCount:  0,
				conditions: [8]int8{0, 0, 0, 0, 0, 0, 0, 0},
			},
			args: args{
				sign: PLAYER_ONE,
				spot: Spot{Row: 0, Col: 0},
			},
			want:    EMPTY,
			wantErr: true,
		},
		{
			name: "0;0 PLAYER_ONE WINNING",
			fields: fields{
				board: Board{
					{EMPTY, PLAYER_ONE, PLAYER_ONE},
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
				},
				moveCount:  0,
				conditions: [8]int8{2, 0, 0, 0, 0, 0, 0, 0},
			},
			args: args{
				sign: PLAYER_ONE,
				spot: Spot{Row: 0, Col: 0},
			},
			want:    PLAYER_ONE,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board:      tt.fields.board,
				moveCount:  tt.fields.moveCount,
				conditions: tt.fields.conditions,
			}
			got, err := g.SetSpot(tt.args.sign, tt.args.spot)
			if (err != nil) != tt.wantErr {
				t.Errorf("GameStruct.SetSpot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GameStruct.SetSpot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameStruct_IsWinning(t *testing.T) {
	type fields struct {
		board      Board
		moveCount  uint8
		conditions [8]int8
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "not winning",
			fields: fields{
				board: Board{
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
				},
				moveCount:  0,
				conditions: [8]int8{0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: false,
		},
		{
			name: "Winning PLAYER_ONE",
			fields: fields{
				board: Board{
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
				},
				moveCount:  0,
				conditions: [8]int8{0, 0, 0, 0, 0, 0, 3, 0},
			},
			want: true,
		},
		{
			name: "Winning PLAYER_TWO",
			fields: fields{
				board: Board{
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
				},
				moveCount:  0,
				conditions: [8]int8{-3, 0, 0, 0, 0, 0, 0, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board:      tt.fields.board,
				moveCount:  tt.fields.moveCount,
				conditions: tt.fields.conditions,
			}
			if got := g.IsWinning(); got != tt.want {
				t.Errorf("GameStruct.IsWinning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameStruct_IsFull(t *testing.T) {
	type fields struct {
		board      Board
		moveCount  uint8
		conditions [8]int8
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Not full board",
			fields: fields{
				board: Board{
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
				},
				moveCount:  3,
				conditions: [8]int8{3, 0, 0, 0, 0, 0, 0, 0},
			},
			want: false,
		},
		{
			name: "Full board",
			fields: fields{
				board: Board{
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
					{EMPTY, EMPTY, EMPTY},
				},
				moveCount:  9,
				conditions: [8]int8{3, 0, 0, 0, 0, 0, 0, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board:      tt.fields.board,
				moveCount:  tt.fields.moveCount,
				conditions: tt.fields.conditions,
			}
			if got := g.IsFull(); got != tt.want {
				t.Errorf("GameStruct.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameStruct_ResetGame(t *testing.T) {
	type fields struct {
		board      Board
		moveCount  uint8
		conditions [8]int8
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Reset Board Test",
			fields: fields{
				board: Board{
					{EMPTY, PLAYER_ONE, EMPTY},
					{EMPTY, EMPTY, PLAYER_TWO},
					{PLAYER_ONE, EMPTY, EMPTY},
				},
				moveCount:  12,
				conditions: [8]int8{1, 2, 3, 4, 5, 6, 7, 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				board:      tt.fields.board,
				moveCount:  tt.fields.moveCount,
				conditions: tt.fields.conditions,
			}
			g.ResetGame()
		})
	}
}
