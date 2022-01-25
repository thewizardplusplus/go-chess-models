package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

// Rook ...
type Rook struct{ Base }

// NewRook ...
func NewRook(color models.Color, position models.Position) Rook {
	base := NewBase(models.Rook, color, position)
	return Rook{base}
}

// ApplyPosition ...
func (piece Rook) ApplyPosition(position models.Position) models.Piece {
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
	var makePosition func(i int) models.Position
	if fileSteps == 0 {
		a, b = start.Rank, finish.Rank
		makePosition = func(i int) models.Position {
			return models.Position{
				File: start.File,
				Rank: i,
			}
		}
	} else {
		a, b = start.File, finish.File
		makePosition = func(i int) models.Position {
			return models.Position{
				File: i,
				Rank: start.Rank,
			}
		}
	}

	return !search(storage, a, b, makePosition)
}
