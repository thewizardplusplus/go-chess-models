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
	kindsInFEN = map[Kind]rune{
		King:   'k',
		Queen:  'q',
		Rook:   'r',
		Bishop: 'b',
		Knight: 'n',
		Pawn:   'p',
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
	if color == Black {
		kindCase = unicode.LowerCase
	} else {
		kindCase = unicode.UpperCase
	}

	kindInFEN := kindsInFEN[kind]
	kindInFEN =
		unicode.To(kindCase, kindInFEN)

	return kindInFEN
}
