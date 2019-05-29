package chessmodels

// Move ...
type Move struct {
	Start  Position
	Finish Position
}

// PieceStorage ...
type PieceStorage interface {
	Size() Size
	Pieces() []Piece
	CheckMove(move Move) error
}

// MoveGenerator ...
type MoveGenerator struct {
	PieceStorage PieceStorage
}

// MovesForColor ...
//
// It doesn't guarantee an order
// of returned moves.
func (
	generator MoveGenerator,
) MovesForColor(color Color) []Move {
	var moves []Move
	pieces := generator.PieceStorage.Pieces()
	for _, piece := range pieces {
		if piece.Color() != color {
			continue
		}

		position := piece.Position()
		positionMoves := generator.
			MovesForPosition(position)
		moves = append(moves, positionMoves...)
	}

	return moves
}

// MovesForPosition ...
func (
	generator MoveGenerator,
) MovesForPosition(start Position) []Move {
	var moves []Move
	size := generator.PieceStorage.Size()
	width, height := size.Width, size.Height
	for rank := 0; rank < height; rank++ {
		for file := 0; file < width; file++ {
			finish := Position{file, rank}
			move := Move{start, finish}
			err := generator.PieceStorage.
				CheckMove(move)
			if err != nil {
				continue
			}

			moves = append(moves, move)
		}
	}

	return moves
}
