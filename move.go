package chessmodels

// Move ...
type Move struct {
	Start  Position
	Finish Position
}

// MoveChecker ...
type MoveChecker interface {
	CheckMove(
		move Move,
		allowedCheck bool,
	) error
}

// MoveGenerator ...
type MoveGenerator struct {
	Board       Board
	MoveChecker MoveChecker
}

// LegalMovesForColor ...
func (
	generator MoveGenerator,
) LegalMovesForColor(
	color Color,
	allowedCheck bool,
) []Move {
	var moves []Move
	pieces := generator.Board.pieces
	for _, piece := range pieces {
		if piece.Color() != color {
			continue
		}

		position := piece.Position()
		positionMoves :=
			generator.LegalMovesForPosition(
				position,
				allowedCheck,
			)
		moves = append(moves, positionMoves...)
	}

	return moves
}

// LegalMovesForPosition ...
func (
	generator MoveGenerator,
) LegalMovesForPosition(
	start Position,
	allowedCheck bool,
) []Move {
	var moves []Move
	width := generator.Board.size.Width
	height := generator.Board.size.Height
	for rank := 0; rank < height; rank++ {
		for file := 0; file < width; file++ {
			finish := Position{file, rank}
			move := Move{start, finish}
			err := generator.MoveChecker.CheckMove(
				move,
				allowedCheck,
			)
			if err != nil {
				continue
			}

			moves = append(moves, move)
		}
	}

	return moves
}
