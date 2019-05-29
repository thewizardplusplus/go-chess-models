package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

// Knight ...
type Knight struct{ Base }

// NewKnight ...
func NewKnight(
	color models.Color,
	position models.Position,
) Knight {
	kind := models.Knight
	base := Base{kind, color, position}
	return Knight{base}
}

// ApplyPosition ...
func (piece Knight) ApplyPosition(
	position models.Position,
) models.Piece {
	base := piece.Base.ApplyPosition(position)
	return Knight{base}
}

// CheckMove ...
func (piece Knight) CheckMove(
	move models.Move,
	board models.Board,
) bool {
	return false
}
