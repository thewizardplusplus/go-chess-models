package chessmodels

// MoveGenerator ...
type MoveGenerator struct{}

// MovesForColor ...
//
// It doesn't guarantee an order of returned moves.
//
// It doesn't take into account possible checks and can generate such moves.
//
// It returns an error only on a king capture.
func (generator MoveGenerator) MovesForColor(
	storage PieceStorage,
	color Color,
) ([]Move, error) {
	var moves []Move
	for _, piece := range storage.Pieces() {
		if piece.Color() != color {
			continue
		}

		positionMoves, err :=
			generator.MovesForPosition(storage, piece.Position())
		if err != nil {
			return nil, err
		}

		moves = append(moves, positionMoves...)
	}

	return moves, nil
}

// MovesForPosition ...
//
// It doesn't take into account possible checks and can generate such moves.
//
// It returns an error only on a king capture.
func (generator MoveGenerator) MovesForPosition(
	storage PieceStorage,
	position Position,
) ([]Move, error) {
	var moves []Move
	for _, finish := range storage.Size().Positions() {
		move := Move{position, finish}
		if err := storage.CheckMove(move); err != nil {
			// if the move captures a king, break a generating
			if err == ErrKingCapture {
				return nil, err
			}

			// on other errors just skip this move
			continue
		}

		moves = append(moves, move)
	}

	return moves, nil
}
