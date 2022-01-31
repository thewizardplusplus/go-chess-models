package chessmodels_test

import (
	"fmt"
	"sort"

	models "github.com/thewizardplusplus/go-chess-models"
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

func ExampleMapBoard_CheckMove() {
	board := models.NewMapBoard(models.Size{Width: 5, Height: 5}, []models.Piece{
		pieces.NewRook(models.Black, models.Position{File: 2, Rank: 2}),
		pieces.NewBishop(models.White, models.Position{File: 3, Rank: 3}),
	})

	moveOne := models.Move{
		Start:  models.Position{File: 2, Rank: 2},
		Finish: models.Position{File: 3, Rank: 3},
	}
	fmt.Printf("%+v: %v\n", moveOne, board.CheckMove(moveOne))

	moveTwo := models.Move{
		Start:  models.Position{File: 3, Rank: 3},
		Finish: models.Position{File: 2, Rank: 2},
	}
	fmt.Printf("%+v: %v\n", moveTwo, board.CheckMove(moveTwo))

	// Output:
	// {Start:{File:2 Rank:2} Finish:{File:3 Rank:3}}: illegal move
	// {Start:{File:3 Rank:3} Finish:{File:2 Rank:2}}: <nil>
}

func ExampleMapBoard_ApplyMove() {
	board := models.NewMapBoard(models.Size{Width: 5, Height: 5}, []models.Piece{
		pieces.NewRook(models.Black, models.Position{File: 2, Rank: 2}),
		pieces.NewBishop(models.White, models.Position{File: 3, Rank: 3}),
	})
	pieces := board.Pieces()
	sort.Sort(ByPosition(pieces))
	fmt.Printf("%+v\n", pieces)

	updatedBoard := board.ApplyMove(models.Move{
		Start:  models.Position{File: 3, Rank: 3},
		Finish: models.Position{File: 2, Rank: 2},
	})
	updatedPieces := updatedBoard.Pieces()
	sort.Sort(ByPosition(updatedPieces))
	fmt.Printf("%+v\n", updatedPieces)

	// Output:
	// [{Base:{kind:2 color:0 position:{File:2 Rank:2}}} {Base:{kind:3 color:1 position:{File:3 Rank:3}}}]
	// [{Base:{kind:3 color:1 position:{File:2 Rank:2}}}]
}

func ExampleMoveGenerator_MovesForColor() {
	board := models.NewMapBoard(models.Size{Width: 5, Height: 5}, []models.Piece{
		pieces.NewRook(models.Black, models.Position{File: 2, Rank: 2}),
		pieces.NewKnight(models.White, models.Position{File: 3, Rank: 3}),
		pieces.NewPawn(models.White, models.Position{File: 4, Rank: 3}),
	})

	var generator models.MoveGenerator
	moves, _ := generator.MovesForColor(board, models.White)

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
