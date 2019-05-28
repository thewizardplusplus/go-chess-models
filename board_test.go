package chessmodels

import (
	"reflect"
	"sort"
	"testing"
)

type ByPosition []Piece

func (group ByPosition) Len() int {
	return len(group)
}

func (group ByPosition) Swap(i, j int) {
	group[i], group[j] = group[j], group[i]
}

func (group ByPosition) Less(
	i, j int,
) bool {
	return positionLess(
		group[i].Position(),
		group[j].Position(),
	)
}

func positionLess(a, b Position) bool {
	return a.File < b.File &&
		a.Rank < b.Rank
}

func TestNewBoard(test *testing.T) {
	board := NewBoard(Size{5, 5}, []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})

	expectedBoard := Board{
		size: Size{5, 5},
		pieces: pieceGroup{
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

func TestBoardPieces(test *testing.T) {
	board := NewBoard(Size{5, 5}, []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})
	pieces := board.Pieces()
	sort.Sort(ByPosition(pieces))

	expectedPieces := []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	}
	if !reflect.DeepEqual(
		pieces,
		expectedPieces,
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
		pieces: pieceGroup{
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
