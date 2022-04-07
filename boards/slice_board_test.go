package boards

import (
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

func TestNewSliceBoard(test *testing.T) {
	board := NewSliceBoard(common.Size{5, 5}, []common.Piece{
		MockPiece{position: common.Position{2, 3}},
		MockPiece{position: common.Position{4, 2}},
	})

	expectedBoard := pieceStorageWrapper{
		BasePieceStorage: SliceBoard{
			BaseBoard: NewBaseBoard(common.Size{5, 5}),

			pieces: []common.Piece{
				14: MockPiece{position: common.Position{4, 2}},
				17: MockPiece{position: common.Position{2, 3}},
				24: nil,
			},
		},
	}
	if !reflect.DeepEqual(board, expectedBoard) {
		test.Fail()
	}
}

func TestSliceBoardPiece(test *testing.T) {
	type fields struct {
		size   common.Size
		pieces []common.Piece
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
				pieces: []common.Piece{
					14: MockPiece{position: common.Position{4, 2}},
					17: MockPiece{position: common.Position{2, 3}},
					24: nil,
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
				pieces: []common.Piece{
					14: MockPiece{position: common.Position{4, 2}},
					17: MockPiece{position: common.Position{2, 3}},
					24: nil,
				},
			},
			args:      args{common.Position{0, 0}},
			wantPiece: nil,
			wantOk:    false,
		},
	} {
		board := SliceBoard{
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

func TestSliceBoardApplyMove(test *testing.T) {
	board := NewSliceBoard(common.Size{5, 5}, []common.Piece{
		MockPiece{position: common.Position{2, 3}},
		MockPiece{position: common.Position{4, 2}},
	})
	nextBoard := board.ApplyMove(common.Move{
		Start:  common.Position{4, 2},
		Finish: common.Position{1, 2},
	})

	expectedBoard := pieceStorageWrapper{
		BasePieceStorage: SliceBoard{
			BaseBoard: NewBaseBoard(common.Size{5, 5}),

			pieces: []common.Piece{
				14: MockPiece{position: common.Position{4, 2}},
				17: MockPiece{position: common.Position{2, 3}},
				24: nil,
			},
		},
	}
	if !reflect.DeepEqual(board, expectedBoard) {
		test.Fail()
	}

	expectedNextBoard := pieceStorageWrapper{
		BasePieceStorage: SliceBoard{
			BaseBoard: NewBaseBoard(common.Size{5, 5}),

			pieces: []common.Piece{
				11: MockPiece{position: common.Position{1, 2}},
				17: MockPiece{position: common.Position{2, 3}},
				24: nil,
			},
		},
	}
	if !reflect.DeepEqual(nextBoard, expectedNextBoard) {
		test.Fail()
	}
}
