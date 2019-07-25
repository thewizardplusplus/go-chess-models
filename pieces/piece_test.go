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

func TestParsePiece(test *testing.T) {
	type args struct {
		fen rune
	}
	type data struct {
		args      args
		wantPiece models.Piece
		wantErr   bool
	}

	for _, data := range []data{
		data{
			args: args{'K'},
			wantPiece: NewKing(
				models.White,
				models.Position{},
			),
			wantErr: false,
		},
		data{
			args: args{'q'},
			wantPiece: NewQueen(
				models.Black,
				models.Position{},
			),
			wantErr: false,
		},
		data{
			args:      args{'a'},
			wantPiece: nil,
			wantErr:   true,
		},
	} {
		gotPiece, gotErr :=
			ParseDefaultPiece(data.args.fen)

		if !reflect.DeepEqual(
			gotPiece,
			data.wantPiece,
		) {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}
