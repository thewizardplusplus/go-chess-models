// +build long

package chessmodels_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

type PositionMap map[int]int

var (
	kings   = "4k3/8/8/8/8/8/8/4K3"
	queens  = "3qk3/8/8/8/8/8/8/3QK3"
	rooks   = "r3k2r/8/8/8/8/8/8/R3K2R"
	bishops = "2b1kb2/8/8/8/8/8/8/2B1KB2"
	knights = "1n2k1n1/8/8/8/8/8/8/1N2K1N1"
	pawns   = "4k3/pppppppp/8/8" +
		"/8/8/PPPPPPPP/4K3"
	initial = "rnbqkbnr/pppppppp/8/8" +
		"/8/8/PPPPPPPP/RNBQKBNR"
	kiwipete = "r3k2r/p1ppqpb1/bn2pnp1/3PN3" +
		"/1p2P3/2N2Q1p/PPPBBPPP/R3K2R"
)

func TestPerft(test *testing.T) {
	type args struct {
		boardInFEN string
		color      models.Color
		deep       int
	}
	type data struct {
		name string
		args args
		want int
	}

	for index, data := range []data{
		data{
			name: "kings",
			args: args{
				boardInFEN: kings,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		data{
			name: "kings",
			args: args{
				boardInFEN: kings,
				color:      models.White,
				deep:       1,
			},
			want: 5,
		},
		data{
			name: "kings",
			args: args{
				boardInFEN: kings,
				color:      models.White,
				deep:       2,
			},
			want: 25,
		},
		data{
			name: "kings",
			args: args{
				boardInFEN: kings,
				color:      models.White,
				deep:       3,
			},
			want: 170,
		},
		data{
			name: "kings",
			args: args{
				boardInFEN: kings,
				color:      models.White,
				deep:       4,
			},
			want: 1156,
		},
		data{
			name: "kings",
			args: args{
				boardInFEN: kings,
				color:      models.White,
				deep:       5,
			},
			want: 7922,
		},
		data{
			name: "queens",
			args: args{
				boardInFEN: queens,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		data{
			name: "queens",
			args: args{
				boardInFEN: queens,
				color:      models.White,
				deep:       1,
			},
			want: 20,
		},
		data{
			name: "queens",
			args: args{
				boardInFEN: queens,
				color:      models.White,
				deep:       2,
			},
			want: 301,
		},
		data{
			name: "queens",
			args: args{
				boardInFEN: queens,
				color:      models.White,
				deep:       3,
			},
			want: 6063,
		},
		data{
			name: "rooks",
			args: args{
				boardInFEN: rooks,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		data{
			name: "rooks",
			args: args{
				boardInFEN: rooks,
				color:      models.White,
				deep:       1,
			},
			want: 24,
		},
		data{
			name: "rooks",
			args: args{
				boardInFEN: rooks,
				color:      models.White,
				deep:       2,
			},
			want: 482,
		},
		data{
			name: "rooks",
			args: args{
				boardInFEN: rooks,
				color:      models.White,
				deep:       3,
			},
			want: 11522,
		},
		data{
			name: "bishops",
			args: args{
				boardInFEN: bishops,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		data{
			name: "bishops",
			args: args{
				boardInFEN: bishops,
				color:      models.White,
				deep:       1,
			},
			want: 18,
		},
		data{
			name: "bishops",
			args: args{
				boardInFEN: bishops,
				color:      models.White,
				deep:       2,
			},
			want: 305,
		},
		data{
			name: "bishops",
			args: args{
				boardInFEN: bishops,
				color:      models.White,
				deep:       3,
			},
			want: 5575,
		},
		data{
			name: "knights",
			args: args{
				boardInFEN: knights,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		data{
			name: "knights",
			args: args{
				boardInFEN: knights,
				color:      models.White,
				deep:       1,
			},
			want: 11,
		},
		data{
			name: "knights",
			args: args{
				boardInFEN: knights,
				color:      models.White,
				deep:       2,
			},
			want: 121,
		},
		data{
			name: "knights",
			args: args{
				boardInFEN: knights,
				color:      models.White,
				deep:       3,
			},
			want: 1551,
		},
		data{
			name: "knights",
			args: args{
				boardInFEN: knights,
				color:      models.White,
				deep:       4,
			},
			want: 19764,
		},
		data{
			name: "pawns",
			args: args{
				boardInFEN: pawns,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		data{
			name: "pawns",
			args: args{
				boardInFEN: pawns,
				color:      models.White,
				deep:       1,
			},
			want: 10,
		},
		data{
			name: "pawns",
			args: args{
				boardInFEN: pawns,
				color:      models.White,
				deep:       2,
			},
			want: 100,
		},
		data{
			name: "pawns",
			args: args{
				boardInFEN: pawns,
				color:      models.White,
				deep:       3,
			},
			want: 1030,
		},
		data{
			name: "pawns",
			args: args{
				boardInFEN: pawns,
				color:      models.White,
				deep:       4,
			},
			want: 10609,
		},
		data{
			name: "initial",
			args: args{
				boardInFEN: initial,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		data{
			name: "initial",
			args: args{
				boardInFEN: initial,
				color:      models.White,
				deep:       1,
			},
			want: 12,
		},
		data{
			name: "initial",
			args: args{
				boardInFEN: initial,
				color:      models.White,
				deep:       2,
			},
			want: 144,
		},
		data{
			name: "initial",
			args: args{
				boardInFEN: initial,
				color:      models.White,
				deep:       3,
			},
			want: 2124,
		},
		data{
			name: "kiwipete",
			args: args{
				boardInFEN: kiwipete,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		data{
			name: "kiwipete",
			args: args{
				boardInFEN: kiwipete,
				color:      models.White,
				deep:       1,
			},
			want: 44,
		},
		data{
			name: "kiwipete",
			args: args{
				boardInFEN: kiwipete,
				color:      models.White,
				deep:       2,
			},
			want: 1740,
		},
	} {
		prefix := fmt.Sprintf(
			"%s/#%d",
			data.name,
			index,
		)
		storage, err := models.ParseBoard(
			data.args.boardInFEN,
			pieces.NewPiece,
		)
		if err != nil {
			test.Logf("%s: %v", prefix, err)
			test.Fail()

			continue
		}

		var moves []string
		got := perft(
			storage,
			data.args.color,
			data.args.deep,
			func(
				move models.Move,
				count int,
			) {
				moves = append(moves, fmt.Sprintf(
					"%v: %d",
					move,
					count,
				))
			},
		)

		if got != data.want {
			sort.Strings(moves)

			msg := "%s: %d/%d\n" +
				strings.Join(moves, "\n")
			test.Logf(msg, prefix, got, data.want)

			test.Fail()
		}
	}
}

func perft(
	storage models.PieceStorage,
	color models.Color,
	deep int,
	logger func(move models.Move, count int),
) int {
	// check for a check should be first,
	// including before a termination check,
	// because a terminated evaluation
	// doesn't make sense for a check position
	moves, err := models.MoveGenerator{}.
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
		moveCount := perft(
			nextStorage,
			nextColor,
			deep-1,
			nil, // log only a top level
		)
		if logger != nil {
			logger(move, moveCount)
		}

		count += moveCount
	}

	return count
}
