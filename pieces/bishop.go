package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

// Bishop ...
type Bishop struct{ Base }

// NewBishop ...
func NewBishop(
	color models.Color,
	position models.Position,
) Bishop {
	kind := models.Bishop
	base := Base{kind, color, position}
	return Bishop{base}
}

// ApplyPosition ...
func (piece Bishop) ApplyPosition(
	position models.Position,
) models.Piece {
	base := piece.Base.ApplyPosition(position)
	return Bishop{base}
}

// CheckMove ...
func (piece Bishop) CheckMove(
	move models.Move,
	board models.Board,
) bool {
	return false
}
