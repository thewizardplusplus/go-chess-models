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

// Position ...
type Position struct {
	File int // column
	Rank int // row
}

// Piece ...
type Piece interface {
	Kind() Kind
	Color() Color
	Position() Position
	ApplyPosition(position Position) Piece

	// It shouldn't check
	// that move positions is inside the board.
	//
	// It shouldn't check
	// that the move finish position
	// isn't equal to its start position.
	//
	// It shouldn't check
	// that the start move position
	// corresponds to the piece position.
	//
	// It shouldn't check
	// that there isn't a friendly piece
	// on the move finish position.
	//
	// It shouldn't check
	// for a check before or after the move.
	CheckMove(move Move, board Board) bool
}

type pieceGroup map[Position]Piece

func newPieceGroup(
	pieces []Piece,
) pieceGroup {
	group := make(pieceGroup)
	for _, piece := range pieces {
		group.Add(piece)
	}

	return group
}

func (group pieceGroup) Add(piece Piece) {
	group[piece.Position()] = piece
}

// It doesn't check that the move is correct.
func (group pieceGroup) Move(move Move) {
	piece := group[move.Start]
	delete(group, move.Start)

	movedPiece := piece.
		ApplyPosition(move.Finish)
	group.Add(movedPiece)
}

func (group pieceGroup) Copy() pieceGroup {
	groupCopy := make(pieceGroup)
	for _, piece := range group {
		groupCopy.Add(piece)
	}

	return groupCopy
}
