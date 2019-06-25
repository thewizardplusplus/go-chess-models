package chessmodels

import (
	"reflect"
	"testing"
)

func TestParseRank(test *testing.T) {
	type args struct {
		rankIndex    int
		rankInFEN    string
		pieceFactory PieceFactory
	}
	type data struct {
		args        args
		wantPieces  []Piece
		wantMaxFile int
		wantErr     bool
	}

	for _, data := range []data{} {
		gotPieces, gotMaxFile, gotErr :=
			parseRank(
				data.args.rankIndex,
				data.args.rankInFEN,
				data.args.pieceFactory,
			)

		if !reflect.DeepEqual(
			gotPieces,
			data.wantPieces,
		) {
			test.Fail()
		}
		if gotMaxFile != data.wantMaxFile {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}
