package chessmodels_test

import (
	"fmt"
	"sort"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

type MockPiece struct {
	kind      models.Kind
	color     models.Color
	position  models.Position
	checkMove func(move models.Move, storage models.PieceStorage) bool
}

func (piece MockPiece) Kind() models.Kind {
	return piece.kind
}

func (piece MockPiece) Color() models.Color {
	return piece.color
}

func (piece MockPiece) Position() models.Position {
	return piece.position
}

func (piece MockPiece) ApplyPosition(position models.Position) models.Piece {
	return MockPiece{
		kind:      piece.kind,
		color:     piece.color,
		position:  position,
		checkMove: piece.checkMove,
	}
}

func (piece MockPiece) CheckMove(
	move models.Move,
	storage models.PieceStorage,
) bool {
	if piece.checkMove == nil {
		panic("not implemented")
	}

	return piece.checkMove(move, storage)
}

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
	a := group[i].Position()
	b := group[j].Position()
	if a.File == b.File {
		return a.Rank < b.Rank
	}

	return a.File < b.File
}

func ExampleBoard_CheckMove() {
	board := models.NewBoard(models.Size{Width: 5, Height: 5}, []models.Piece{
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

func ExampleBoard_CheckMove_withMockPieces() {
	board := models.NewBoard(models.Size{Width: 5, Height: 5}, []models.Piece{
		MockPiece{
			kind:     models.Rook,
			color:    models.Black,
			position: models.Position{File: 2, Rank: 2},
			checkMove: func(move models.Move, storage models.PieceStorage) bool {
				return false
			},
		},
		MockPiece{
			kind:     models.Bishop,
			color:    models.White,
			position: models.Position{File: 3, Rank: 3},
			checkMove: func(move models.Move, storage models.PieceStorage) bool {
				return true
			},
		},
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

func ExampleBoard_ApplyMove() {
	board := models.NewBoard(models.Size{Width: 5, Height: 5}, []models.Piece{
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

func ExampleBoard_ApplyMove_withMockPieces() {
	board := models.NewBoard(models.Size{Width: 5, Height: 5}, []models.Piece{
		MockPiece{
			kind:     models.Rook,
			position: models.Position{File: 2, Rank: 2},
		},
		MockPiece{
			kind:     models.Bishop,
			position: models.Position{File: 3, Rank: 3},
		},
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
	// [{kind:2 color:0 position:{File:2 Rank:2} checkMove:<nil>} {kind:3 color:0 position:{File:3 Rank:3} checkMove:<nil>}]
	// [{kind:3 color:0 position:{File:2 Rank:2} checkMove:<nil>}]
}
