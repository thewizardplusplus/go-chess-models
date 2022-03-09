package chessmodels

import (
	"errors"
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

type MockBasePieceStorage struct {
	size Size

	piece func(position common.Position) (piece Piece, ok bool)
}

func (storage MockBasePieceStorage) Size() Size {
	return storage.size
}

func (storage MockBasePieceStorage) Piece(
	position common.Position,
) (piece Piece, ok bool) {
	if storage.piece == nil {
		panic("not implemented")
	}

	return storage.piece(position)
}

func (storage MockBasePieceStorage) ApplyMove(move common.Move) PieceStorage {
	panic("not implemented")
}

type MockPieceGroupGetter struct {
	pieces []Piece
}

func (pieceGroupGetter MockPieceGroupGetter) Pieces() []Piece {
	return pieceGroupGetter.pieces
}

type MockMoveChecker struct {
	checkMove func(move common.Move) error
}

func (moveChecker MockMoveChecker) CheckMove(move common.Move) error {
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
					piece: func(position common.Position) (piece Piece, ok bool) {
						if position != (common.Position{2, 3}) && position != (common.Position{4, 2}) {
							return nil, false
						}

						piece = MockPiece{position: position}
						return piece, true
					},
				},
			},
			want: []Piece{
				MockPiece{position: common.Position{4, 2}},
				MockPiece{position: common.Position{2, 3}},
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
							MockPiece{position: common.Position{4, 2}},
							MockPiece{position: common.Position{2, 3}},
						},
					},
				},
			},
			want: []Piece{
				MockPiece{position: common.Position{4, 2}},
				MockPiece{position: common.Position{2, 3}},
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

func TestDefaultBoardWrapperCheckMove_withMoveCheckerInterface(
	test *testing.T,
) {
	type fields struct {
		BasePieceStorage BasePieceStorage
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
				BasePieceStorage: struct {
					MockBasePieceStorage
					MockMoveChecker
				}{
					MockMoveChecker: MockMoveChecker{
						checkMove: func(move common.Move) error {
							if !reflect.DeepEqual(move, common.Move{
								Start:  common.Position{0, 0},
								Finish: common.Position{1, 1},
							}) {
								test.Fail()
							}

							return errors.New("dummy")
						},
					},
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{1, 1},
				},
			},
			want: errors.New("dummy"),
		},
		{
			fields: fields{
				BasePieceStorage: struct {
					MockBasePieceStorage
					MockMoveChecker
				}{
					MockMoveChecker: MockMoveChecker{
						checkMove: func(move common.Move) error {
							if !reflect.DeepEqual(move, common.Move{
								Start:  common.Position{0, 0},
								Finish: common.Position{1, 1},
							}) {
								test.Fail()
							}

							return nil
						},
					},
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
		board := DefaultBoardWrapper{
			BasePieceStorage: data.fields.BasePieceStorage,
		}
		got := board.CheckMove(data.args.move)

		if !reflect.DeepEqual(got, data.want) {
			test.Fail()
		}
	}
}
