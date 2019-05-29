package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestRookInterface(test *testing.T) {
	rookType := reflect.TypeOf(Rook{})
	pieceType := reflect.
		TypeOf((*models.Piece)(nil)).
		Elem()
	if !rookType.Implements(pieceType) {
		test.Fail()
	}
}

func TestNewRook(test *testing.T) {
	piece := NewRook(
		models.White,
		models.Position{File: 2, Rank: 3},
	)

	expectedPiece := Rook{
		Base: Base{
			kind:  models.Rook,
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

func TestRookApplyPosition(test *testing.T) {
	piece := NewRook(
		models.White,
		models.Position{File: 2, Rank: 3},
	)
	nextPiece := piece.ApplyPosition(
		models.Position{
			File: 4,
			Rank: 2,
		},
	)

	expectedPiece := Rook{
		Base: Base{
			kind:  models.Rook,
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

	expectedNextPiece := Rook{
		Base: Base{
			kind:  models.Rook,
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
