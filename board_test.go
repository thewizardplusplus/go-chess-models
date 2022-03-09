package chessmodels

import (
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

func TestNewBoard(test *testing.T) {
	size := Size{5, 5}
	pieces := []Piece{
		MockPiece{position: common.Position{2, 3}},
		MockPiece{position: common.Position{4, 2}},
	}
	board := NewBoard(size, pieces)

	expectedBoard := NewSliceBoard(size, pieces)
	if !reflect.DeepEqual(board, expectedBoard) {
		test.Fail()
	}
}
