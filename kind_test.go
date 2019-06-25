package chessmodels

import (
	"testing"
)

func TestParseKind(test *testing.T) {
	type args struct {
		kindInFEN rune
	}
	type data struct {
		args     args
		wantKind Kind
		wantErr  bool
	}

	for _, data := range []data{
		data{
			args:     args{'K'},
			wantKind: King,
			wantErr:  false,
		},
		data{
			args:     args{'q'},
			wantKind: Queen,
			wantErr:  false,
		},
		data{
			args:     args{'a'},
			wantKind: 0,
			wantErr:  true,
		},
	} {
		gotKind, gotErr :=
			ParseKind(data.args.kindInFEN)

		if gotKind != data.wantKind {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}
