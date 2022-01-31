package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
)

func TestNewKing(test *testing.T) {
	piece := NewKing(models.White, models.Position{
		File: 2,
		Rank: 3,
	})

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
	if !reflect.DeepEqual(piece, expectedPiece) {
		test.Fail()
	}
}

func TestKingApplyPosition(test *testing.T) {
	piece := NewKing(models.White, models.Position{
		File: 2,
		Rank: 3,
	})
	nextPiece := piece.ApplyPosition(models.Position{
		File: 4,
		Rank: 2,
	})

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
	if !reflect.DeepEqual(piece, expectedPiece) {
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
	if !reflect.DeepEqual(nextPiece, expectedNextPiece) {
		test.Fail()
	}
}

func TestKingCheckMove(test *testing.T) {
	storage, err :=
		uci.DecodePieceStorage("5/5/2K2/5/5", NewPiece, models.NewMapBoard)
	if err != nil {
		test.Fail()
		return
	}

	var generator models.MoveGenerator
	moves, err := generator.MovesForPosition(storage, models.Position{
		File: 2,
		Rank: 2,
	})

	expectedMoves := []models.Move{
		{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 1,
				Rank: 1,
			},
		},
		{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 2,
				Rank: 1,
			},
		},
		{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 3,
				Rank: 1,
			},
		},
		{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 1,
				Rank: 2,
			},
		},
		{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 3,
				Rank: 2,
			},
		},
		{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 1,
				Rank: 3,
			},
		},
		{
			Start: models.Position{
				File: 2,
				Rank: 2,
			},
			Finish: models.Position{
				File: 2,
				Rank: 3,
			},
		},
		{
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
	if !reflect.DeepEqual(moves, expectedMoves) {
		test.Fail()
	}
	if err != nil {
		test.Fail()
	}
}
