package chessmodels

import (
	"reflect"
	"testing"
)

func TestPieces(test *testing.T) {
	storage := MockPieceStorage{
		size: Size{5, 5},
		piece: func(position Position) (piece Piece, ok bool) {
			if position != (Position{2, 3}) && position != (Position{4, 2}) {
				return nil, false
			}

			piece = MockPiece{position: position}
			return piece, true
		},
	}
	pieces := Pieces(storage)

	expectedPieces := []Piece{
		MockPiece{position: Position{4, 2}},
		MockPiece{position: Position{2, 3}},
	}
	if !reflect.DeepEqual(pieces, expectedPieces) {
		test.Fail()
	}
}
