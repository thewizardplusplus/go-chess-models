package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

// Pawn ...
type Pawn struct{ Base }

// NewPawn ...
func NewPawn(
	color models.Color,
	position models.Position,
) Pawn {
	kind := models.Pawn
	base := Base{kind, color, position}
	return Pawn{base}
}

// ApplyPosition ...
func (piece Pawn) ApplyPosition(
	position models.Position,
) models.Piece {
	base := piece.Base.ApplyPosition(position)
	return Pawn{base}
}

// CheckMove ...
func (piece Pawn) CheckMove(
	move models.Move,
	board models.Board,
) bool {
	return false
}
