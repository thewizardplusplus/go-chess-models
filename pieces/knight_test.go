package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestNewKnight(test *testing.T) {
	piece := NewKnight(
		models.White,
		models.Position{File: 2, Rank: 3},
	)

	expectedPiece := Knight{
		Base: Base{
			kind:  models.Knight,
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

func TestKnightApplyPosition(
	test *testing.T,
) {
	piece := NewKnight(
		models.White,
		models.Position{File: 2, Rank: 3},
	)
	nextPiece := piece.ApplyPosition(
		models.Position{
			File: 4,
			Rank: 2,
		},
	)

	expectedPiece := Knight{
		Base: Base{
			kind:  models.Knight,
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

	expectedNextPiece := Knight{
		Base: Base{
			kind:  models.Knight,
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

func TestKnightCheckMove(test *testing.T) {
	storage, err := models.ParseBoard(
		"5/5/2N2/5/5",
		func(fen rune) (models.Piece, error) {
			return ParsePiece(
				fen,
				func(
					kind models.Kind,
					color models.Color,
				) models.Piece {
					return NewPiece(
						kind,
						color,
						models.Position{},
					)
				},
			)
		},
	)
	if err != nil {
		test.Fail()
		return
	}

	generator := models.MoveGenerator{}
	moves, err := generator.MovesForPosition(
		storage,
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
				Rank: 0,
			},
		},
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 3,
				Rank: 0,
			},
		},
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 0,
				Rank: 1,
			},
		},
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 4,
				Rank: 1,
			},
		},
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 0,
				Rank: 3,
			},
		},
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 4,
				Rank: 3,
			},
		},
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 1,
				Rank: 4,
			},
		},
		models.Move{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 3,
				Rank: 4,
			},
		},
	}
	if !reflect.DeepEqual(
		moves,
		expectedMoves,
	) {
		test.Fail()
	}
	if err != nil {
		test.Fail()
	}
}
