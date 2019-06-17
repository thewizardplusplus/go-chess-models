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

// Size ...
type Size struct {
	Width  int
	Height int
}

// PieceStorage ...
type PieceStorage interface {
	Size() Size
	Piece(
		position Position,
	) (piece Piece, ok bool)
	Pieces() []Piece
	ApplyMove(move Move) PieceStorage
	CheckMove(move Move) error
	CheckMoves(moves []Move) error
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
) Board {
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
	if move.Start == move.Finish {
		return ErrNoMove
	}

	piece, ok := board.pieces[move.Start]
	if !ok {
		return ErrNoPiece
	}

	target, ok := board.pieces[move.Finish]
	if ok && target.Color() == piece.Color() {
		return ErrFriendlyTarget
	}

	if !piece.CheckMove(move, board) {
		return ErrIllegalMove
	}

	return nil
}

// CheckMoves ...
//
// It checks only for absence
// of a king capture.
func (board Board) CheckMoves(
	moves []Move,
) error {
	for _, move := range moves {
		piece, ok := board.pieces[move.Finish]
		if ok && piece.Kind() == King {
			return ErrKingCapture
		}
	}

	return nil
}
