package pieces

import (
	"math"

	models "github.com/thewizardplusplus/go-chess-models"
)

func min(a int, b int) int {
	fa, fb := float64(a), float64(b)
	return int(math.Min(fa, fb))
}

func max(a int, b int) int {
	fa, fb := float64(a), float64(b)
	return int(math.Max(fa, fb))
}

func steps(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}

func search(
	storage models.PieceStorage,
	a int,
	b int,
	makePosition func(i int) models.Position,
) (ok bool) {
	start := min(a, b)
	finish := max(a, b)
	for i := start + 1; i < finish; i++ {
		position := makePosition(i)
		if _, ok := storage.Piece(position); ok {
			return true
		}
	}

	return false
}
