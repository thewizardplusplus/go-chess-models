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

	errUnknownKind = errors.New(
		"unknown kind",
	)
)

// ParsePiece ...
func ParsePiece(
	kindInFEN rune,
) (Kind, Color, error) {
	kind, ok := kinds[kindInFEN]
	if !ok {
		return 0, 0, errUnknownKind
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
func (kind Kind) ToFEN(
	color Color,
) (rune, error) {
	kindInFEN, ok := kindsInFEN[kind]
	if !ok {
		return 0, errUnknownKind
	}

	var kindCase int
	if color == Black {
		kindCase = unicode.LowerCase
	} else {
		kindCase = unicode.UpperCase
	}
	kindInFEN =
		unicode.To(kindCase, kindInFEN)

	return kindInFEN, nil
}
