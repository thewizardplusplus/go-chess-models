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

func TestBaseApplyPosition(
	test *testing.T,
) {
	piece := Base{
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

	expectedPiece := Base{
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

	expectedNextPiece := Base{
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

func TestBaseString(test *testing.T) {
	type fields struct {
		kind  models.Kind
		color models.Color
	}
	type data struct {
		fields fields
		want   string
	}

	for _, data := range []data{
		data{
			fields: fields{
				kind:  models.King,
				color: models.White,
			},
			want: "K",
		},
		data{
			fields: fields{
				kind:  models.Queen,
				color: models.Black,
			},
			want: "q",
		},
	} {
		piece := Base{
			kind:  data.fields.kind,
			color: data.fields.color,
		}
		got := piece.String()

		if got != data.want {
			test.Fail()
		}
	}
}
