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
	ErrCheck = errors.New("check")
)

// Size ...
type Size struct {
	Width  int
	Height int
}

// Board ...
type Board struct {
	size   Size
	pieces PieceGroup
}

// NewBoard ...
func NewBoard(
	size Size,
	pieces PieceGroup,
) Board {
	return Board{size, pieces}
}

// ApplyMove ...
// It doesn't check that the move is correct.
func (board Board) ApplyMove(
	move Move,
) Board {
	pieces := board.pieces.Copy()
	pieces.Move(move)

	return NewBoard(board.size, pieces)
}

// CheckMove ...
// It doesn't check that move positions is inside the board.
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

	checkColor, ok := board.CheckColor()
	if ok && checkColor == piece.Color() {
		return ErrCheck
	}

	return nil
}

// CheckColor checks for a check
// and returns a color of a side under it.
//
// If there isn't a check on the board
// the false will be returned
// as a second result.
func (board Board) CheckColor() (
	checkColor Color,
	ok bool,
) {
	return Black, false
}
