package chessmodels

// DefaultBoardWrapper ...
type DefaultBoardWrapper struct {
	BasePieceStorage
}

// Pieces ...
func (board DefaultBoardWrapper) Pieces() []Piece {
	if pieceGroupGetter, ok := board.BasePieceStorage.(PieceGroupGetter); ok {
		return pieceGroupGetter.Pieces()
	}

	return Pieces(board)
}

// CheckMove ...
//
// It doesn't check for a check before or after the move.
func (board DefaultBoardWrapper) CheckMove(move Move) error {
	return CheckMove(board, move)
}
