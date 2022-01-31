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
