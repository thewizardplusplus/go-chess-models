package chessmodels

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
	Position() Position
	ApplyPosition(position Position) Piece
}

// PieceGroup ...
type PieceGroup map[Position]Piece

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
