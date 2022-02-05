package chessmodels

// BaseBoard ...
type BaseBoard struct {
	size Size
}

// NewBaseBoard ...
func NewBaseBoard(size Size) BaseBoard {
	return BaseBoard{size}
}

// Size ...
func (board BaseBoard) Size() Size {
	return board.size
}
