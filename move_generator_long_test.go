// +build long

package chessmodels_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

var (
	kings   = "4k3/8/8/8/8/8/8/4K3"
	queens  = "3qk3/8/8/8/8/8/8/3QK3"
	rooks   = "r3k2r/8/8/8/8/8/8/R3K2R"
	bishops = "2b1kb2/8/8/8/8/8/8/2B1KB2"
	knights = "1n2k1n1/8/8/8/8/8/8/1N2K1N1"
	pawns   = "4k3/pppppppp/8/8/8/8/PPPPPPPP/4K3"
)

func TestPerft(test *testing.T) {
	type args struct {
		boardInFEN string
		color      common.Color
		deep       int
	}
	type storage struct {
		name    string
		factory uci.PieceStorageFactory
	}
	type data struct {
		name string
		args args
		want int
	}

	for _, storage := range []storage{
		{
			name:    "MapBoard",
			factory: models.NewMapBoard,
		},
		{
			name:    "SliceBoard",
			factory: models.NewSliceBoard,
		},
	} {
		for _, data := range []data{
			{
				name: "kings",
				args: args{
					boardInFEN: kings,
					color:      common.White,
					deep:       0,
				},
				want: 1,
			},
			{
				name: "kings",
				args: args{
					boardInFEN: kings,
					color:      common.White,
					deep:       1,
				},
				want: 5,
			},
			{
				name: "kings",
				args: args{
					boardInFEN: kings,
					color:      common.White,
					deep:       2,
				},
				want: 25,
			},
			{
				name: "kings",
				args: args{
					boardInFEN: kings,
					color:      common.White,
					deep:       3,
				},
				want: 170,
			},
			{
				name: "kings",
				args: args{
					boardInFEN: kings,
					color:      common.White,
					deep:       4,
				},
				want: 1156,
			},
			{
				name: "kings",
				args: args{
					boardInFEN: kings,
					color:      common.White,
					deep:       5,
				},
				want: 7922,
			},
			{
				name: "queens",
				args: args{
					boardInFEN: queens,
					color:      common.White,
					deep:       0,
				},
				want: 1,
			},
			{
				name: "queens",
				args: args{
					boardInFEN: queens,
					color:      common.White,
					deep:       1,
				},
				want: 20,
			},
			{
				name: "queens",
				args: args{
					boardInFEN: queens,
					color:      common.White,
					deep:       2,
				},
				want: 301,
			},
			{
				name: "queens",
				args: args{
					boardInFEN: queens,
					color:      common.White,
					deep:       3,
				},
				want: 6063,
			},
			{
				name: "rooks",
				args: args{
					boardInFEN: rooks,
					color:      common.White,
					deep:       0,
				},
				want: 1,
			},
			{
				name: "rooks",
				args: args{
					boardInFEN: rooks,
					color:      common.White,
					deep:       1,
				},
				want: 24,
			},
			{
				name: "rooks",
				args: args{
					boardInFEN: rooks,
					color:      common.White,
					deep:       2,
				},
				want: 482,
			},
			{
				name: "rooks",
				args: args{
					boardInFEN: rooks,
					color:      common.White,
					deep:       3,
				},
				want: 11522,
			},
			{
				name: "bishops",
				args: args{
					boardInFEN: bishops,
					color:      common.White,
					deep:       0,
				},
				want: 1,
			},
			{
				name: "bishops",
				args: args{
					boardInFEN: bishops,
					color:      common.White,
					deep:       1,
				},
				want: 18,
			},
			{
				name: "bishops",
				args: args{
					boardInFEN: bishops,
					color:      common.White,
					deep:       2,
				},
				want: 305,
			},
			{
				name: "bishops",
				args: args{
					boardInFEN: bishops,
					color:      common.White,
					deep:       3,
				},
				want: 5575,
			},
			{
				name: "knights",
				args: args{
					boardInFEN: knights,
					color:      common.White,
					deep:       0,
				},
				want: 1,
			},
			{
				name: "knights",
				args: args{
					boardInFEN: knights,
					color:      common.White,
					deep:       1,
				},
				want: 11,
			},
			{
				name: "knights",
				args: args{
					boardInFEN: knights,
					color:      common.White,
					deep:       2,
				},
				want: 121,
			},
			{
				name: "knights",
				args: args{
					boardInFEN: knights,
					color:      common.White,
					deep:       3,
				},
				want: 1551,
			},
			{
				name: "knights",
				args: args{
					boardInFEN: knights,
					color:      common.White,
					deep:       4,
				},
				want: 19764,
			},
			{
				name: "pawns",
				args: args{
					boardInFEN: pawns,
					color:      common.White,
					deep:       0,
				},
				want: 1,
			},
			{
				name: "pawns",
				args: args{
					boardInFEN: pawns,
					color:      common.White,
					deep:       1,
				},
				want: 10,
			},
			{
				name: "pawns",
				args: args{
					boardInFEN: pawns,
					color:      common.White,
					deep:       2,
				},
				want: 100,
			},
			{
				name: "pawns",
				args: args{
					boardInFEN: pawns,
					color:      common.White,
					deep:       3,
				},
				want: 1030,
			},
			{
				name: "pawns",
				args: args{
					boardInFEN: pawns,
					color:      common.White,
					deep:       4,
				},
				want: 10609,
			},
			{
				name: "initial",
				args: args{
					boardInFEN: initial,
					color:      common.White,
					deep:       0,
				},
				want: 1,
			},
			{
				name: "initial",
				args: args{
					boardInFEN: initial,
					color:      common.White,
					deep:       1,
				},
				want: 12,
			},
			{
				name: "initial",
				args: args{
					boardInFEN: initial,
					color:      common.White,
					deep:       2,
				},
				want: 144,
			},
			{
				name: "initial",
				args: args{
					boardInFEN: initial,
					color:      common.White,
					deep:       3,
				},
				want: 2124,
			},
			{
				name: "kiwipete",
				args: args{
					boardInFEN: kiwipete,
					color:      common.White,
					deep:       0,
				},
				want: 1,
			},
			{
				name: "kiwipete",
				args: args{
					boardInFEN: kiwipete,
					color:      common.White,
					deep:       1,
				},
				want: 44,
			},
			{
				name: "kiwipete",
				args: args{
					boardInFEN: kiwipete,
					color:      common.White,
					deep:       2,
				},
				want: 1740,
			},
		} {
			prefix := fmt.Sprintf("%s/%s/%dPly", storage.name, data.name, data.args.deep)
			storage, err := uci.DecodePieceStorage(
				data.args.boardInFEN,
				pieces.NewPiece,
				storage.factory,
			)
			if err != nil {
				test.Errorf("%s: %v", prefix, err)
				continue
			}

			var topLevelMoves []string
			var generator models.MoveGenerator
			got := models.Perft(
				generator,
				storage,
				data.args.color,
				data.args.deep,
				func(move common.Move, count int, deep int) {
					// log only the top-level moves
					if deep == data.args.deep {
						topLevelMoves = append(topLevelMoves, fmt.Sprintf("%v: %d", move, count))
					}
				},
			)

			if got != data.want {
				sort.Strings(topLevelMoves)

				message := "%s: %d/%d\n" + strings.Join(topLevelMoves, "\n")
				test.Errorf(message, prefix, got, data.want)
			}
		}
	}
}
