package pieces

import (
	"reflect"
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
		{
			args: args{
				kind:  models.King,
				color: models.White,
				position: models.Position{
					File: 2,
					Rank: 3,
				},
			},
			want: NewKing(models.White, models.Position{
				File: 2,
				Rank: 3,
			}),
		},
		{
			args: args{
				kind:  models.Queen,
				color: models.Black,
				position: models.Position{
					File: 4,
					Rank: 2,
				},
			},
			want: NewQueen(models.Black, models.Position{
				File: 4,
				Rank: 2,
			}),
		},
	} {
		got := NewPiece(data.args.kind, data.args.color, data.args.position)

		if !reflect.DeepEqual(got, data.want) {
			test.Fail()
		}
	}
}
