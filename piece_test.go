package chessmodels

import (
	"reflect"
	"testing"
)

type MockPiece struct {
	kind     Kind
	color    Color
	position Position
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
		kind:     piece.kind,
		color:    piece.color,
		position: position,
	}
}

func (piece MockPiece) CheckMove(
	move Move,
	board Board,
) bool {
	return true
}

func TestNewPieceGroup(test *testing.T) {
	pieces := NewPieceGroup([]Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})

	expectedPieces := PieceGroup{
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
	pieces := make(PieceGroup)
	pieces.Add(MockPiece{
		position: Position{2, 3},
	})
	pieces.Add(MockPiece{
		position: Position{4, 2},
	})

	expectedPieces := PieceGroup{
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
	pieces := make(PieceGroup)
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

	expectedPieces := PieceGroup{
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
	pieces := make(PieceGroup)
	pieces.Add(MockPiece{
		position: Position{2, 3},
	})

	piecesCopy := pieces.Copy()
	pieces.Add(MockPiece{
		position: Position{4, 2},
	})

	expectedPieces := PieceGroup{
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

	expectedPiecesCopy := PieceGroup{
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

func TestPieceGroupByColor(
	test *testing.T,
) {
	pieces := make(PieceGroup)
	pieces.Add(MockPiece{
		color:    Black,
		position: Position{0, 5},
	})
	pieces.Add(MockPiece{
		color:    White,
		position: Position{1, 2},
	})
	pieces.Add(MockPiece{
		color:    Black,
		position: Position{2, 3},
	})
	pieces.Add(MockPiece{
		color:    White,
		position: Position{4, 2},
	})

	whitePieces := pieces.ByColor(White)
	expectedWhitePieces := []Piece{
		MockPiece{
			color:    White,
			position: Position{1, 2},
		},
		MockPiece{
			color:    White,
			position: Position{4, 2},
		},
	}
	if !reflect.DeepEqual(
		whitePieces,
		expectedWhitePieces,
	) {
		test.Fail()
	}
}
