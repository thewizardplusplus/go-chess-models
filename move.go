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

// DefaultMoveGenerator ...
type DefaultMoveGenerator struct {
	Board          Board
	MoveChecker    MoveChecker
	IsCheckAllowed bool
}

// MovesForColor ...
func (
	generator DefaultMoveGenerator,
) MovesForColor(color Color) []Move {
	var moves []Move
	pieces := generator.Board.pieces
	for _, piece := range pieces {
		if piece.Color() != color {
			continue
		}

		position := piece.Position()
		positionMoves :=
			generator.MovesForPosition(position)
		moves = append(moves, positionMoves...)
	}

	return moves
}

// MovesForPosition ...
func (
	generator DefaultMoveGenerator,
) MovesForPosition(start Position) []Move {
	var moves []Move
	width := generator.Board.size.Width
	height := generator.Board.size.Height
	for rank := 0; rank < height; rank++ {
		for file := 0; file < width; file++ {
			finish := Position{file, rank}
			move := Move{start, finish}
			err := generator.MoveChecker.CheckMove(
				move,
				generator.IsCheckAllowed,
			)
			if err != nil {
				continue
			}

			moves = append(moves, move)
		}
	}

	return moves
}
