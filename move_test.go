package chessmodels

import (
	"testing"
)

func TestMoveIsZero(test *testing.T) {
	type fields struct {
		start  Position
		finish Position
	}
	type data struct {
		fields fields
		want   bool
	}

	for _, data := range []data{
		{
			fields: fields{
				start:  Position{1, 2},
				finish: Position{3, 4},
			},
			want: false,
		},
		{
			fields: fields{
				start:  Position{1, 0},
				finish: Position{3, 0},
			},
			want: false,
		},
		{
			fields: fields{
				start:  Position{1, 2},
				finish: Position{0, 0},
			},
			want: false,
		},
		{
			fields: fields{
				start:  Position{0, 0},
				finish: Position{0, 0},
			},
			want: true,
		},
	} {
		move := Move{
			Start:  data.fields.start,
			Finish: data.fields.finish,
		}
		got := move.IsZero()

		if got != data.want {
			test.Fail()
		}
	}
}

func TestMoveIsEmpty(test *testing.T) {
	type fields struct {
		start  Position
		finish Position
	}
	type data struct {
		fields fields
		want   bool
	}

	for _, data := range []data{
		{
			fields: fields{
				start:  Position{1, 2},
				finish: Position{3, 4},
			},
			want: false,
		},
		{
			fields: fields{
				start:  Position{1, 2},
				finish: Position{1, 4},
			},
			want: false,
		},
		{
			fields: fields{
				start:  Position{1, 2},
				finish: Position{3, 2},
			},
			want: false,
		},
		{
			fields: fields{
				start:  Position{1, 2},
				finish: Position{1, 2},
			},
			want: true,
		},
	} {
		move := Move{
			Start:  data.fields.start,
			Finish: data.fields.finish,
		}
		got := move.IsEmpty()

		if got != data.want {
			test.Fail()
		}
	}
}

func TestCheckMove(test *testing.T) {
	type fields struct {
		size   Size
		pieces pieceGroup
	}
	type args struct {
		move Move
	}
	type data struct {
		fields fields
		args   args
		want   error
	}

	for _, data := range []data{
		{
			fields: fields{
				size:   Size{2, 2},
				pieces: nil,
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{0, 0},
				},
			},
			want: ErrNoMove,
		},
		{
			fields: fields{
				size:   Size{2, 2},
				pieces: nil,
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{-1, -1},
				},
			},
			want: ErrOutOfSize,
		},
		{
			fields: fields{
				size:   Size{2, 2},
				pieces: nil,
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{1, 1},
				},
			},
			want: ErrNoPiece,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: pieceGroup{
					Position{0, 0}: MockPiece{
						color:    Black,
						position: Position{0, 0},
					},
					Position{1, 1}: MockPiece{
						color:    Black,
						position: Position{1, 1},
					},
				},
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{1, 1},
				},
			},
			want: ErrFriendlyTarget,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: pieceGroup{
					Position{0, 0}: MockPiece{
						position: Position{0, 0},
						checkMove: func(move Move, storage PieceStorage) bool {
							return false
						},
					},
				},
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{1, 1},
				},
			},
			want: ErrIllegalMove,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: pieceGroup{
					Position{0, 0}: MockPiece{
						color:    Black,
						position: Position{0, 0},
						checkMove: func(move Move, storage PieceStorage) bool {
							return true
						},
					},
					Position{1, 1}: MockPiece{
						kind:     King,
						color:    White,
						position: Position{1, 1},
					},
				},
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{1, 1},
				},
			},
			want: ErrKingCapture,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: pieceGroup{
					Position{0, 0}: MockPiece{
						position: Position{0, 0},
						checkMove: func(move Move, storage PieceStorage) bool {
							return true
						},
					},
				},
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{1, 1},
				},
			},
			want: nil,
		},
	} {
		board := MapBoard{
			BaseBoard: BaseBoard{
				size: data.fields.size,
			},

			pieces: data.fields.pieces,
		}
		got := CheckMove(board, data.args.move)

		if got != data.want {
			test.Fail()
		}
	}
}
