package chessmodels

// Size ...
type Size struct {
	Width  int
	Height int
}

// Board ...
type Board struct {
	size   Size
	pieces PieceGroup
}

// NewBoard ...
func NewBoard(
	size Size,
	pieces PieceGroup,
) Board {
	return Board{size, pieces}
}

// ApplyMove ...
// It doesn't check that the move is correct.
func (board Board) ApplyMove(
	move Move,
) Board {
	pieces := board.pieces.Copy()
	pieces.Move(move)

	return NewBoard(board.size, pieces)
}
