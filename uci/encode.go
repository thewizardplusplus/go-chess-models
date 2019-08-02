package uci

import (
	"strconv"
	"strings"
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

// EncodePieceStorage ...
//
// It converts the piece storage to FEN.
func EncodePieceStorage(
	storage models.PieceStorage,
) string {
	var rank string
	var shift int
	resetShift := func() {
		if shift != 0 {
			rank += strconv.Itoa(shift)
			shift = 0
		}
	}

	var ranks []string
	positions := storage.Size().Positions()
	for _, position := range positions {
		piece, ok := storage.Piece(position)
		if ok {
			resetShift()

			rank += EncodePiece(piece)
		} else {
			shift++
		}

		lastFile := storage.Size().Height - 1
		if position.File == lastFile {
			resetShift()

			ranks = append(ranks, rank)
			rank = ""
		}
	}

	reverse(ranks)
	return strings.Join(ranks, "/")
}
