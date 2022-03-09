package chessmodels

import (
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

func TestPieces(test *testing.T) {
	storage := MockPieceStorage{
		MockBasePieceStorage: MockBasePieceStorage{
			size: common.Size{5, 5},
			piece: func(position common.Position) (piece Piece, ok bool) {
				if position != (common.Position{2, 3}) && position != (common.Position{4, 2}) {
					return nil, false
				}

				piece = MockPiece{position: position}
				return piece, true
			},
		},
	}
	pieces := Pieces(storage)

	expectedPieces := []Piece{
		MockPiece{position: common.Position{4, 2}},
		MockPiece{position: common.Position{2, 3}},
	}
	if !reflect.DeepEqual(pieces, expectedPieces) {
		test.Fail()
	}
}
