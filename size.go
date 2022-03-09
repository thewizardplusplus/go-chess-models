package chessmodels

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

// PositionHandler ...
type PositionHandler func(position common.Position) error

// Size ...
type Size struct {
	Width  int
	Height int
}

// HasPosition ...
func (size Size) HasPosition(position common.Position) bool {
	return less(position.File, size.Width) && less(position.Rank, size.Height)
}

// HasMove ...
func (size Size) HasMove(move common.Move) bool {
	return size.HasPosition(move.Start) && size.HasPosition(move.Finish)
}

// PositionIndex ...
func (size Size) PositionIndex(position common.Position) int {
	return size.Width*position.Rank + position.File
}

// PositionCount ...
func (size Size) PositionCount() int {
	return size.Width * size.Height
}

// Positions ...
func (size Size) Positions() []common.Position {
	positions := make([]common.Position, 0, size.PositionCount())
	size.IteratePositions(func(position common.Position) error { // nolint: errcheck, gosec, lll
		positions = append(positions, position)
		return nil
	})

	return positions
}

// IteratePositions ...
func (size Size) IteratePositions(handler PositionHandler) error {
	for rank := 0; rank < size.Height; rank++ {
		for file := 0; file < size.Width; file++ {
			position := common.Position{file, rank}
			if err := handler(position); err != nil {
				return err
			}
		}
	}

	return nil
}

func less(value int, limit int) bool {
	return value >= 0 && value < limit
}
