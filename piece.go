package chessmodels

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

// Piece ...
type Piece interface {
	Kind() common.Kind
	Color() common.Color
	Position() common.Position
	ApplyPosition(position common.Position) Piece

	// It shouldn't check that move positions is inside the board.
	//
	// It shouldn't check that the move finish position isn't equal
	// to its start position.
	//
	// It shouldn't check that the start move position corresponds
	// to the piece position.
	//
	// It shouldn't check that there isn't a friendly piece
	// on the move finish position.
	//
	// It shouldn't check that there isn't an enemy king
	// on the move finish position.
	//
	// It shouldn't check for a check before or after the move.
	CheckMove(move common.Move, storage PieceStorage) bool
}

// BasePieceStorage ...
type BasePieceStorage interface {
	Size() common.Size
	Piece(position common.Position) (piece Piece, ok bool)

	// It shouldn't check that the move is correct.
	ApplyMove(move common.Move) PieceStorage
}

// PieceGroupGetter ...
type PieceGroupGetter interface {
	Pieces() []Piece
}

// MoveChecker ...
type MoveChecker interface {
	// It shouldn't check for a check before or after the move.
	CheckMove(move common.Move) error
}

// PieceStorage ...
type PieceStorage interface {
	BasePieceStorage
	PieceGroupGetter
	MoveChecker
}

// Pieces ...
func Pieces(storage PieceStorage) []Piece {
	var pieces []Piece
	storage.Size().IteratePositions(func(position common.Position) error { // nolint: errcheck, gosec, lll
		if piece, ok := storage.Piece(position); ok {
			pieces = append(pieces, piece)
		}

		return nil
	})

	return pieces
}
