package chessmodels

// MoveGenerator ...
type MoveGenerator struct{}

// MovesForColor ...
//
// It doesn't guarantee an order
// of returned moves.
//
// It doesn't take into account
// possible checks and can generate
// such moves.
//
// It returns an error
// only on a king capture.
func (
	generator MoveGenerator,
) MovesForColor(
	storage PieceStorage,
	color Color,
) ([]Move, error) {
	var moves []Move
	for _, piece := range storage.Pieces() {
		if piece.Color() != color {
			continue
		}

		position := piece.Position()
		positionMoves, err := generator.
			MovesForPosition(storage, position)
		if err != nil {
			return nil, err
		}

		moves = append(moves, positionMoves...)
	}

	return moves, nil
}

// MovesForPosition ...
//
// It doesn't take into account
// possible checks and can generate
// such moves.
//
// It returns an error
// only on a king capture.
func (
	generator MoveGenerator,
) MovesForPosition(
	storage PieceStorage,
	position Position,
) ([]Move, error) {
	var moves []Move
	positions := storage.Size().Positions()
	for _, finish := range positions {
		move := Move{position, finish}
		err := storage.CheckMove(move)
		if err != nil {
			// if the move captures a king,
			// break a generating
			if err == ErrKingCapture {
				return nil, err
			}

			// on other errors
			// just skip this move
			continue
		}

		moves = append(moves, move)
	}

	return moves, nil
}