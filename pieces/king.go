package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

// King ...
type King struct{ base }

// NewKing ...
func NewKing(
	color models.Color,
	position models.Position,
) King {
	kind := models.King
	base := base{kind, color, position}
	return King{base}
}

// ApplyPosition ...
func (piece King) ApplyPosition(
	position models.Position,
) models.Piece {
	base := piece.base.ApplyPosition(position)
	return King{base}
}

// CheckMove ...
func (piece King) CheckMove(
	move models.Move,
	storage models.PieceStorage,
) bool {
	start, finish := move.Start, move.Finish
	fileSteps := steps(start.File, finish.File)
	rankSteps := steps(start.Rank, finish.Rank)
	return fileSteps <= 1 && rankSteps <= 1
}
