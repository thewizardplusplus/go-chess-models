package chessmodels

// DefaultBoardWrapper ...
type DefaultBoardWrapper struct {
	BasePieceStorage
}

// Pieces ...
func (board DefaultBoardWrapper) Pieces() []Piece {
	return Pieces(board)
}

// CheckMove ...
//
// It doesn't check for a check before or after the move.
func (board DefaultBoardWrapper) CheckMove(move Move) error {
	return CheckMove(board, move)
}
