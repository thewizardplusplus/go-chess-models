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
