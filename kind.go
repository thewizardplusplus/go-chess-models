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

type kindGroup map[rune]Kind

func (kinds kindGroup) Add(
	fen rune,
	fenCase int,
	kind Kind,
) {
	casedFEN := unicode.To(fenCase, fen)
	kinds[casedFEN] = kind
}

var (
	kinds = kindGroup{
		'k': King,
		'q': Queen,
		'r': Rook,
		'b': Bishop,
		'n': Knight,
		'p': Pawn,
	}
	kindsInFEN = map[Kind]rune{}

	errUnknownKind = errors.New(
		"unknown kind",
	)
)

func init() {
	completedKinds := make(kindGroup)
	lower, upper :=
		unicode.LowerCase, unicode.UpperCase
	for fen, kind := range kinds {
		completedKinds.Add(fen, lower, kind)
		completedKinds.Add(fen, upper, kind)

		// force a lower case to be independent
		// of a definition of the kinds variable
		kindsInFEN[kind] = unicode.ToLower(fen)
	}

	kinds = completedKinds
}

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
