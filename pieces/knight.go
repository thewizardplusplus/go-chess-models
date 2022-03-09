package pieces

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

// Knight ...
type Knight struct{ Base }

// NewKnight ...
func NewKnight(color common.Color, position common.Position) Knight {
	base := NewBase(common.Knight, color, position)
	return Knight{base}
}

// ApplyPosition ...
func (piece Knight) ApplyPosition(position common.Position) common.Piece {
	base := piece.Base.ApplyPosition(position)
	return Knight{base}
}

// CheckMove ...
func (piece Knight) CheckMove(
	move common.Move,
	storage common.PieceStorage,
) bool {
	start, finish := move.Start, move.Finish
	fileSteps := steps(start.File, finish.File)
	rankSteps := steps(start.Rank, finish.Rank)
	return (fileSteps == 1 && rankSteps == 2) || (fileSteps == 2 && rankSteps == 1)
}
