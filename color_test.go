package chessmodels

import (
	"testing"
)

func TestParseColor(test *testing.T) {
	type args struct {
		kindInFEN byte
	}
	type data struct {
		args args
		want Color
	}

	for _, data := range []data{
		data{
			args: args{'K'},
			want: White,
		},
		data{
			args: args{'k'},
			want: Black,
		},
	} {
		got := ParseColor(data.args.kindInFEN)

		if got != data.want {
			test.Fail()
		}
	}
}

func TestColorNegative(test *testing.T) {
	type data struct {
		color Color
		want  Color
	}

	for _, data := range []data{
		data{
			color: Black,
			want:  White,
		},
		data{
			color: White,
			want:  Black,
		},
	} {
		got := data.color.Negative()

		if got != data.want {
			test.Fail()
		}
	}
}
