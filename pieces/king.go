package pieces

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

// King ...
type King struct{ Base }

// NewKing ...
func NewKing(color common.Color, position common.Position) King {
	base := NewBase(common.King, color, position)
	return King{base}
}

// ApplyPosition ...
func (piece King) ApplyPosition(position common.Position) common.Piece {
	base := piece.Base.ApplyPosition(position)
	return King{base}
}

// CheckMove ...
func (piece King) CheckMove(
	move common.Move,
	storage common.PieceStorage,
) bool {
	start, finish := move.Start, move.Finish
	fileSteps := steps(start.File, finish.File)
	rankSteps := steps(start.Rank, finish.Rank)
	return fileSteps <= 1 && rankSteps <= 1
}
