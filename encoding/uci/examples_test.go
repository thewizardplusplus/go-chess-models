package uci_test

import (
	"fmt"
	"sort"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

type ByPosition []models.Piece

func (group ByPosition) Len() int {
	return len(group)
}

func (group ByPosition) Swap(i, j int) {
	group[i], group[j] = group[j], group[i]
}

func (group ByPosition) Less(
	i, j int,
) bool {
	a, b := group[i].Position(), group[j].Position()
	if a.File == b.File {
		return a.Rank < b.Rank
	}

	return a.File < b.File
}

func ExampleDecodeMove() {
	move, _ := uci.DecodeMove("d4c3")
	fmt.Printf("%+v\n", move)

	// Output: {Start:{File:3 Rank:3} Finish:{File:2 Rank:2}}
}

func ExampleEncodeMove() {
	move := uci.EncodeMove(common.Move{
		Start:  common.Position{File: 3, Rank: 3},
		Finish: common.Position{File: 2, Rank: 2},
	})
	fmt.Printf("%v\n", move)

	// Output: d4c3
}

func ExampleDecodePieceStorage() {
	const fen = "8/8/8/8/3B4/2r5/8/8"
	storage, _ := uci.DecodePieceStorage(fen, pieces.NewPiece, models.NewBoard)
	pieces := storage.Pieces()
	sort.Sort(ByPosition(pieces))
	fmt.Printf("%+v\n", pieces)

	// Output: [{Base:{kind:2 color:0 position:{File:2 Rank:2}}} {Base:{kind:3 color:1 position:{File:3 Rank:3}}}]
}

func ExampleEncodePieceStorage() {
	board := models.NewBoard(common.Size{Width: 5, Height: 5}, []models.Piece{
		pieces.NewRook(common.Black, common.Position{File: 2, Rank: 2}),
		pieces.NewBishop(common.White, common.Position{File: 3, Rank: 3}),
	})
	fen := uci.EncodePieceStorage(board)
	fmt.Printf("%v\n", fen)

	// Output: 5/3B1/2r2/5/5
}
