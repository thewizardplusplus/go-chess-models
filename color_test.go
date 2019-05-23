package chessmodels

import (
	"testing"
)

func TestColorNegative(test *testing.T) {
	type data struct {
		color Color
		want  Color
	}

	for _, data := range []data{
		data{Black, White},
		data{White, Black},
	} {
		got := data.color.Negative()
		if got != data.want {
			test.Fail()
		}
	}
}
