package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/common"
)

// Bishop ...
type Bishop struct{ Base }

// NewBishop ...
func NewBishop(color common.Color, position common.Position) Bishop {
	base := NewBase(common.Bishop, color, position)
	return Bishop{base}
}

// ApplyPosition ...
func (piece Bishop) ApplyPosition(position common.Position) models.Piece {
	base := piece.Base.ApplyPosition(position)
	return Bishop{base}
}

// CheckMove ...
func (piece Bishop) CheckMove(
	move models.Move,
	storage models.PieceStorage,
) bool {
	start, finish := move.Start, move.Finish
	fileSteps := steps(start.File, finish.File)
	rankSteps := steps(start.Rank, finish.Rank)
	if fileSteps != rankSteps {
		return false
	}

	// if file in the move are descending, these will be scanned from a finish
	// to a start (see the search() function)
	//
	// scanning direction of ranks should correspond to it
	rankStart, rankFinish := start.Rank, finish.Rank
	if start.File > finish.File {
		rankStart, rankFinish = rankFinish, rankStart
	}

	rankSign := 1
	if rankStart > rankFinish {
		rankSign = -1
	}

	fileMin := min(start.File, finish.File)
	return !search(storage, start.File, finish.File, func(i int) common.Position {
		step := i - fileMin
		return common.Position{
			File: i,
			Rank: rankStart + step*rankSign,
		}
	})
}
