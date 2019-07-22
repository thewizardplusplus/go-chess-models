package pieces

import (
	"unicode"

	models "github.com/thewizardplusplus/go-chess-models"
)

// Base ...
type Base struct {
	kind     models.Kind
	color    models.Color
	position models.Position
}

// Kind ...
func (piece Base) Kind() models.Kind {
	return piece.kind
}

// Color ...
func (piece Base) Color() models.Color {
	return piece.color
}

// Position ...
func (
	piece Base,
) Position() models.Position {
	return piece.position
}

// ApplyPosition ...
func (piece Base) ApplyPosition(
	position models.Position,
) Base {
	kind, color := piece.kind, piece.color
	return Base{kind, color, position}
}

// String ...
//
// It converts the piece to FEN
// (only a kind and a color, not a position).
func (piece Base) String() string {
	var kindCase int
	switch piece.color {
	case models.Black:
		kindCase = unicode.LowerCase
	case models.White:
		kindCase = unicode.UpperCase
	}

	var kindInFEN rune
	switch piece.kind {
	case models.King:
		kindInFEN = 'k'
	case models.Queen:
		kindInFEN = 'q'
	case models.Rook:
		kindInFEN = 'r'
	case models.Bishop:
		kindInFEN = 'b'
	case models.Knight:
		kindInFEN = 'n'
	case models.Pawn:
		kindInFEN = 'p'
	}

	fen := unicode.To(kindCase, kindInFEN)
	return string(fen)
}
