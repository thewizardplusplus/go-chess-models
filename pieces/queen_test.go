package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestQueenInterface(test *testing.T) {
	queenType := reflect.TypeOf(Queen{})
	pieceType := reflect.
		TypeOf((*models.Piece)(nil)).
		Elem()
	if !queenType.Implements(pieceType) {
		test.Fail()
	}
}

func TestNewQueen(test *testing.T) {
	piece := NewQueen(
		models.White,
		models.Position{File: 2, Rank: 3},
	)

	expectedPiece := Queen{
		Base: Base{
			kind:  models.Queen,
			color: models.White,
			position: models.Position{
				File: 2,
				Rank: 3,
			},
		},
	}
	if !reflect.DeepEqual(
		piece,
		expectedPiece,
	) {
		test.Fail()
	}
}

func TestQueenApplyPosition(test *testing.T) {
	piece := NewQueen(
		models.White,
		models.Position{File: 2, Rank: 3},
	)
	nextPiece := piece.ApplyPosition(
		models.Position{
			File: 4,
			Rank: 2,
		},
	)

	expectedPiece := Queen{
		Base: Base{
			kind:  models.Queen,
			color: models.White,
			position: models.Position{
				File: 2,
				Rank: 3,
			},
		},
	}
	if !reflect.DeepEqual(
		piece,
		expectedPiece,
	) {
		test.Fail()
	}

	expectedNextPiece := Queen{
		Base: Base{
			kind:  models.Queen,
			color: models.White,
			position: models.Position{
				File: 4,
				Rank: 2,
			},
		},
	}
	if !reflect.DeepEqual(
		nextPiece,
		expectedNextPiece,
	) {
		test.Fail()
	}
}
