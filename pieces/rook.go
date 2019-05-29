package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

// Rook ...
type Rook struct{ Base }

// NewRook ...
func NewRook(
	color models.Color,
	position models.Position,
) Rook {
	kind := models.Rook
	base := Base{kind, color, position}
	return Rook{base}
}

// ApplyPosition ...
func (piece Rook) ApplyPosition(
	position models.Position,
) models.Piece {
	base := piece.Base.ApplyPosition(position)
	return Rook{base}
}

// CheckMove ...
func (piece Rook) CheckMove(
	move models.Move,
	board models.Board,
) bool {
	return false
}
