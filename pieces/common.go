package pieces

import (
	"math"

	models "github.com/thewizardplusplus/go-chess-models"
)

func steps(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}

func search(
	board models.Board,
	a int,
	b int,
	makePosition func(i int) models.Position,
) bool {
	fa, fb := float64(a), float64(b)
	start := int(math.Min(fa, fb))
	finish := int(math.Max(fa, fb))
	for i := start + 1; i < finish; i++ {
		position := makePosition(i)
		_, ok := board.Piece(position)
		if ok {
			return true
		}
	}

	return false
}
