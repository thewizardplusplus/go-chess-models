package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestKnightInterface(test *testing.T) {
	knightType := reflect.TypeOf(Knight{})
	pieceType := reflect.
		TypeOf((*models.Piece)(nil)).
		Elem()
	if !knightType.Implements(pieceType) {
		test.Fail()
	}
}

func TestNewKnight(test *testing.T) {
	piece := NewKnight(
		models.White,
		models.Position{File: 2, Rank: 3},
	)

	expectedPiece := Knight{
		Base: Base{
			kind:  models.Knight,
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

func TestKnightApplyPosition(test *testing.T) {
	piece := NewKnight(
		models.White,
		models.Position{File: 2, Rank: 3},
	)
	nextPiece := piece.ApplyPosition(
		models.Position{
			File: 4,
			Rank: 2,
		},
	)

	expectedPiece := Knight{
		Base: Base{
			kind:  models.Knight,
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

	expectedNextPiece := Knight{
		Base: Base{
			kind:  models.Knight,
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
