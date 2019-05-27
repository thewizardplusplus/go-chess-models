package chessmodels

import (
	"reflect"
	"testing"
)

func TestNewBoard(test *testing.T) {
	board := NewBoard(Size{5, 5}, []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})

	expectedBoard := Board{
		size: Size{5, 5},
		pieces: PieceGroup{
			Position{2, 3}: MockPiece{
				position: Position{2, 3},
			},
			Position{4, 2}: MockPiece{
				position: Position{4, 2},
			},
		},
	}
	if !reflect.DeepEqual(
		board,
		expectedBoard,
	) {
		test.Fail()
	}
}

func TestBoardSize(test *testing.T) {
	board := NewBoard(Size{5, 5}, nil)

	if !reflect.DeepEqual(
		board.size,
		Size{5, 5},
	) {
		test.Fail()
	}
}

func TestBoardApplyMove(test *testing.T) {
	board := NewBoard(Size{5, 5}, []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})
	nextBoard := board.ApplyMove(Move{
		Start:  Position{4, 2},
		Finish: Position{6, 5},
	})

	expectedNextBoard := Board{
		size: Size{5, 5},
		pieces: PieceGroup{
			Position{2, 3}: MockPiece{
				position: Position{2, 3},
			},
			Position{6, 5}: MockPiece{
				position: Position{6, 5},
			},
		},
	}
	if !reflect.DeepEqual(
		nextBoard,
		expectedNextBoard,
	) {
		test.Fail()
	}
}
