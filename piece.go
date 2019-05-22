package chessmodels

// Position ...
type Position struct {
	File int // column
	Rank int // row
}

// Piece ...
type Piece interface {
	Position() Position
}

// PieceGroup ...
type PieceGroup map[Position]Piece

// Add ...
func (group PieceGroup) Add(piece Piece) {
	group[piece.Position()] = piece
}
