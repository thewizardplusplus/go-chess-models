package uci

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func TestDecodePiece(test *testing.T) {
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
			wantPiece: pieces.NewKing(
				models.White,
				models.Position{},
			),
			wantErr: false,
		},
		data{
			args: args{'q'},
			wantPiece: pieces.NewQueen(
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
		gotPiece, gotErr := DecodePiece(
			data.args.fen,
			pieces.NewPiece,
		)

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
