package chessmodels

import (
	"reflect"
	"testing"
)

func TestBoardApplyMove(test *testing.T) {
	board := NewBoard(Size{5, 5}, []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})
	boardCopy := board.ApplyMove(Move{
		Start:  Position{4, 2},
		Finish: Position{6, 5},
	})

	expectedBoard := Board{
		size:  Size{5, 5},
		moves: Size{5, 5}.Moves(),
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

	expectedBoardCopy := Board{
		size:  Size{5, 5},
		moves: Size{5, 5}.Moves(),
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
		boardCopy,
		expectedBoardCopy,
	) {
		test.Fail()
	}
}
