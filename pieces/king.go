package pieces

import (
	"math"

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
	oneStep := func(a int, b int) bool {
		return math.Abs(float64(a-b)) <= 1
	}

	start, finish := move.Start, move.Finish
	return oneStep(start.File, finish.File) &&
		oneStep(start.Rank, finish.Rank)
}
