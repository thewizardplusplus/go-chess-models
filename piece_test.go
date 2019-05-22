package chessmodels

import (
	"reflect"
	"testing"
)

type MockPiece struct {
	position Position
}

func (piece MockPiece) Position() Position {
	return piece.position
}

func TestPieceGroupAdd(test *testing.T) {
	pieces := make(PieceGroup)
	pieces.Add(MockPiece{Position{2, 3}})
	pieces.Add(MockPiece{Position{4, 2}})

	expectedPieces := PieceGroup{
		Position{2, 3}: MockPiece{Position{2, 3}},
		Position{4, 2}: MockPiece{Position{4, 2}},
	}
	if !reflect.DeepEqual(pieces, expectedPieces) {
		test.Fail()
	}
}
