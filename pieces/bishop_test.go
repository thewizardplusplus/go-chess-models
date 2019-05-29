package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestBishopInterface(test *testing.T) {
	bishopType := reflect.TypeOf(Bishop{})
	pieceType := reflect.
		TypeOf((*models.Piece)(nil)).
		Elem()
	if !bishopType.Implements(pieceType) {
		test.Fail()
	}
}

func TestNewBishop(test *testing.T) {
	piece := NewBishop(
		models.White,
		models.Position{File: 2, Rank: 3},
	)

	expectedPiece := Bishop{
		Base: Base{
			kind:  models.Bishop,
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

func TestBishopApplyPosition(test *testing.T) {
	piece := NewBishop(
		models.White,
		models.Position{File: 2, Rank: 3},
	)
	nextPiece := piece.ApplyPosition(
		models.Position{
			File: 4,
			Rank: 2,
		},
	)

	expectedPiece := Bishop{
		Base: Base{
			kind:  models.Bishop,
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

	expectedNextPiece := Bishop{
		Base: Base{
			kind:  models.Bishop,
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
