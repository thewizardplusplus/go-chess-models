package chessmodels

import (
	"unicode"
)

// Color ...
type Color int

// ...
const (
	Black Color = iota
	White
)

// ParseColor ...
//
// It parses a color of a piece
// from its kind in FEN.
//
// It doesn't check that the kind is valid,
// and relies only on its case.
func ParseColor(kindInFEN byte) Color {
	if unicode.IsUpper(rune(kindInFEN)) {
		return White
	}

	return Black
}

// Negative ...
func (color Color) Negative() Color {
	if color == Black {
		return White
	}

	return Black
}
