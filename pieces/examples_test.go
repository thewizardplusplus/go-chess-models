package pieces_test

import (
	"fmt"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func ExampleKnight_CheckMove() {
	board := models.NewBoard(models.Size{Width: 5, Height: 5}, []models.Piece{
		pieces.NewKnight(models.Black, models.Position{File: 2, Rank: 2}),
	})

	moveOne := models.Move{
		Start:  models.Position{File: 2, Rank: 2},
		Finish: models.Position{File: 2, Rank: 3},
	}
	fmt.Printf("%+v: %v\n", moveOne, board.CheckMove(moveOne))

	moveTwo := models.Move{
		Start:  models.Position{File: 2, Rank: 2},
		Finish: models.Position{File: 4, Rank: 3},
	}
	fmt.Printf("%+v: %v\n", moveTwo, board.CheckMove(moveTwo))

	// Output:
	// {Start:{File:2 Rank:2} Finish:{File:2 Rank:3}}: illegal move
	// {Start:{File:2 Rank:2} Finish:{File:4 Rank:3}}: <nil>
}

func ExampleKnight_ApplyPosition() {
	piece := pieces.NewKnight(models.Black, models.Position{File: 2, Rank: 2})
	fmt.Printf("%+v\n", piece)

	updatedPiece := piece.ApplyPosition(models.Position{File: 4, Rank: 3})
	fmt.Printf("%+v\n", updatedPiece)

	// Output:
	// {Base:{kind:4 color:0 position:{File:2 Rank:2}}}
	// {Base:{kind:4 color:0 position:{File:4 Rank:3}}}
}
