package chessmodels

import (
	"reflect"
	"testing"
)

func TestNewSliceBoard(test *testing.T) {
	board := NewSliceBoard(Size{5, 5}, []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})

	expectedBoard := SliceBoard{
		size: Size{5, 5},
		pieces: []Piece{
			14: MockPiece{position: Position{4, 2}},
			17: MockPiece{position: Position{2, 3}},
			24: nil,
		},
	}
	if !reflect.DeepEqual(board, expectedBoard) {
		test.Fail()
	}
}

func TestSliceBoardSize(test *testing.T) {
	board := NewSliceBoard(Size{5, 5}, nil)
	size := board.Size()

	if !reflect.DeepEqual(size, Size{5, 5}) {
		test.Fail()
	}
}

func TestSliceBoardPiece(test *testing.T) {
	type fields struct {
		size   Size
		pieces []Piece
	}
	type args struct {
		position Position
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
				size: Size{5, 5},
				pieces: []Piece{
					14: MockPiece{position: Position{4, 2}},
					17: MockPiece{position: Position{2, 3}},
					24: nil,
				},
			},
			args: args{Position{2, 3}},
			wantPiece: MockPiece{
				position: Position{2, 3},
			},
			wantOk: true,
		},
		{
			fields: fields{
				size: Size{5, 5},
				pieces: []Piece{
					14: MockPiece{position: Position{4, 2}},
					17: MockPiece{position: Position{2, 3}},
					24: nil,
				},
			},
			args:      args{Position{0, 0}},
			wantPiece: nil,
			wantOk:    false,
		},
	} {
		board := SliceBoard{
			size:   data.fields.size,
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
