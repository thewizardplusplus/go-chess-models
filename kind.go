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

type kindGroup map[byte]Kind

func (kinds kindGroup) Add(
	fen byte,
	fenCase int,
	kind Kind,
) {
	casedFEN :=
		byte(unicode.To(fenCase, rune(fen)))
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
)

func init() {
	completedKinds := make(kindGroup)
	lower, upper :=
		unicode.LowerCase, unicode.UpperCase
	for fen, kind := range kinds {
		completedKinds.Add(fen, lower, kind)
		completedKinds.Add(fen, upper, kind)
	}

	kinds = completedKinds
}

// ParseKind ...
//
// It parses a kind of a piece
// from one in FEN.
func ParseKind(
	kindInFEN byte,
) (Kind, error) {
	kind, ok := kinds[kindInFEN]
	if !ok {
		return 0, errors.New("unknown kind")
	}

	return kind, nil
}
