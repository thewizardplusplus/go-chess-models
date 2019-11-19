package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

// Knight ...
type Knight struct{ Base }

// NewKnight ...
func NewKnight(color models.Color, position models.Position) Knight {
	base := NewBase(models.Knight, color, position)
	return Knight{base}
}

// ApplyPosition ...
func (piece Knight) ApplyPosition(position models.Position) models.Piece {
	base := piece.Base.ApplyPosition(position)
	return Knight{base}
}

// CheckMove ...
func (piece Knight) CheckMove(
	move models.Move,
	storage models.PieceStorage,
) bool {
	start, finish := move.Start, move.Finish
	fileSteps := steps(start.File, finish.File)
	rankSteps := steps(start.Rank, finish.Rank)
	return fileSteps == 1 && rankSteps == 2 || fileSteps == 2 && rankSteps == 1
}
