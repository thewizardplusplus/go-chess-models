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
		positionMoves := generator.
			MovesForPosition(storage, position)
		moves = append(moves, positionMoves...)
	}

	return moves
}

// MovesForPosition ...
func (
	generator MoveGenerator,
) MovesForPosition(
	storage PieceStorage,
	position Position,
) []Move {
	var moves []Move
	width := storage.Size().Width
	height := storage.Size().Height
	for rank := 0; rank < height; rank++ {
		for file := 0; file < width; file++ {
			finish := Position{file, rank}
			move := Move{position, finish}
			err := storage.CheckMove(move)
			if err != nil {
				continue
			}

			moves = append(moves, move)
		}
	}

	return moves
}
