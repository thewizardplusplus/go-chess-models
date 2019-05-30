package pieces

import (
	"math"

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
	board models.Board,
) bool {
	min := func(a int, b int) int {
		fa, fb := float64(a), float64(b)
		return int(math.Min(fa, fb))
	}
	max := func(a int, b int) int {
		fa, fb := float64(a), float64(b)
		return int(math.Max(fa, fb))
	}

	type p func(i int) models.Position
	search := func(a int, b int, p p) bool {
		start := min(a, b)
		finish := max(a, b)
		for i := start + 1; i < finish; i++ {
			position := p(i)
			_, ok := board.Piece(position)
			if ok {
				return true
			}
		}

		return false
	}

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
