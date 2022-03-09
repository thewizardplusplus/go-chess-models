package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/common"
)

// Rook ...
type Rook struct{ Base }

// NewRook ...
func NewRook(color common.Color, position common.Position) Rook {
	base := NewBase(common.Rook, color, position)
	return Rook{base}
}

// ApplyPosition ...
func (piece Rook) ApplyPosition(position common.Position) models.Piece {
	base := piece.Base.ApplyPosition(position)
	return Rook{base}
}

// CheckMove ...
func (piece Rook) CheckMove(
	move models.Move,
	storage models.PieceStorage,
) bool {
	start, finish := move.Start, move.Finish
	fileSteps := steps(start.File, finish.File)
	rankSteps := steps(start.Rank, finish.Rank)
	if fileSteps != 0 && rankSteps != 0 {
		return false
	}

	var a, b int
	var makePosition func(i int) common.Position
	if fileSteps == 0 {
		a, b = start.Rank, finish.Rank
		makePosition = func(i int) common.Position {
			return common.Position{
				File: start.File,
				Rank: i,
			}
		}
	} else {
		a, b = start.File, finish.File
		makePosition = func(i int) common.Position {
			return common.Position{
				File: i,
				Rank: start.Rank,
			}
		}
	}

	return !search(storage, a, b, makePosition)
}
