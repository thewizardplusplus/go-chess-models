package chessmodels

// Size ...
type Size struct {
	Width  int
	Height int
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
