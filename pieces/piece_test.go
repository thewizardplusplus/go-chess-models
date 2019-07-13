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
		args      args
		wantPiece models.Piece
	}

	for _, data := range []data{
		data{
			args: args{
				kind:     models.King,
				color:    models.White,
				position: models.Position{2, 3},
			},
			wantPiece: NewKing(
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
			wantPiece: NewQueen(
				models.Black,
				models.Position{4, 2},
			),
		},
	} {
		gotPiece := NewPiece(
			data.args.kind,
			data.args.color,
			data.args.position,
		)

		if gotPiece != data.wantPiece {
			test.Fail()
		}
	}
}
