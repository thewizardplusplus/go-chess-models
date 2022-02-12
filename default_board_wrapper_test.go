package chessmodels

import (
	"reflect"
	"testing"
)

type MockBasePieceStorage struct {
	size Size

	piece func(position Position) (piece Piece, ok bool)
}

func (storage MockBasePieceStorage) Size() Size {
	return storage.size
}

func (storage MockBasePieceStorage) Piece(
	position Position,
) (piece Piece, ok bool) {
	if storage.piece == nil {
		panic("not implemented")
	}

	return storage.piece(position)
}

func (storage MockBasePieceStorage) ApplyMove(move Move) PieceStorage {
	panic("not implemented")
}

func TestDefaultBoardWrapperPieces(test *testing.T) {
	baseStorage := MockBasePieceStorage{
		size: Size{5, 5},
		piece: func(position Position) (piece Piece, ok bool) {
			if position != (Position{2, 3}) && position != (Position{4, 2}) {
				return nil, false
			}

			piece = MockPiece{position: position}
			return piece, true
		},
	}
	board := defaultBoardWrapper{baseStorage}
	pieces := board.Pieces()

	expectedPieces := []Piece{
		MockPiece{position: Position{4, 2}},
		MockPiece{position: Position{2, 3}},
	}
	if !reflect.DeepEqual(pieces, expectedPieces) {
		test.Fail()
	}
}
