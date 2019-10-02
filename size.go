package chessmodels

// Size ...
type Size struct {
	Width  int
	Height int
}

// HasPosition ...
func (size Size) HasPosition(
	position Position,
) bool {
	return isLess(position.File, size.Width) &&
		isLess(position.Rank, size.Height)
}

// HasMove ...
func (size Size) HasMove(move Move) bool {
	return size.HasPosition(move.Start) &&
		size.HasPosition(move.Finish)
}

// Positions ...
func (size Size) Positions() []Position {
	var positions []Position
	width, height := size.Width, size.Height
	for rank := 0; rank < height; rank++ {
		for file := 0; file < width; file++ {
			position := Position{file, rank}
			positions = append(positions, position)
		}
	}

	return positions
}

func isLess(value int, limit int) bool {
	return value >= 0 && value < limit
}
