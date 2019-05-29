package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestKingInterface(test *testing.T) {
	kingType := reflect.TypeOf(King{})
	pieceType := reflect.
		TypeOf((*models.Piece)(nil)).
		Elem()
	if !kingType.Implements(pieceType) {
		test.Fail()
	}
}

func TestNewKing(test *testing.T) {
	piece := NewKing(
		models.White,
		models.Position{File: 2, Rank: 3},
	)

	expectedPiece := King{
		Base: Base{
			kind:  models.King,
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

func TestKingApplyPosition(test *testing.T) {
	piece := NewKing(
		models.White,
		models.Position{File: 2, Rank: 3},
	)
	nextPiece := piece.ApplyPosition(
		models.Position{
			File: 4,
			Rank: 2,
		},
	)

	expectedPiece := King{
		Base: Base{
			kind:  models.King,
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

	expectedNextPiece := King{
		Base: Base{
			kind:  models.King,
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
