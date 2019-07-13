package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestBaseKind(test *testing.T) {
	piece := base{kind: models.Pawn}
	kind := piece.Kind()

	if kind != models.Pawn {
		test.Fail()
	}
}

func TestBaseColor(test *testing.T) {
	piece := base{color: models.White}
	color := piece.Color()

	if color != models.White {
		test.Fail()
	}
}

func TestBasePosition(test *testing.T) {
	piece := base{
		position: models.Position{
			File: 2,
			Rank: 3,
		},
	}
	position := piece.Position()

	expectedPosition := models.Position{
		File: 2,
		Rank: 3,
	}
	if !reflect.DeepEqual(
		position,
		expectedPosition,
	) {
		test.Fail()
	}
}

func TestBaseApplyPosition(test *testing.T) {
	piece := base{
		kind:  models.Pawn,
		color: models.White,
		position: models.Position{
			File: 2,
			Rank: 3,
		},
	}
	nextPiece := piece.ApplyPosition(
		models.Position{
			File: 4,
			Rank: 2,
		},
	)

	expectedPiece := base{
		kind:  models.Pawn,
		color: models.White,
		position: models.Position{
			File: 2,
			Rank: 3,
		},
	}
	if !reflect.DeepEqual(
		piece,
		expectedPiece,
	) {
		test.Fail()
	}

	expectedNextPiece := base{
		kind:  models.Pawn,
		color: models.White,
		position: models.Position{
			File: 4,
			Rank: 2,
		},
	}
	if !reflect.DeepEqual(
		nextPiece,
		expectedNextPiece,
	) {
		test.Fail()
	}
}
