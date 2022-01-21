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

// PerftHandler ...
type PerftHandler func(move Move, count int, deep int)

// Perft ...
func Perft(
	storage PieceStorage,
	color Color,
	deep int,
	handler PerftHandler,
) int {
	// check for a check should be first, including before a termination check,
	// because a terminated evaluation doesn't make sense for a check position
	var generator MoveGenerator
	moves, err := generator.MovesForColor(storage, color)
	if err != nil {
		return 0
	}

	if deep == 0 {
		return 1
	}

	var count int
	for _, move := range moves {
		nextStorage := storage.ApplyMove(move)
		nextColor := color.Negative()
		moveCount := Perft(
			nextStorage,
			nextColor,
			deep-1,
			handler,
		)
		if handler != nil {
			handler(move, moveCount, deep)
		}

		count += moveCount
	}

	return count
}
