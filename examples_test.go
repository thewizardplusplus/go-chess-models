package chessmodels_test

import (
	"fmt"
	"sort"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func ExampleMoveGenerator_MovesForColor() {
	board := boards.NewMapBoard(common.Size{Width: 5, Height: 5}, []common.Piece{
		pieces.NewRook(common.Black, common.Position{File: 2, Rank: 2}),
		pieces.NewKnight(common.White, common.Position{File: 3, Rank: 3}),
		pieces.NewPawn(common.White, common.Position{File: 4, Rank: 3}),
	})

	var generator models.MoveGenerator
	moves, _ := generator.MovesForColor(board, common.White)

	// sorting only by the final point will be sufficient for the reproducibility
	// of this example
	sort.Slice(moves, func(i int, j int) bool {
		a, b := moves[i].Finish, moves[j].Finish
		if a.File == b.File {
			return a.Rank < b.Rank
		}

		return a.File < b.File
	})

	for _, move := range moves {
		fmt.Printf("%+v\n", move)
	}

	// Output:
	// {Start:{File:3 Rank:3} Finish:{File:1 Rank:2}}
	// {Start:{File:3 Rank:3} Finish:{File:1 Rank:4}}
	// {Start:{File:3 Rank:3} Finish:{File:2 Rank:1}}
	// {Start:{File:3 Rank:3} Finish:{File:4 Rank:1}}
	// {Start:{File:4 Rank:3} Finish:{File:4 Rank:4}}
}
