package chessmodels

import (
	"testing"
)

func TestKindToFEN(test *testing.T) {
	type args struct {
		color Color
	}
	type data struct {
		kind    Kind
		args    args
		wantFEN rune
		wantErr bool
	}

	for _, data := range []data{
		data{
			kind:    King,
			args:    args{White},
			wantFEN: 'K',
			wantErr: false,
		},
		data{
			kind:    Queen,
			args:    args{Black},
			wantFEN: 'q',
			wantErr: false,
		},
		data{
			kind:    'a',
			args:    args{Black},
			wantFEN: 0,
			wantErr: true,
		},
	} {
		gotFEN, gotErr :=
			data.kind.ToFEN(data.args.color)

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
