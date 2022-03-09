package chessmodels

import (
	"errors"

	"github.com/thewizardplusplus/go-chess-models/common"
)

// ...
var (
	ErrNoMove         = errors.New("no move")
	ErrOutOfSize      = errors.New("out of size")
	ErrNoPiece        = errors.New("no piece")
	ErrFriendlyTarget = errors.New("friendly target")
	ErrIllegalMove    = errors.New("illegal move")
	ErrKingCapture    = errors.New("king capture")
)

// CheckMove ...
//
// It doesn't check for a check before or after the move.
func CheckMove(storage PieceStorage, move common.Move) error {
	if move.IsEmpty() {
		return ErrNoMove
	}

	if !storage.Size().HasMove(move) {
		return ErrOutOfSize
	}

	piece, ok := storage.Piece(move.Start)
	if !ok {
		return ErrNoPiece
	}

	target, hasTarget := storage.Piece(move.Finish)
	if hasTarget && target.Color() == piece.Color() {
		return ErrFriendlyTarget
	}

	if !piece.CheckMove(move, storage) {
		return ErrIllegalMove
	}

	// this check should be occurred only for legal moves
	// (i.e. after all rest checks)
	//
	// it's necessary because of a wider area of ​​influence of this error
	// compared to rest ones
	if hasTarget && target.Kind() == common.King {
		return ErrKingCapture
	}

	return nil
}
