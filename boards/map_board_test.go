package boards

import (
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

func TestNewMapBoard(test *testing.T) {
	board := NewMapBoard(common.Size{5, 5}, []common.Piece{
		MockPiece{position: common.Position{2, 3}},
		MockPiece{position: common.Position{4, 2}},
	})

	expectedBoard := pieceStorageWrapper{
		BasePieceStorage: MapBoard{
			BaseBoard: NewBaseBoard(common.Size{5, 5}),

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
		wantPiece common.Piece
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
			BaseBoard: NewBaseBoard(data.fields.size),

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
	board := NewMapBoard(common.Size{5, 5}, []common.Piece{
		MockPiece{position: common.Position{2, 3}},
		MockPiece{position: common.Position{4, 2}},
	})
	nextBoard := board.ApplyMove(common.Move{
		Start:  common.Position{4, 2},
		Finish: common.Position{1, 2},
	})

	expectedBoard := pieceStorageWrapper{
		BasePieceStorage: MapBoard{
			BaseBoard: NewBaseBoard(common.Size{5, 5}),

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

	expectedNextBoard := pieceStorageWrapper{
		BasePieceStorage: MapBoard{
			BaseBoard: NewBaseBoard(common.Size{5, 5}),

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
