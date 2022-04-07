package boards

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

type pieceStorageWithoutPieceGroupGetter interface {
	common.BasePieceStorage
	common.MoveChecker
}

type pieceGroupGetterWrapper struct {
	pieceStorageWithoutPieceGroupGetter
}

// Pieces ...
func (wrapper pieceGroupGetterWrapper) Pieces() []common.Piece {
	return common.Pieces(wrapper)
}

type pieceStorageWithoutMoveChecker interface {
	common.BasePieceStorage
	common.PieceGroupGetter
}

type moveCheckerWrapper struct {
	pieceStorageWithoutMoveChecker
}

// CheckMove ...
//
// It doesn't check for a check before or after the move.
func (wrapper moveCheckerWrapper) CheckMove(move common.Move) error {
	return common.CheckMove(wrapper, move)
}

type pieceStorageWrapper struct {
	common.BasePieceStorage
}

// Pieces ...
func (wrapper pieceStorageWrapper) Pieces() []common.Piece {
	return common.Pieces(wrapper)
}

// CheckMove ...
//
// It doesn't check for a check before or after the move.
func (wrapper pieceStorageWrapper) CheckMove(move common.Move) error {
	return common.CheckMove(wrapper, move)
}

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
