package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

// King ...
type King struct{ Base }

// NewKing ...
func NewKing(
	color models.Color,
	position models.Position,
) King {
	kind := models.King
	base := Base{kind, color, position}
	return King{base}
}

// ApplyPosition ...
func (piece King) ApplyPosition(
	position models.Position,
) models.Piece {
	base := piece.Base.ApplyPosition(position)
	return King{base}
}

// CheckMove ...
func (piece King) CheckMove(
	move models.Move,
	board models.Board,
) bool {
	return false
}
