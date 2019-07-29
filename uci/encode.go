package uci

import (
	"unicode"

	models "github.com/thewizardplusplus/go-chess-models"
)

// EncodePiece ...
//
// It converts the piece to FEN
// (only a kind and a color, not a position).
func EncodePiece(piece models.Piece) string {
	var kindCase int
	switch piece.Color() {
	case models.Black:
		kindCase = unicode.LowerCase
	case models.White:
		kindCase = unicode.UpperCase
	}

	var kindInFEN rune
	switch piece.Kind() {
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
