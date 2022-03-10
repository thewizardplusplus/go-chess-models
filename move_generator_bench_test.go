package chessmodels_test

import (
	"fmt"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

var (
	initial  = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	kiwipete = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R"
)

func BenchmarkPerft(benchmark *testing.B) {
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
	}

	for _, storage := range []storage{
		{
			name:    "MapBoard",
			factory: boards.NewMapBoard,
		},
		{
			name:    "SliceBoard",
			factory: models.NewSliceBoard,
		},
	} {
		for _, data := range []data{
			{
				name: "initial",
				args: args{
					boardInFEN: initial,
					color:      common.White,
					deep:       1,
				},
			},
			{
				name: "initial",
				args: args{
					boardInFEN: initial,
					color:      common.White,
					deep:       2,
				},
			},
			{
				name: "initial",
				args: args{
					boardInFEN: initial,
					color:      common.White,
					deep:       3,
				},
			},
			{
				name: "kiwipete",
				args: args{
					boardInFEN: kiwipete,
					color:      common.White,
					deep:       1,
				},
			},
			{
				name: "kiwipete",
				args: args{
					boardInFEN: kiwipete,
					color:      common.White,
					deep:       2,
				},
			},
		} {
			prefix := fmt.Sprintf("%s/%s/%dPly", storage.name, data.name, data.args.deep)
			benchmark.Run(prefix, func(benchmark *testing.B) {
				storage, err := uci.DecodePieceStorage(
					data.args.boardInFEN,
					pieces.NewPiece,
					storage.factory,
				)
				if err != nil {
					benchmark.Errorf("%s: %v", prefix, err)
					return
				}

				benchmark.ResetTimer()

				for i := 0; i < benchmark.N; i++ {
					var generator models.MoveGenerator
					models.Perft(generator, storage, data.args.color, data.args.deep, nil)
				}
			})
		}
	}
}
