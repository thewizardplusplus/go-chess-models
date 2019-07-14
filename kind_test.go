package chessmodels

import (
	"testing"
)

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

func TestKindToFEN(test *testing.T) {
	type args struct {
		color Color
	}
	type data struct {
		kind Kind
		args args
		want rune
	}

	for _, data := range []data{
		data{
			kind: King,
			args: args{White},
			want: 'K',
		},
		data{
			kind: Queen,
			args: args{Black},
			want: 'q',
		},
	} {
		got := data.kind.ToFEN(data.args.color)

		if got != data.want {
			test.Fail()
		}
	}
}
