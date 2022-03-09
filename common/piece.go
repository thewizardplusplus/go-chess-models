package common

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
