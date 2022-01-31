package chessmodels

import (
	"reflect"
	"testing"
)

func TestNewPieceGroup(test *testing.T) {
	pieces := newPieceGroup([]Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})

	expectedPieces := pieceGroup{
		Position{2, 3}: MockPiece{
			position: Position{2, 3},
		},
		Position{4, 2}: MockPiece{
			position: Position{4, 2},
		},
	}
	if !reflect.DeepEqual(pieces, expectedPieces) {
		test.Fail()
	}
}
