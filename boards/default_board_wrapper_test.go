package boards

import (
	"errors"
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

type MockPiece struct {
	kind     common.Kind
	color    common.Color
	position common.Position

	checkMove func(move common.Move, storage common.PieceStorage) bool
}

func (piece MockPiece) Kind() common.Kind {
	return piece.kind
}

func (piece MockPiece) Color() common.Color {
	return piece.color
}

func (piece MockPiece) Position() common.Position {
	return piece.position
}

func (piece MockPiece) ApplyPosition(position common.Position) common.Piece {
	return MockPiece{
		kind:      piece.kind,
		color:     piece.color,
		position:  position,
		checkMove: piece.checkMove,
	}
}

func (piece MockPiece) CheckMove(move common.Move, storage common.PieceStorage) bool {
	if piece.checkMove == nil {
		panic("not implemented")
	}

	return piece.checkMove(move, storage)
}

type MockBasePieceStorage struct {
	size common.Size

	piece func(position common.Position) (piece common.Piece, ok bool)
}

func (storage MockBasePieceStorage) Size() common.Size {
	return storage.size
}

func (storage MockBasePieceStorage) Piece(
	position common.Position,
) (piece common.Piece, ok bool) {
	if storage.piece == nil {
		panic("not implemented")
	}

	return storage.piece(position)
}

func (storage MockBasePieceStorage) ApplyMove(move common.Move) common.PieceStorage {
	panic("not implemented")
}

type MockPieceGroupGetter struct {
	pieces []common.Piece
}

func (pieceGroupGetter MockPieceGroupGetter) Pieces() []common.Piece {
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

func TestPieceGroupGetterWrapperPieces(test *testing.T) {
	board := pieceGroupGetterWrapper{
		pieceStorageWithoutPieceGroupGetter: struct {
			MockBasePieceStorage
			MockMoveChecker
		}{
			MockBasePieceStorage: MockBasePieceStorage{
				size: common.Size{5, 5},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					if position != (common.Position{2, 3}) && position != (common.Position{4, 2}) {
						return nil, false
					}

					piece = MockPiece{position: position}
					return piece, true
				},
			},
		},
	}
	pieces := board.Pieces()

	expectedPieces := []common.Piece{
		MockPiece{position: common.Position{4, 2}},
		MockPiece{position: common.Position{2, 3}},
	}
	if !reflect.DeepEqual(pieces, expectedPieces) {
		test.Fail()
	}
}

func TestMoveCheckerWrapperCheckMove(test *testing.T) {
	type fields struct {
		size  common.Size
		piece func(position common.Position) (piece common.Piece, ok bool)
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
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					return nil, false
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{0, 0},
				},
			},
			want: common.ErrNoMove,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					return nil, false
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{-1, -1},
				},
			},
			want: common.ErrOutOfSize,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					return nil, false
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{1, 1},
				},
			},
			want: common.ErrNoPiece,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
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
			want: common.ErrFriendlyTarget,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					if position != (common.Position{0, 0}) {
						return nil, false
					}

					piece = MockPiece{
						position: position,
						checkMove: func(move common.Move, storage common.PieceStorage) bool {
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
			want: common.ErrIllegalMove,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					switch position {
					case common.Position{0, 0}:
						piece = MockPiece{
							color:    common.Black,
							position: common.Position{0, 0},
							checkMove: func(move common.Move, storage common.PieceStorage) bool {
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
			want: common.ErrKingCapture,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					if position != (common.Position{0, 0}) {
						return nil, false
					}

					piece = MockPiece{
						position: position,
						checkMove: func(move common.Move, storage common.PieceStorage) bool {
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
		board := moveCheckerWrapper{
			pieceStorageWithoutMoveChecker: struct {
				MockBasePieceStorage
				MockPieceGroupGetter
			}{
				MockBasePieceStorage: MockBasePieceStorage{
					size:  data.fields.size,
					piece: data.fields.piece,
				},
			},
		}
		got := board.CheckMove(data.args.move)

		if got != data.want {
			test.Fail()
		}
	}
}

func TestPieceStorageWrapperPieces(test *testing.T) {
	board := pieceStorageWrapper{
		BasePieceStorage: MockBasePieceStorage{
			size: common.Size{5, 5},
			piece: func(position common.Position) (piece common.Piece, ok bool) {
				if position != (common.Position{2, 3}) && position != (common.Position{4, 2}) {
					return nil, false
				}

				piece = MockPiece{position: position}
				return piece, true
			},
		},
	}
	pieces := board.Pieces()

	expectedPieces := []common.Piece{
		MockPiece{position: common.Position{4, 2}},
		MockPiece{position: common.Position{2, 3}},
	}
	if !reflect.DeepEqual(pieces, expectedPieces) {
		test.Fail()
	}
}

func TestPieceStorageWrapperCheckMove(test *testing.T) {
	type fields struct {
		size  common.Size
		piece func(position common.Position) (piece common.Piece, ok bool)
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
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					return nil, false
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{0, 0},
				},
			},
			want: common.ErrNoMove,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					return nil, false
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{-1, -1},
				},
			},
			want: common.ErrOutOfSize,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					return nil, false
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{1, 1},
				},
			},
			want: common.ErrNoPiece,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
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
			want: common.ErrFriendlyTarget,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					if position != (common.Position{0, 0}) {
						return nil, false
					}

					piece = MockPiece{
						position: position,
						checkMove: func(move common.Move, storage common.PieceStorage) bool {
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
			want: common.ErrIllegalMove,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					switch position {
					case common.Position{0, 0}:
						piece = MockPiece{
							color:    common.Black,
							position: common.Position{0, 0},
							checkMove: func(move common.Move, storage common.PieceStorage) bool {
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
			want: common.ErrKingCapture,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					if position != (common.Position{0, 0}) {
						return nil, false
					}

					piece = MockPiece{
						position: position,
						checkMove: func(move common.Move, storage common.PieceStorage) bool {
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
		board := pieceStorageWrapper{
			BasePieceStorage: MockBasePieceStorage{
				size:  data.fields.size,
				piece: data.fields.piece,
			},
		}
		got := board.CheckMove(data.args.move)

		if got != data.want {
			test.Fail()
		}
	}
}

func TestWrapBasePieceStorage(test *testing.T) {
	type args struct {
		baseStorage common.BasePieceStorage
	}
	type data struct {
		args args
		want common.PieceStorage
	}

	for _, data := range []data{
		{
			args: args{
				baseStorage: struct {
					MockBasePieceStorage
					MockMoveChecker
				}{
					MockBasePieceStorage: MockBasePieceStorage{
						size:  common.Size{5, 5},
						piece: nil,
					},
				},
			},
			want: pieceGroupGetterWrapper{
				pieceStorageWithoutPieceGroupGetter: struct {
					MockBasePieceStorage
					MockMoveChecker
				}{
					MockBasePieceStorage: MockBasePieceStorage{
						size:  common.Size{5, 5},
						piece: nil,
					},
				},
			},
		},
		{
			args: args{
				baseStorage: struct {
					MockBasePieceStorage
					MockPieceGroupGetter
				}{
					MockBasePieceStorage: MockBasePieceStorage{
						size:  common.Size{5, 5},
						piece: nil,
					},
				},
			},
			want: moveCheckerWrapper{
				pieceStorageWithoutMoveChecker: struct {
					MockBasePieceStorage
					MockPieceGroupGetter
				}{
					MockBasePieceStorage: MockBasePieceStorage{
						size:  common.Size{5, 5},
						piece: nil,
					},
				},
			},
		},
		{
			args: args{
				baseStorage: MockBasePieceStorage{
					size:  common.Size{5, 5},
					piece: nil,
				},
			},
			want: pieceStorageWrapper{
				BasePieceStorage: MockBasePieceStorage{
					size:  common.Size{5, 5},
					piece: nil,
				},
			},
		},
	} {
		got := WrapBasePieceStorage(data.args.baseStorage)

		if !reflect.DeepEqual(got, data.want) {
			test.Fail()
		}
	}
}

func TestDefaultBoardWrapperPieces(test *testing.T) {
	type fields struct {
		BasePieceStorage common.BasePieceStorage
	}
	type data struct {
		fields fields
		want   []common.Piece
	}

	for _, data := range []data{
		{
			fields: fields{
				BasePieceStorage: MockBasePieceStorage{
					size: common.Size{5, 5},
					piece: func(position common.Position) (piece common.Piece, ok bool) {
						if position != (common.Position{2, 3}) && position != (common.Position{4, 2}) {
							return nil, false
						}

						piece = MockPiece{position: position}
						return piece, true
					},
				},
			},
			want: []common.Piece{
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
						pieces: []common.Piece{
							MockPiece{position: common.Position{4, 2}},
							MockPiece{position: common.Position{2, 3}},
						},
					},
				},
			},
			want: []common.Piece{
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
		size  common.Size
		piece func(position common.Position) (piece common.Piece, ok bool)
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
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					return nil, false
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{0, 0},
				},
			},
			want: common.ErrNoMove,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					return nil, false
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{-1, -1},
				},
			},
			want: common.ErrOutOfSize,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					return nil, false
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{0, 0},
					Finish: common.Position{1, 1},
				},
			},
			want: common.ErrNoPiece,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
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
			want: common.ErrFriendlyTarget,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					if position != (common.Position{0, 0}) {
						return nil, false
					}

					piece = MockPiece{
						position: position,
						checkMove: func(move common.Move, storage common.PieceStorage) bool {
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
			want: common.ErrIllegalMove,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					switch position {
					case common.Position{0, 0}:
						piece = MockPiece{
							color:    common.Black,
							position: common.Position{0, 0},
							checkMove: func(move common.Move, storage common.PieceStorage) bool {
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
			want: common.ErrKingCapture,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				piece: func(position common.Position) (piece common.Piece, ok bool) {
					if position != (common.Position{0, 0}) {
						return nil, false
					}

					piece = MockPiece{
						position: position,
						checkMove: func(move common.Move, storage common.PieceStorage) bool {
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
		BasePieceStorage common.BasePieceStorage
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
