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

// ToFEN ...
func (kind Kind) ToFEN(color Color) rune {
	var kindCase int
	switch color {
	case Black:
		kindCase = unicode.LowerCase
	case White:
		kindCase = unicode.UpperCase
	}

	var kindInFEN rune
	switch kind {
	case King:
		kindInFEN = 'k'
	case Queen:
		kindInFEN = 'q'
	case Rook:
		kindInFEN = 'r'
	case Bishop:
		kindInFEN = 'b'
	case Knight:
		kindInFEN = 'n'
	case Pawn:
		kindInFEN = 'p'
	}

	return unicode.To(kindCase, kindInFEN)
}
