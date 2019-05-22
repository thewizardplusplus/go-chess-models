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

func (piece MockPiece) ApplyPosition(
	position Position,
) Piece {
	return MockPiece{position}
}

func TestPieceGroupAdd(test *testing.T) {
	pieces := make(PieceGroup)
	pieces.Add(MockPiece{Position{2, 3}})
	pieces.Add(MockPiece{Position{4, 2}})

	expectedPieces := PieceGroup{
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

func TestPieceGroupMove(test *testing.T) {
	pieces := make(PieceGroup)
	pieces.Add(MockPiece{Position{2, 3}})
	pieces.Add(MockPiece{Position{4, 2}})
	pieces.Move(Move{
		Start:  Position{4, 2},
		Finish: Position{6, 5},
	})

	expectedPieces := PieceGroup{
		Position{2, 3}: MockPiece{
			position: Position{2, 3},
		},
		Position{6, 5}: MockPiece{
			position: Position{6, 5},
		},
	}
	if !reflect.DeepEqual(pieces, expectedPieces) {
		test.Fail()
	}
}

func TestPieceGroupCopy(test *testing.T) {
	pieces := make(PieceGroup)
	pieces.Add(MockPiece{Position{2, 3}})

	piecesCopy := pieces.Copy()
	pieces.Add(MockPiece{Position{4, 2}})

	expectedPieces := PieceGroup{
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

	expectedPiecesCopy := PieceGroup{
		Position{2, 3}: MockPiece{
			position: Position{2, 3},
		},
	}
	if !reflect.DeepEqual(piecesCopy, expectedPiecesCopy) {
		test.Fail()
	}
}
