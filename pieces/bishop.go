package pieces

import (
	"math"

	models "github.com/thewizardplusplus/go-chess-models"
)

// Bishop ...
type Bishop struct{ Base }

// NewBishop ...
func NewBishop(
	color models.Color,
	position models.Position,
) Bishop {
	kind := models.Bishop
	base := Base{kind, color, position}
	return Bishop{base}
}

// ApplyPosition ...
func (piece Bishop) ApplyPosition(
	position models.Position,
) models.Piece {
	base := piece.Base.ApplyPosition(position)
	return Bishop{base}
}

// CheckMove ...
func (piece Bishop) CheckMove(
	move models.Move,
	board models.Board,
) bool {
	min := func(a int, b int) int {
		fa, fb := float64(a), float64(b)
		return int(math.Min(fa, fb))
	}

	start, finish := move.Start, move.Finish
	fileSteps := steps(start.File, finish.File)
	rankSteps := steps(start.Rank, finish.Rank)
	if fileSteps != rankSteps {
		return false
	}

	fileMin := min(start.File, finish.File)
	rankMin := min(start.Rank, finish.Rank)
	ok := search(
		board,
		start.File,
		finish.File,
		func(i int) models.Position {
			step := i - fileMin
			return models.Position{
				File: i,
				Rank: rankMin + step,
			}
		},
	)
	return !ok
}
