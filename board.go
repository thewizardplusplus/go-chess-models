package chessmodels

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
/*func (board Board) CheckMove(
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
  if ok && piece.Color() == target.Color() {
    return ErrFriendlyTarget
  }

  return piece.CheckMove(board, move)
}*/
