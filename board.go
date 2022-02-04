package chessmodels

// Board ...
type Board = SliceBoard

// NewBoard ...
func NewBoard(size Size, pieces []Piece) PieceStorage {
	return NewSliceBoard(size, pieces)
}
