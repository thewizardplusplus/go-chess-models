package chessmodels

// SliceBoard ...
type SliceBoard struct {
	size   Size
	pieces []Piece
}

// Size ...
func (board SliceBoard) Size() Size {
	panic("not implemented")
}

// Piece ...
func (board SliceBoard) Piece(position Position) (piece Piece, ok bool) {
	panic("not implemented")
}

// Pieces ...
func (board SliceBoard) Pieces() []Piece {
	panic("not implemented")
}

// ApplyMove ...
//
// It doesn't check that the move is correct.
func (board SliceBoard) ApplyMove(move Move) PieceStorage {
	panic("not implemented")
}

// CheckMove ...
//
// It doesn't check for a check before or after the move.
func (board SliceBoard) CheckMove(move Move) error {
	panic("not implemented")
}
