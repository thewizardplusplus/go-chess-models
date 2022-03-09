package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/common"
)

// Pawn ...
type Pawn struct{ Base }

// NewPawn ...
func NewPawn(color common.Color, position models.Position) Pawn {
	base := NewBase(common.Pawn, color, position)
	return Pawn{base}
}

// ApplyPosition ...
func (piece Pawn) ApplyPosition(position models.Position) models.Piece {
	base := piece.Base.ApplyPosition(position)
	return Pawn{base}
}

// CheckMove ...
func (piece Pawn) CheckMove(
	move models.Move,
	storage models.PieceStorage,
) bool {
	start, finish := move.Start, move.Finish
	fileSteps := steps(start.File, finish.File)
	if _, ok := storage.Piece(finish); !ok {
		if fileSteps != 0 {
			return false
		}
	} else {
		if fileSteps != 1 {
			return false
		}
	}

	rankSteps := finish.Rank - start.Rank
	switch piece.color {
	case common.Black:
		if rankSteps != -1 {
			return false
		}
	case common.White:
		if rankSteps != 1 {
			return false
		}
	}

	return true
}
