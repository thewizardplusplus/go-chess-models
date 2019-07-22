package chessmodels

import (
	"errors"
	"unicode"
)

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

// ParsePiece ...
func ParsePiece(
	kindInFEN rune,
) (Kind, Color, error) {
	var kind Kind
	switch unicode.ToLower(kindInFEN) {
	case 'k':
		kind = King
	case 'q':
		kind = Queen
	case 'r':
		kind = Rook
	case 'b':
		kind = Bishop
	case 'n':
		kind = Knight
	case 'p':
		kind = Pawn
	default:
		return 0, 0, errors.New("unknown kind")
	}

	var color Color
	if unicode.IsLower(kindInFEN) {
		color = Black
	} else {
		color = White
	}

	return kind, color, nil
}
