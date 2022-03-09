package chessmodels

import (
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

func TestCheckMove(test *testing.T) {
	type fields struct {
		size  Size
		piece func(position common.Position) (piece Piece, ok bool)
	}
	type args struct {
		move common.Move
	}
	type data struct {
		fields fields
		args   args
		want   error
	}

	for _, data := range []data{
		{
			fields: fields{
				size: Size{2, 2},
				piece: func(position common.Position) (piece Piece, ok bool) {
					return nil, false
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{0, 0},
				},
			},
			want: ErrNoMove,
		},
		{
			fields: fields{
				size: Size{2, 2},
				piece: func(position common.Position) (piece Piece, ok bool) {
					return nil, false
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{-1, -1},
				},
			},
			want: ErrOutOfSize,
		},
		{
			fields: fields{
				size: Size{2, 2},
				piece: func(position common.Position) (piece Piece, ok bool) {
					return nil, false
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{1, 1},
				},
			},
			want: ErrNoPiece,
		},
		{
			fields: fields{
				size: Size{2, 2},
				piece: func(position common.Position) (piece Piece, ok bool) {
					if position != (common.Position{0, 0}) && position != (common.Position{1, 1}) {
						return nil, false
					}

					piece = MockPiece{color: common.Black, position: position}
					return piece, true
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{1, 1},
				},
			},
			want: ErrFriendlyTarget,
		},
		{
			fields: fields{
				size: Size{2, 2},
				piece: func(position common.Position) (piece Piece, ok bool) {
					if position != (common.Position{0, 0}) {
						return nil, false
					}

					piece = MockPiece{
						position: position,
						checkMove: func(move common.Move, storage PieceStorage) bool {
							return false
						},
					}
					return piece, true
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{1, 1},
				},
			},
			want: ErrIllegalMove,
		},
		{
			fields: fields{
				size: Size{2, 2},
				piece: func(position common.Position) (piece Piece, ok bool) {
					switch position {
					case common.Position{0, 0}:
						piece = MockPiece{
							color:    common.Black,
							position: common.Position{0, 0},
							checkMove: func(move common.Move, storage PieceStorage) bool {
								return true
							},
						}
					case common.Position{1, 1}:
						piece = MockPiece{
							kind:     common.King,
							color:    common.White,
							position: common.Position{1, 1},
						}
					}

					ok = piece != nil
					return piece, ok
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{1, 1},
				},
			},
			want: ErrKingCapture,
		},
		{
			fields: fields{
				size: Size{2, 2},
				piece: func(position common.Position) (piece Piece, ok bool) {
					if position != (common.Position{0, 0}) {
						return nil, false
					}

					piece = MockPiece{
						position: position,
						checkMove: func(move common.Move, storage PieceStorage) bool {
							return true
						},
					}
					return piece, true
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{1, 1},
				},
			},
			want: nil,
		},
	} {
		storage := MockPieceStorage{
			MockBasePieceStorage: MockBasePieceStorage{
				size:  data.fields.size,
				piece: data.fields.piece,
			},
		}
		got := CheckMove(storage, data.args.move)

		if got != data.want {
			test.Fail()
		}
	}
}
