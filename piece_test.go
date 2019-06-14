package chessmodels

import (
	"reflect"
	"testing"
)

type MockPiece struct {
	kind     Kind
	color    Color
	position Position

	checkMove func(
		move Move,
		board Board,
	) bool
}

func (piece MockPiece) Kind() Kind {
	return piece.kind
}

func (piece MockPiece) Color() Color {
	return piece.color
}

func (piece MockPiece) Position() Position {
	return piece.position
}

func (piece MockPiece) ApplyPosition(
	position Position,
) Piece {
	return MockPiece{
		kind:      piece.kind,
		color:     piece.color,
		position:  position,
		checkMove: piece.checkMove,
	}
}

func (piece MockPiece) CheckMove(
	move Move,
	board Board,
) bool {
	if piece.checkMove == nil {
		panic("not implemented")
	}

	return piece.checkMove(move, board)
}

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
	if !reflect.DeepEqual(
		pieces,
		expectedPieces,
	) {
		test.Fail()
	}
}

func TestPieceGroupAdd(test *testing.T) {
	pieces := make(pieceGroup)
	pieces.Add(MockPiece{
		position: Position{2, 3},
	})
	pieces.Add(MockPiece{
		position: Position{4, 2},
	})

	expectedPieces := pieceGroup{
		Position{2, 3}: MockPiece{
			position: Position{2, 3},
		},
		Position{4, 2}: MockPiece{
			position: Position{4, 2},
		},
	}
	if !reflect.DeepEqual(
		pieces,
		expectedPieces,
	) {
		test.Fail()
	}
}

func TestPieceGroupMove(test *testing.T) {
	pieces := make(pieceGroup)
	pieces.Add(MockPiece{
		position: Position{2, 3},
	})
	pieces.Add(MockPiece{
		position: Position{4, 2},
	})
	pieces.Move(Move{
		Start:  Position{4, 2},
		Finish: Position{6, 5},
	})

	expectedPieces := pieceGroup{
		Position{2, 3}: MockPiece{
			position: Position{2, 3},
		},
		Position{6, 5}: MockPiece{
			position: Position{6, 5},
		},
	}
	if !reflect.DeepEqual(
		pieces,
		expectedPieces,
	) {
		test.Fail()
	}
}

func TestPieceGroupCopy(test *testing.T) {
	pieces := make(pieceGroup)
	pieces.Add(MockPiece{
		position: Position{2, 3},
	})

	piecesCopy := pieces.Copy()
	pieces.Add(MockPiece{
		position: Position{4, 2},
	})

	expectedPiecesCopy := pieceGroup{
		Position{2, 3}: MockPiece{
			position: Position{2, 3},
		},
	}
	if !reflect.DeepEqual(
		piecesCopy,
		expectedPiecesCopy,
	) {
		test.Fail()
	}
}
