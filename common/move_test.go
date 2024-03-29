package common

import (
	"testing"
)

type MockPiece struct {
	kind     Kind
	color    Color
	position Position

	checkMove func(move Move, storage PieceStorage) bool
}

func (piece MockPiece) Kind() Kind {
	return piece.kind
}

func (piece MockPiece) Color() Color {
	return piece.color
}

func (piece MockPiece) Position() Position {
	return piece.position
}

func (piece MockPiece) ApplyPosition(position Position) Piece {
	return MockPiece{
		kind:      piece.kind,
		color:     piece.color,
		position:  position,
		checkMove: piece.checkMove,
	}
}

func (piece MockPiece) CheckMove(move Move, storage PieceStorage) bool {
	if piece.checkMove == nil {
		panic("not implemented")
	}

	return piece.checkMove(move, storage)
}

type MockBasePieceStorage struct {
	size Size

	piece func(position Position) (piece Piece, ok bool)
}

func (storage MockBasePieceStorage) Size() Size {
	return storage.size
}

func (storage MockBasePieceStorage) Piece(position Position) (
	piece Piece,
	ok bool,
) {
	if storage.piece == nil {
		panic("not implemented")
	}

	return storage.piece(position)
}

func (storage MockBasePieceStorage) ApplyMove(move Move) PieceStorage {
	panic("not implemented")
}

type MockPieceGroupGetter struct {
	pieces []Piece
}

func (pieceGroupGetter MockPieceGroupGetter) Pieces() []Piece {
	return pieceGroupGetter.pieces
}

type MockMoveChecker struct {
	checkMove func(move Move) error
}

func (moveChecker MockMoveChecker) CheckMove(move Move) error {
	if moveChecker.checkMove == nil {
		panic("not implemented")
	}

	return moveChecker.checkMove(move)
}

type MockPieceStorage struct {
	MockBasePieceStorage
	MockPieceGroupGetter
	MockMoveChecker
}

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
		size  Size
		piece func(position Position) (piece Piece, ok bool)
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
				size: Size{2, 2},
				piece: func(position Position) (piece Piece, ok bool) {
					return nil, false
				},
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
				size: Size{2, 2},
				piece: func(position Position) (piece Piece, ok bool) {
					return nil, false
				},
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
				size: Size{2, 2},
				piece: func(position Position) (piece Piece, ok bool) {
					return nil, false
				},
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
				piece: func(position Position) (piece Piece, ok bool) {
					if position != (Position{0, 0}) && position != (Position{1, 1}) {
						return nil, false
					}

					piece = MockPiece{color: Black, position: position}
					return piece, true
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
				piece: func(position Position) (piece Piece, ok bool) {
					if position != (Position{0, 0}) {
						return nil, false
					}

					piece = MockPiece{
						position: position,
						checkMove: func(move Move, storage PieceStorage) bool {
							return false
						},
					}
					return piece, true
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
				piece: func(position Position) (piece Piece, ok bool) {
					switch position {
					case Position{0, 0}:
						piece = MockPiece{
							color:    Black,
							position: Position{0, 0},
							checkMove: func(move Move, storage PieceStorage) bool {
								return true
							},
						}
					case Position{1, 1}:
						piece = MockPiece{
							kind:     King,
							color:    White,
							position: Position{1, 1},
						}
					}

					ok = piece != nil
					return piece, ok
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
				piece: func(position Position) (piece Piece, ok bool) {
					if position != (Position{0, 0}) {
						return nil, false
					}

					piece = MockPiece{
						position: position,
						checkMove: func(move Move, storage PieceStorage) bool {
							return true
						},
					}
					return piece, true
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
