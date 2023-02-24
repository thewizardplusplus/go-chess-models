package boards_test

import (
	"fmt"
	"sort"

	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

type ByPosition []common.Piece

func (group ByPosition) Len() int {
	return len(group)
}

func (group ByPosition) Swap(i int, j int) {
	group[i], group[j] = group[j], group[i]
}

func (group ByPosition) Less(i int, j int) bool {
	a, b := group[i].Position(), group[j].Position()
	if a.File == b.File {
		return a.Rank < b.Rank
	}

	return a.File < b.File
}

func ExampleMapBoard_CheckMove() { // nolint: govet
	board := boards.NewMapBoard(common.Size{Width: 5, Height: 5}, []common.Piece{
		pieces.NewRook(common.Black, common.Position{File: 2, Rank: 2}),
		pieces.NewBishop(common.White, common.Position{File: 3, Rank: 3}),
	})

	moveOne := common.Move{
		Start:  common.Position{File: 2, Rank: 2},
		Finish: common.Position{File: 3, Rank: 3},
	}
	fmt.Printf("%+v: %v\n", moveOne, board.CheckMove(moveOne))

	moveTwo := common.Move{
		Start:  common.Position{File: 3, Rank: 3},
		Finish: common.Position{File: 2, Rank: 2},
	}
	fmt.Printf("%+v: %v\n", moveTwo, board.CheckMove(moveTwo))

	// Output:
	// {Start:{File:2 Rank:2} Finish:{File:3 Rank:3}}: illegal move
	// {Start:{File:3 Rank:3} Finish:{File:2 Rank:2}}: <nil>
}

func ExampleMapBoard_ApplyMove() {
	board := boards.NewMapBoard(common.Size{Width: 5, Height: 5}, []common.Piece{
		pieces.NewRook(common.Black, common.Position{File: 2, Rank: 2}),
		pieces.NewBishop(common.White, common.Position{File: 3, Rank: 3}),
	})
	pieces := board.Pieces()
	sort.Sort(ByPosition(pieces))
	fmt.Printf("%+v\n", pieces)

	updatedBoard := board.ApplyMove(common.Move{
		Start:  common.Position{File: 3, Rank: 3},
		Finish: common.Position{File: 2, Rank: 2},
	})
	updatedPieces := updatedBoard.Pieces()
	sort.Sort(ByPosition(updatedPieces))
	fmt.Printf("%+v\n", updatedPieces)

	// Output:
	// [{Base:{kind:2 color:0 position:{File:2 Rank:2}}} {Base:{kind:3 color:1 position:{File:3 Rank:3}}}]
	// [{Base:{kind:3 color:1 position:{File:2 Rank:2}}}]
}
