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
	pieces []Piece,
) Board {
	return Board{
		size:   size,
		pieces: NewPieceGroup(pieces),
	}
}

// ApplyMove ...
// It doesn't check that the move is correct.
func (board Board) ApplyMove(
	move Move,
) Board {
	pieces := board.pieces.Copy()
	pieces.Move(move)

	return Board{
		size:   board.size,
		pieces: pieces,
	}
}

// CheckMove ...
// It doesn't check that move positions is inside the board.
func (board Board) CheckMove(
	move Move,
	allowedCheck bool,
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

	if !allowedCheck {
		nextBoard := board.ApplyMove(move)
		if nextBoard.IsCheckForColor(
			piece.Color(),
		) {
			return ErrCheck
		}
	}

	return nil
}

// IsCheckForColor ...
func (board Board) IsCheckForColor(
	color Color,
) bool {
	generator := MoveGenerator{board, board}
	moves := generator.LegalMovesForColor(
		color.Negative(),
		true, // allowedCheck
	)
	for _, move := range moves {
		piece, ok := board.pieces[move.Finish]
		if ok && piece.Kind() == King {
			return true
		}
	}

	return false
}
