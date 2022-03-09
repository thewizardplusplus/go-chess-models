package chessmodels

import (
	"reflect"
	"testing"
)

type MockBasePieceStorage struct {
	size Size

	piece func(position Position) (piece Piece, ok bool)
}

func (storage MockBasePieceStorage) Size() Size {
	return storage.size
}

func (storage MockBasePieceStorage) Piece(
	position Position,
) (piece Piece, ok bool) {
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

func TestDefaultBoardWrapperPieces(test *testing.T) {
	type fields struct {
		BasePieceStorage BasePieceStorage
	}
	type data struct {
		fields fields
		want   []Piece
	}

	for _, data := range []data{
		{
			fields: fields{
				BasePieceStorage: MockBasePieceStorage{
					size: Size{5, 5},
					piece: func(position Position) (piece Piece, ok bool) {
						if position != (Position{2, 3}) && position != (Position{4, 2}) {
							return nil, false
						}

						piece = MockPiece{position: position}
						return piece, true
					},
				},
			},
			want: []Piece{
				MockPiece{position: Position{4, 2}},
				MockPiece{position: Position{2, 3}},
			},
		},
		{
			fields: fields{
				BasePieceStorage: struct {
					MockBasePieceStorage
					MockPieceGroupGetter
				}{
					MockPieceGroupGetter: MockPieceGroupGetter{
						pieces: []Piece{
							MockPiece{position: Position{4, 2}},
							MockPiece{position: Position{2, 3}},
						},
					},
				},
			},
			want: []Piece{
				MockPiece{position: Position{4, 2}},
				MockPiece{position: Position{2, 3}},
			},
		},
	} {
		board := DefaultBoardWrapper{
			BasePieceStorage: data.fields.BasePieceStorage,
		}
		got := board.Pieces()

		if !reflect.DeepEqual(got, data.want) {
			test.Fail()
		}
	}
}

func TestDefaultBoardWrapperCheckMove(test *testing.T) {
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
		baseStorage := MockBasePieceStorage{
			size:  data.fields.size,
			piece: data.fields.piece,
		}
		board := DefaultBoardWrapper{baseStorage}
		got := board.CheckMove(data.args.move)

		if got != data.want {
			test.Fail()
		}
	}
}
