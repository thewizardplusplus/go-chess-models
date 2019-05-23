package chessmodels

// Kind ...
type Kind int

// ...
const (
	King Kind = iota
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

// Color ...
type Color int

// ...
const (
	Black Color = iota
	White
)

// Negative ...
func (color Color) Negative() Color {
	if color == Black {
		return White
	}

	return Black
}

// Position ...
type Position struct {
	File int // column
	Rank int // row
}

// Move ...
type Move struct {
	Start  Position
	Finish Position
}

// Piece ...
type Piece interface {
	Kind() Kind
	Color() Color
	Position() Position
	ApplyPosition(position Position) Piece
	CheckMove(move Move, board Board) bool
}

// PieceGroup ...
type PieceGroup map[Position]Piece

// NewPieceGroup ...
func NewPieceGroup(
	pieces []Piece,
) PieceGroup {
	group := make(PieceGroup)
	for _, piece := range pieces {
		group.Add(piece)
	}

	return group
}

// Add ...
func (group PieceGroup) Add(piece Piece) {
	group[piece.Position()] = piece
}

// Move ...
// It doesn't check that the move is correct.
func (group PieceGroup) Move(move Move) {
	piece := group[move.Start].
		ApplyPosition(move.Finish)
	group.Add(piece)
	delete(group, move.Start)
}

// Copy ...
func (group PieceGroup) Copy() PieceGroup {
	groupCopy := make(PieceGroup)
	for position, piece := range group {
		groupCopy[position] = piece
	}

	return groupCopy
}
