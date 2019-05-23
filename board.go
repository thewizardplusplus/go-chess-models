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

// Positions ...
func (size Size) Positions() []Position {
	var positions []Position
	width, height := size.Width, size.Height
	for rank := 0; rank < height; rank++ {
		for file := 0; file < width; file++ {
			position := Position{file, rank}
			positions = append(
				positions,
				position,
			)
		}
	}

	return positions
}

// Board ...
type Board struct {
	size      Size
	positions []Position
	pieces    PieceGroup
}

// NewBoard ...
func NewBoard(
	size Size,
	pieces PieceGroup,
) Board {
	positions := size.Positions()
	return Board{size, positions, pieces}
}

// ApplyMove ...
// It doesn't check that the move is correct.
func (board Board) ApplyMove(
	move Move,
) Board {
	pieces := board.pieces.Copy()
	pieces.Move(move)

	return Board{
		size:      board.size,
		positions: board.positions,
		pieces:    pieces,
	}
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

	if board.IsCheckForColor(piece.Color()) {
		return ErrCheck
	}

	return nil
}

// IsCheckForColor ...
func (board Board) IsCheckForColor(
	color Color,
) bool {
	moves := board.
		LegalMovesForColor(color.Negative())
	for _, move := range moves {
		piece, ok := board.pieces[move.Finish]
		if ok && piece.Kind() == King {
			return true
		}
	}

	return false
}

// LegalMoves ...
func (board Board) LegalMovesForColor(
	color Color,
) []Move {
	var moves []Move
	for _, piece := range board.pieces {
		if piece.Color() != color {
			continue
		}

		piecePosition := piece.Position()
		pieceMoves := board.
			LegalMovesForPosition(piecePosition)
		moves = append(moves, pieceMoves...)
	}

	return moves
}

// LegalMovesForPosition ...
func (board Board) LegalMovesForPosition(
	start Position,
) []Move {
	var moves []Move
	for _, finish := range board.positions {
		move := Move{start, finish}
		err := board.CheckMove(move)
		if err == nil {
			moves = append(moves, move)
		}
	}

	return moves
}
