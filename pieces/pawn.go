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
	storage models.PieceStorage,
) bool {
	start, finish := move.Start, move.Finish

	_, ok := storage.Piece(finish)
	fileSteps := steps(start.File, finish.File)
	switch ok {
	case false:
		if fileSteps != 0 {
			return false
		}
	case true:
		if fileSteps != 1 {
			return false
		}
	}

	rankSteps := finish.Rank - start.Rank
	switch piece.color {
	case models.Black:
		if rankSteps != -1 {
			return false
		}
	case models.White:
		if rankSteps != 1 {
			return false
		}
	}

	return true
}
