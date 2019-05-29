package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

// Queen ...
type Queen struct{ Base }

// NewQueen ...
func NewQueen(
	color models.Color,
	position models.Position,
) Queen {
	kind := models.Queen
	base := Base{kind, color, position}
	return Queen{base}
}

// ApplyPosition ...
func (piece Queen) ApplyPosition(
	position models.Position,
) models.Piece {
	base := piece.Base.ApplyPosition(position)
	return Queen{base}
}

// CheckMove ...
func (piece Queen) CheckMove(
	move models.Move,
	board models.Board,
) bool {
	return false
}
