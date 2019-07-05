package chessmodels

import (
	"testing"
)

func TestKindToFEN(test *testing.T) {
	type data struct {
		kind    Kind
		wantFEN rune
		wantErr bool
	}

	for _, data := range []data{
		data{
			kind:    King,
			wantFEN: 'k',
			wantErr: false,
		},
		data{
			kind:    Queen,
			wantFEN: 'q',
			wantErr: false,
		},
		data{
			kind:    23,
			wantFEN: 0,
			wantErr: true,
		},
	} {
		gotFEN, gotErr := data.kind.ToFEN()

		if gotFEN != data.wantFEN {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}

func TestParsePiece(test *testing.T) {
	type args struct {
		kindInFEN rune
	}
	type data struct {
		args      args
		wantKind  Kind
		wantColor Color
		wantErr   bool
	}

	for _, data := range []data{
		data{
			args:      args{'K'},
			wantKind:  King,
			wantColor: White,
			wantErr:   false,
		},
		data{
			args:      args{'q'},
			wantKind:  Queen,
			wantColor: Black,
			wantErr:   false,
		},
		data{
			args:      args{'a'},
			wantKind:  0,
			wantColor: 0,
			wantErr:   true,
		},
	} {
		gotKind, gotColor, gotErr :=
			ParsePiece(data.args.kindInFEN)

		if gotKind != data.wantKind {
			test.Fail()
		}

		if gotColor != data.wantColor {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}
