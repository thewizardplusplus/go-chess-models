package chessmodels

import (
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

type MockPiece struct {
	kind     common.Kind
	color    common.Color
	position common.Position

	checkMove func(move common.Move, storage PieceStorage) bool
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

func (piece MockPiece) ApplyPosition(position common.Position) Piece {
	return MockPiece{
		kind:      piece.kind,
		color:     piece.color,
		position:  position,
		checkMove: piece.checkMove,
	}
}

func (piece MockPiece) CheckMove(move common.Move, storage PieceStorage) bool {
	if piece.checkMove == nil {
		panic("not implemented")
	}

	return piece.checkMove(move, storage)
}

func TestNewMapBoard(test *testing.T) {
	board := NewMapBoard(common.Size{5, 5}, []Piece{
		MockPiece{position: common.Position{2, 3}},
		MockPiece{position: common.Position{4, 2}},
	})

	expectedBoard := DefaultBoardWrapper{
		BasePieceStorage: MapBoard{
			BaseBoard: BaseBoard{
				size: common.Size{5, 5},
			},

			pieces: pieceGroup{
				common.Position{2, 3}: MockPiece{
					position: common.Position{2, 3},
				},
				common.Position{4, 2}: MockPiece{
					position: common.Position{4, 2},
				},
			},
		},
	}
	if !reflect.DeepEqual(board, expectedBoard) {
		test.Fail()
	}
}

func TestMapBoardPiece(test *testing.T) {
	type fields struct {
		size   common.Size
		pieces pieceGroup
	}
	type args struct {
		position common.Position
	}
	type data struct {
		fields    fields
		args      args
		wantPiece Piece
		wantOk    bool
	}

	for _, data := range []data{
		{
			fields: fields{
				size: common.Size{5, 5},
				pieces: pieceGroup{
					common.Position{2, 3}: MockPiece{
						position: common.Position{2, 3},
					},
					common.Position{4, 2}: MockPiece{
						position: common.Position{4, 2},
					},
				},
			},
			args: args{common.Position{2, 3}},
			wantPiece: MockPiece{
				position: common.Position{2, 3},
			},
			wantOk: true,
		},
		{
			fields: fields{
				size: common.Size{5, 5},
				pieces: pieceGroup{
					common.Position{2, 3}: MockPiece{
						position: common.Position{2, 3},
					},
					common.Position{4, 2}: MockPiece{
						position: common.Position{4, 2},
					},
				},
			},
			args:      args{common.Position{0, 0}},
			wantPiece: nil,
			wantOk:    false,
		},
	} {
		board := MapBoard{
			BaseBoard: BaseBoard{
				size: data.fields.size,
			},

			pieces: data.fields.pieces,
		}
		gotPiece, gotOk := board.Piece(data.args.position)

		if !reflect.DeepEqual(gotPiece, data.wantPiece) {
			test.Fail()
		}
		if gotOk != data.wantOk {
			test.Fail()
		}
	}
}

func TestMapBoardApplyMove(test *testing.T) {
	board := NewMapBoard(common.Size{5, 5}, []Piece{
		MockPiece{position: common.Position{2, 3}},
		MockPiece{position: common.Position{4, 2}},
	})
	nextBoard := board.ApplyMove(common.Move{
		Start:  common.Position{4, 2},
		Finish: common.Position{1, 2},
	})

	expectedBoard := DefaultBoardWrapper{
		BasePieceStorage: MapBoard{
			BaseBoard: BaseBoard{
				size: common.Size{5, 5},
			},

			pieces: pieceGroup{
				common.Position{2, 3}: MockPiece{
					position: common.Position{2, 3},
				},
				common.Position{4, 2}: MockPiece{
					position: common.Position{4, 2},
				},
			},
		},
	}
	if !reflect.DeepEqual(board, expectedBoard) {
		test.Fail()
	}

	expectedNextBoard := DefaultBoardWrapper{
		BasePieceStorage: MapBoard{
			BaseBoard: BaseBoard{
				size: common.Size{5, 5},
			},

			pieces: pieceGroup{
				common.Position{1, 2}: MockPiece{
					position: common.Position{1, 2},
				},
				common.Position{2, 3}: MockPiece{
					position: common.Position{2, 3},
				},
			},
		},
	}
	if !reflect.DeepEqual(nextBoard, expectedNextBoard) {
		test.Fail()
	}
}
