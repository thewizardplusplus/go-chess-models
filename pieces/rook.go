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
	storage models.PieceStorage,
) bool {
	start, finish := move.Start, move.Finish
	fileSteps := steps(start.File, finish.File)
	rankSteps := steps(start.Rank, finish.Rank)
	if fileSteps != 0 && rankSteps != 0 {
		return false
	}

	var ok bool
	switch 0 {
	case fileSteps:
		ok = search(
			storage,
			start.Rank,
			finish.Rank,
			func(i int) models.Position {
				return models.Position{
					File: start.File,
					Rank: i,
				}
			},
		)
	case rankSteps:
		ok = search(
			storage,
			start.File,
			finish.File,
			func(i int) models.Position {
				return models.Position{
					File: i,
					Rank: start.Rank,
				}
			},
		)
	}

	return !ok
}
