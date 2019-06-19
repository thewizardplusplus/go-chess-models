package chessmodels

import (
	"testing"
)

func TestPerft(test *testing.T) {
	type args struct {
		storage PieceStorage
		color   Color
		deep    int
	}
	type data struct {
		args args
		want int
	}

	for _, data := range []data{
		data{
			args: args{
				storage: NewBoard(
					Size{8, 8},
					[]Piece{},
				),
				color: Black,
				deep:  0,
			},
			want: 0,
		},
	} {
		got := perft(
			data.args.storage,
			data.args.color,
			data.args.deep,
		)

		if got != data.want {
			test.Fail()
		}
	}
}

func perft(
	storage PieceStorage,
	color Color,
	deep int,
) int {
	// check for a check should be first,
	// including before a termination check,
	// because a terminated evaluation
	// doesn't make sense for a check position
	moves, err := MoveGenerator{}.
		MovesForColor(storage, color)
	if err != nil {
		return 0
	}

	if deep == 0 {
		return 1
	}

	var count int
	for _, move := range moves {
		nextStorage := storage.ApplyMove(move)
		nextColor := color.Negative()
		count += perft(
			nextStorage,
			nextColor,
			deep-1,
		)
	}

	return count
}
