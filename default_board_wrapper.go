package chessmodels

type defaultBoardWrapper struct {
	BasePieceStorage
}

func (board defaultBoardWrapper) Pieces() []Piece {
	return Pieces(board)
}

// It doesn't check for a check before or after the move.
func (board defaultBoardWrapper) CheckMove(move Move) error {
	return CheckMove(board, move)
}
