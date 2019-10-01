package chessmodels

import (
	"errors"
)

// ...
var (
	ErrNoMove         = errors.New("no move")
	ErrNoPiece        = errors.New("no piece")
	ErrFriendlyTarget = errors.New(
		"friendly target",
	)
	ErrIllegalMove = errors.New(
		"illegal move",
	)
	ErrKingCapture = errors.New(
		"king capture",
	)
)

// PieceStorage ...
type PieceStorage interface {
	Size() Size
	Piece(
		position Position,
	) (piece Piece, ok bool)
	Pieces() []Piece

	// It shouldn't check that the move
	// is correct.
	ApplyMove(move Move) PieceStorage

	// It shouldn't check that move positions
	// is inside the board.
	//
	// It shouldn't check for a check
	// before or after the move.
	CheckMove(move Move) error
}

// Board ...
type Board struct {
	size   Size
	pieces pieceGroup
}

// NewBoard ...
func NewBoard(
	size Size,
	pieces []Piece,
) PieceStorage {
	pieceGroup := newPieceGroup(pieces)
	return Board{size, pieceGroup}
}

// Size ...
func (board Board) Size() Size {
	return board.size
}

// Piece ...
func (board Board) Piece(
	position Position,
) (piece Piece, ok bool) {
	piece, ok = board.pieces[position]
	return piece, ok
}

// Pieces ...
//
// It doesn't guarantee an order
// of returned pieces.
func (board Board) Pieces() []Piece {
	var pieces []Piece
	for _, piece := range board.pieces {
		pieces = append(pieces, piece)
	}

	return pieces
}

// ApplyMove ...
//
// It doesn't check that the move
// is correct.
func (board Board) ApplyMove(
	move Move,
) PieceStorage {
	pieces := board.pieces.Copy()
	pieces.Move(move)

	return Board{board.size, pieces}
}

// CheckMove ...
//
// It doesn't check that move positions
// is inside the board.
//
// It doesn't check for a check
// before or after the move.
func (board Board) CheckMove(
	move Move,
) error {
	if move.IsEmpty() {
		return ErrNoMove
	}

	piece, ok := board.pieces[move.Start]
	if !ok {
		return ErrNoPiece
	}

	target, hasTarget :=
		board.pieces[move.Finish]
	if hasTarget &&
		target.Color() == piece.Color() {
		return ErrFriendlyTarget
	}

	if !piece.CheckMove(move, board) {
		return ErrIllegalMove
	}

	// this check should be occurred
	// only for legal moves
	// (i.e. after all rest checks)
	//
	// it's necessary because of a wider area
	// of ​​influence of this error
	// compared to rest ones
	if hasTarget && target.Kind() == King {
		return ErrKingCapture
	}

	return nil
}
