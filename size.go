package chessmodels

// Size ...
type Size struct {
	Width  int
	Height int
}

// HasPosition ...
func (size Size) HasPosition(position Position) bool {
	return less(position.File, size.Width) && less(position.Rank, size.Height)
}

// HasMove ...
func (size Size) HasMove(move Move) bool {
	return size.HasPosition(move.Start) && size.HasPosition(move.Finish)
}

// Positions ...
func (size Size) Positions() []Position {
	var positions []Position
	for rank := 0; rank < size.Height; rank++ {
		for file := 0; file < size.Width; file++ {
			positions = append(positions, Position{file, rank})
		}
	}

	return positions
}

func less(value int, limit int) bool {
	return 0 <= value && value < limit
}
