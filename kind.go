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

var (
	kinds = map[rune]Kind{
		// black
		'k': King,
		'q': Queen,
		'r': Rook,
		'b': Bishop,
		'n': Knight,
		'p': Pawn,

		// white
		'K': King,
		'Q': Queen,
		'R': Rook,
		'B': Bishop,
		'N': Knight,
		'P': Pawn,
	}
)

// ParsePiece ...
func ParsePiece(
	kindInFEN rune,
) (Kind, Color, error) {
	kind, ok := kinds[kindInFEN]
	if !ok {
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
