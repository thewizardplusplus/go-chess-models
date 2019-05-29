package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestBaseKind(test *testing.T) {
	piece := Base{kind: models.Pawn}
	kind := piece.Kind()

	if kind != models.Pawn {
		test.Fail()
	}
}

func TestBaseColor(test *testing.T) {
	piece := Base{color: models.White}
	color := piece.Color()

	if color != models.White {
		test.Fail()
	}
}

func TestBasePosition(test *testing.T) {
	piece := Base{
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
