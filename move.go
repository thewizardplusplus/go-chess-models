package chessmodels

// Move ...
type Move struct {
	Start  Position
	Finish Position
}

// MoveGenerator ...
type MoveGenerator struct{}

// MovesForColor ...
//
// It doesn't guarantee an order
// of returned moves.
func (
	generator MoveGenerator,
) MovesForColor(
	storage PieceStorage,
	color Color,
) []Move {
	var moves []Move
	for _, piece := range storage.Pieces() {
		if piece.Color() != color {
			continue
		}

		position := piece.Position()
		positionMoves, _ := generator.
			MovesForPosition(storage, position)
		moves = append(moves, positionMoves...)
	}

	return moves
}

// MovesForPosition ...
//
// It doesn't take into account
// possible checks and can generate
// such moves.
//
// It returns an error only on a king capture.
func (
	generator MoveGenerator,
) MovesForPosition(
	storage PieceStorage,
	position Position,
) ([]Move, error) {
	var moves []Move
	width := storage.Size().Width
	height := storage.Size().Height
	for rank := 0; rank < height; rank++ {
		for file := 0; file < width; file++ {
			finish := Position{file, rank}
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
	}

	return moves, nil
}
