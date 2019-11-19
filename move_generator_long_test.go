// +build long

package chessmodels_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

var (
	kings    = "4k3/8/8/8/8/8/8/4K3"
	queens   = "3qk3/8/8/8/8/8/8/3QK3"
	rooks    = "r3k2r/8/8/8/8/8/8/R3K2R"
	bishops  = "2b1kb2/8/8/8/8/8/8/2B1KB2"
	knights  = "1n2k1n1/8/8/8/8/8/8/1N2K1N1"
	pawns    = "4k3/pppppppp/8/8/8/8/PPPPPPPP/4K3"
	initial  = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	kiwipete = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R"
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
		{
			name: "kings",
			args: args{
				boardInFEN: kings,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		{
			name: "kings",
			args: args{
				boardInFEN: kings,
				color:      models.White,
				deep:       1,
			},
			want: 5,
		},
		{
			name: "kings",
			args: args{
				boardInFEN: kings,
				color:      models.White,
				deep:       2,
			},
			want: 25,
		},
		{
			name: "kings",
			args: args{
				boardInFEN: kings,
				color:      models.White,
				deep:       3,
			},
			want: 170,
		},
		{
			name: "kings",
			args: args{
				boardInFEN: kings,
				color:      models.White,
				deep:       4,
			},
			want: 1156,
		},
		{
			name: "kings",
			args: args{
				boardInFEN: kings,
				color:      models.White,
				deep:       5,
			},
			want: 7922,
		},
		{
			name: "queens",
			args: args{
				boardInFEN: queens,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		{
			name: "queens",
			args: args{
				boardInFEN: queens,
				color:      models.White,
				deep:       1,
			},
			want: 20,
		},
		{
			name: "queens",
			args: args{
				boardInFEN: queens,
				color:      models.White,
				deep:       2,
			},
			want: 301,
		},
		{
			name: "queens",
			args: args{
				boardInFEN: queens,
				color:      models.White,
				deep:       3,
			},
			want: 6063,
		},
		{
			name: "rooks",
			args: args{
				boardInFEN: rooks,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		{
			name: "rooks",
			args: args{
				boardInFEN: rooks,
				color:      models.White,
				deep:       1,
			},
			want: 24,
		},
		{
			name: "rooks",
			args: args{
				boardInFEN: rooks,
				color:      models.White,
				deep:       2,
			},
			want: 482,
		},
		{
			name: "rooks",
			args: args{
				boardInFEN: rooks,
				color:      models.White,
				deep:       3,
			},
			want: 11522,
		},
		{
			name: "bishops",
			args: args{
				boardInFEN: bishops,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		{
			name: "bishops",
			args: args{
				boardInFEN: bishops,
				color:      models.White,
				deep:       1,
			},
			want: 18,
		},
		{
			name: "bishops",
			args: args{
				boardInFEN: bishops,
				color:      models.White,
				deep:       2,
			},
			want: 305,
		},
		{
			name: "bishops",
			args: args{
				boardInFEN: bishops,
				color:      models.White,
				deep:       3,
			},
			want: 5575,
		},
		{
			name: "knights",
			args: args{
				boardInFEN: knights,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		{
			name: "knights",
			args: args{
				boardInFEN: knights,
				color:      models.White,
				deep:       1,
			},
			want: 11,
		},
		{
			name: "knights",
			args: args{
				boardInFEN: knights,
				color:      models.White,
				deep:       2,
			},
			want: 121,
		},
		{
			name: "knights",
			args: args{
				boardInFEN: knights,
				color:      models.White,
				deep:       3,
			},
			want: 1551,
		},
		{
			name: "knights",
			args: args{
				boardInFEN: knights,
				color:      models.White,
				deep:       4,
			},
			want: 19764,
		},
		{
			name: "pawns",
			args: args{
				boardInFEN: pawns,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		{
			name: "pawns",
			args: args{
				boardInFEN: pawns,
				color:      models.White,
				deep:       1,
			},
			want: 10,
		},
		{
			name: "pawns",
			args: args{
				boardInFEN: pawns,
				color:      models.White,
				deep:       2,
			},
			want: 100,
		},
		{
			name: "pawns",
			args: args{
				boardInFEN: pawns,
				color:      models.White,
				deep:       3,
			},
			want: 1030,
		},
		{
			name: "pawns",
			args: args{
				boardInFEN: pawns,
				color:      models.White,
				deep:       4,
			},
			want: 10609,
		},
		{
			name: "initial",
			args: args{
				boardInFEN: initial,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		{
			name: "initial",
			args: args{
				boardInFEN: initial,
				color:      models.White,
				deep:       1,
			},
			want: 12,
		},
		{
			name: "initial",
			args: args{
				boardInFEN: initial,
				color:      models.White,
				deep:       2,
			},
			want: 144,
		},
		{
			name: "initial",
			args: args{
				boardInFEN: initial,
				color:      models.White,
				deep:       3,
			},
			want: 2124,
		},
		{
			name: "kiwipete",
			args: args{
				boardInFEN: kiwipete,
				color:      models.White,
				deep:       0,
			},
			want: 1,
		},
		{
			name: "kiwipete",
			args: args{
				boardInFEN: kiwipete,
				color:      models.White,
				deep:       1,
			},
			want: 44,
		},
		{
			name: "kiwipete",
			args: args{
				boardInFEN: kiwipete,
				color:      models.White,
				deep:       2,
			},
			want: 1740,
		},
	} {
		prefix := fmt.Sprintf("%s/#%d", data.name, index)
		storage, err := uci.DecodePieceStorage(
			data.args.boardInFEN,
			pieces.NewPiece,
			models.NewBoard,
		)
		if err != nil {
			test.Errorf("%s: %v", prefix, err)
			continue
		}

		var moves []string
		got := perft(storage, data.args.color, data.args.deep, func(
			move models.Move,
			count int,
		) {
			moves = append(moves, fmt.Sprintf("%v: %d", move, count))
		})

		if got != data.want {
			sort.Strings(moves)

			message := "%s: %d/%d\n" + strings.Join(moves, "\n")
			test.Errorf(message, prefix, got, data.want)
		}
	}
}

func perft(
	storage models.PieceStorage,
	color models.Color,
	deep int,
	logger func(move models.Move, count int),
) int {
	// check for a check should be first, including before a termination check,
	// because a terminated evaluation doesn't make sense for a check position
	var generator models.MoveGenerator
	moves, err := generator.MovesForColor(storage, color)
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
