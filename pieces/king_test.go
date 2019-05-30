package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestKingInterface(test *testing.T) {
	kingType := reflect.TypeOf(King{})
	pieceType := reflect.
		TypeOf((*models.Piece)(nil)).
		Elem()
	if !kingType.Implements(pieceType) {
		test.Fail()
	}
}

func TestNewKing(test *testing.T) {
	piece := NewKing(
		models.White,
		models.Position{File: 2, Rank: 3},
	)

	expectedPiece := King{
		Base: Base{
			kind:  models.King,
			color: models.White,
			position: models.Position{
				File: 2,
				Rank: 3,
			},
		},
	}
	if !reflect.DeepEqual(
		piece,
		expectedPiece,
	) {
		test.Fail()
	}
}

func TestKingApplyPosition(
	test *testing.T,
) {
	piece := NewKing(
		models.White,
		models.Position{File: 2, Rank: 3},
	)
	nextPiece := piece.ApplyPosition(
		models.Position{
			File: 4,
			Rank: 2,
		},
	)

	expectedPiece := King{
		Base: Base{
			kind:  models.King,
			color: models.White,
			position: models.Position{
				File: 2,
				Rank: 3,
			},
		},
	}
	if !reflect.DeepEqual(
		piece,
		expectedPiece,
	) {
		test.Fail()
	}

	expectedNextPiece := King{
		Base: Base{
			kind:  models.King,
			color: models.White,
			position: models.Position{
				File: 4,
				Rank: 2,
			},
		},
	}
	if !reflect.DeepEqual(
		nextPiece,
		expectedNextPiece,
	) {
		test.Fail()
	}
}

func TestKingCheckMove(test *testing.T) {
	board := models.NewBoard(
		models.Size{5, 5},
		[]models.Piece{
			NewKing(
				models.White,
				models.Position{File: 2, Rank: 2},
			),
		},
	)
	generator := models.MoveGenerator{
		PieceStorage: board,
	}
	moves := generator.MovesForPosition(
		models.Position{File: 2, Rank: 2},
	)

	expectedMoves := []models.Move{
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 1,
				Rank: 1,
			},
		},
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 2,
				Rank: 1,
			},
		},
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 3,
				Rank: 1,
			},
		},
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 1,
				Rank: 2,
			},
		},
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 3,
				Rank: 2,
			},
		},
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 1,
				Rank: 3,
			},
		},
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 2,
				Rank: 3,
			},
		},
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 3,
				Rank: 3,
			},
		},
	}
	if !reflect.DeepEqual(
		moves,
		expectedMoves,
	) {
		test.Fail()
	}
}
