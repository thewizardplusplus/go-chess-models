package pieces

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func steps(a int, b int) int {
	steps := a - b
	if steps < 0 {
		steps = -steps
	}

	return steps
}

func search(
	storage common.PieceStorage,
	a int,
	b int,
	makePosition func(i int) common.Position,
) (ok bool) {
	start, finish := min(a, b), max(a, b)
	for i := start + 1; i < finish; i++ {
		position := makePosition(i)
		if _, ok := storage.Piece(position); ok {
			return true
		}
	}

	return false
}
