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
	kindMap = map[byte]Kind{
		'k': King,
		'q': Queen,
		'r': Rook,
		'b': Bishop,
		'n': Knight,
		'p': Pawn,
	}
)

func init() {
	completedKindMap := make(map[byte]Kind)
	for fen, kind := range kindMap {
		lowerFEN := unicode.ToLower(rune(fen))
		completedKindMap[byte(lowerFEN)] = kind

		upperFEN := unicode.ToUpper(rune(fen))
		completedKindMap[byte(upperFEN)] = kind
	}

	kindMap = completedKindMap
}

// ParseKind ...
//
// It parses a kind of a piece
// from one in FEN.
func ParseKind(
	kindInFEN byte,
) (Kind, error) {
	kind, ok := kindMap[kindInFEN]
	if !ok {
		return 0, errors.New("unknown kind")
	}

	return kind, nil
}
