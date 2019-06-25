package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestPieceFactoryType(test *testing.T) {
	gotType := reflect.TypeOf(NewPiece)
	wantType := reflect.
		TypeOf((*models.PieceFactory)(nil)).
		Elem()
	if !gotType.AssignableTo(wantType) {
		test.Fail()
	}
}

func TestNewPiece(test *testing.T) {
	type args struct {
		kind     models.Kind
		color    models.Color
		position models.Position
	}
	type data struct {
		args      args
		wantPiece models.Piece
		wantErr   bool
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
			wantErr: false,
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
			wantErr: false,
		},
		data{
			args: args{
				kind:     1e6,
				color:    models.Black,
				position: models.Position{4, 2},
			},
			wantPiece: nil,
			wantErr:   true,
		},
	} {
		gotPiece, gotErr := NewPiece(
			data.args.kind,
			data.args.color,
			data.args.position,
		)

		if gotPiece != data.wantPiece {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}
