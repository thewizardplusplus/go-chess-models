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
			positions = append(
				positions,
				position,
			)
		}
	}

	return positions
}

// MovesForPosition ...
func (size Size) MovesForPosition(
	start Position,
) []Move {
	var moves []Move
	for _, finish := range size.Positions() {
		move := Move{start, finish}
		moves = append(moves, move)
	}

	return moves
}

// Moves ...
func (size Size) Moves() MoveGroup {
	moves := make(MoveGroup)
	for _, position := range size.Positions() {
		moves[position] =
			size.MovesForPosition(position)
	}

	return moves
}
