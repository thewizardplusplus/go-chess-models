package pieces

import (
	"math"

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
	step := func(a int, b int) int {
		return int(math.Abs(float64(a - b)))
	}

	start, finish := move.Start, move.Finish
	fileSteps := step(start.File, finish.File)
	rankSteps := step(start.Rank, finish.Rank)
	return fileSteps == 1 && rankSteps == 2 ||
		fileSteps == 2 && rankSteps == 1
}
