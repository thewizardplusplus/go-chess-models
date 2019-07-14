package pieces

import (
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestNewPiece(test *testing.T) {
	type args struct {
		kind     models.Kind
		color    models.Color
		position models.Position
	}
	type data struct {
		args args
		want models.Piece
	}

	for _, data := range []data{
		data{
			args: args{
				kind:     models.King,
				color:    models.White,
				position: models.Position{2, 3},
			},
			want: NewKing(
				models.White,
				models.Position{2, 3},
			),
		},
		data{
			args: args{
				kind:     models.Queen,
				color:    models.Black,
				position: models.Position{4, 2},
			},
			want: NewQueen(
				models.Black,
				models.Position{4, 2},
			),
		},
	} {
		got := NewPiece(
			data.args.kind,
			data.args.color,
			data.args.position,
		)

		if got != data.want {
			test.Fail()
		}
	}
}
