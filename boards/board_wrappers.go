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

// WrapBasePieceStorage ...
func WrapBasePieceStorage(
	baseStorage common.BasePieceStorage,
) common.PieceStorage {
	switch partialStorage := baseStorage.(type) {
	case pieceStorageWithoutPieceGroupGetter:
		return pieceGroupGetterWrapper{partialStorage}
	case pieceStorageWithoutMoveChecker:
		return moveCheckerWrapper{partialStorage}
	default:
		return pieceStorageWrapper{baseStorage}
	}
}
