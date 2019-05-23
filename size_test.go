package chessmodels

import (
	"reflect"
	"testing"
)

func TestSizePositions(test *testing.T) {
	positions := Size{2, 2}.Positions()

	expectedPositions := []Position{
		Position{0, 0},
		Position{1, 0},
		Position{0, 1},
		Position{1, 1},
	}
	if !reflect.DeepEqual(
		positions,
		expectedPositions,
	) {
		test.Fail()
	}
}

func TestSizeMovesForPosition(
	test *testing.T,
) {
	moves := Size{2, 2}.MovesForPosition(
		Position{1, 1},
	)

	expectedMoves := []Move{
		Move{Position{1, 1}, Position{0, 0}},
		Move{Position{1, 1}, Position{1, 0}},
		Move{Position{1, 1}, Position{0, 1}},
		Move{Position{1, 1}, Position{1, 1}},
	}
	if !reflect.DeepEqual(
		moves,
		expectedMoves,
	) {
		test.Fail()
	}
}
