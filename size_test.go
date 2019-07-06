package chessmodels

import (
	"reflect"
	"testing"
)

func TestSizePositions(test *testing.T) {
	positions := Size{3, 3}.Positions()

	expectedPositions := []Position{
		Position{0, 0},
		Position{1, 0},
		Position{2, 0},
		Position{0, 1},
		Position{1, 1},
		Position{2, 1},
		Position{0, 2},
		Position{1, 2},
		Position{2, 2},
	}
	if !reflect.DeepEqual(
		positions,
		expectedPositions,
	) {
		test.Fail()
	}
}
