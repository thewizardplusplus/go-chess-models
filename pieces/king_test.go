package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
)

func TestNewKing(test *testing.T) {
	piece := NewKing(common.White, common.Position{
		File: 2,
		Rank: 3,
	})

	expectedPiece := King{
		Base: Base{
			kind:  common.King,
			color: common.White,
			position: common.Position{
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
	piece := NewKing(common.White, common.Position{
		File: 2,
		Rank: 3,
	})
	nextPiece := piece.ApplyPosition(common.Position{
		File: 4,
		Rank: 2,
	})

	expectedPiece := King{
		Base: Base{
			kind:  common.King,
			color: common.White,
			position: common.Position{
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
			kind:  common.King,
			color: common.White,
			position: common.Position{
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
		uci.DecodePieceStorage("5/5/2K2/5/5", NewPiece, models.NewBoard)
	if err != nil {
		test.Fail()
		return
	}

	var generator models.MoveGenerator
	moves, err := generator.MovesForPosition(storage, common.Position{
		File: 2,
		Rank: 2,
	})

	expectedMoves := []models.Move{
		{
			Start: common.Position{
				File: 2,
				Rank: 2,
			},
			Finish: common.Position{
				File: 1,
				Rank: 1,
			},
		},
		{
			Start: common.Position{
				File: 2,
				Rank: 2,
			},
			Finish: common.Position{
				File: 2,
				Rank: 1,
			},
		},
		{
			Start: common.Position{
				File: 2,
				Rank: 2,
			},
			Finish: common.Position{
				File: 3,
				Rank: 1,
			},
		},
		{
			Start: common.Position{
				File: 2,
				Rank: 2,
			},
			Finish: common.Position{
				File: 1,
				Rank: 2,
			},
		},
		{
			Start: common.Position{
				File: 2,
				Rank: 2,
			},
			Finish: common.Position{
				File: 3,
				Rank: 2,
			},
		},
		{
			Start: common.Position{
				File: 2,
				Rank: 2,
			},
			Finish: common.Position{
				File: 1,
				Rank: 3,
			},
		},
		{
			Start: common.Position{
				File: 2,
				Rank: 2,
			},
			Finish: common.Position{
				File: 2,
				Rank: 3,
			},
		},
		{
			Start: common.Position{
				File: 2,
				Rank: 2,
			},
			Finish: common.Position{
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
