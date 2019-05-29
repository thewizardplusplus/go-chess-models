package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestPawnInterface(test *testing.T) {
	pawnType := reflect.TypeOf(Pawn{})
	pieceType := reflect.
		TypeOf((*models.Piece)(nil)).
		Elem()
	if !pawnType.Implements(pieceType) {
		test.Fail()
	}
}

func TestNewPawn(test *testing.T) {
	piece := NewPawn(
		models.White,
		models.Position{File: 2, Rank: 3},
	)

	expectedPiece := Pawn{
		Base: Base{
			kind:  models.Pawn,
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

func TestPawnApplyPosition(test *testing.T) {
	piece := NewPawn(
		models.White,
		models.Position{File: 2, Rank: 3},
	)
	nextPiece := piece.ApplyPosition(
		models.Position{
			File: 4,
			Rank: 2,
		},
	)

	expectedPiece := Pawn{
		Base: Base{
			kind:  models.Pawn,
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

	expectedNextPiece := Pawn{
		Base: Base{
			kind:  models.Pawn,
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
