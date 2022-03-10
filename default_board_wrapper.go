package chessmodels

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

// DefaultBoardWrapper ...
type DefaultBoardWrapper struct {
	common.BasePieceStorage
}

// Pieces ...
func (board DefaultBoardWrapper) Pieces() []common.Piece {
	if pieceGroupGetter, ok := board.BasePieceStorage.(common.PieceGroupGetter); ok {
		return pieceGroupGetter.Pieces()
	}

	return common.Pieces(board)
}

// CheckMove ...
//
// It doesn't check for a check before or after the move.
func (board DefaultBoardWrapper) CheckMove(move common.Move) error {
	if moveChecker, ok := board.BasePieceStorage.(common.MoveChecker); ok {
		return moveChecker.CheckMove(move)
	}

	return common.CheckMove(board, move)
}
