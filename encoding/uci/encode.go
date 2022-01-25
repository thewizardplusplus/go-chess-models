package uci

import (
	"strconv"
	"strings"
	"unicode"

	models "github.com/thewizardplusplus/go-chess-models"
)

// EncodePosition ...
//
// It converts the position to pure algebraic coordinate notation.
func EncodePosition(position models.Position) string {
	file := string(rune(position.File + minFileSymbol))
	rank := strconv.Itoa(position.Rank + 1)
	return file + rank
}

// EncodeMove ...
//
// It converts the move to pure algebraic coordinate notation.
func EncodeMove(move models.Move) string {
	start := EncodePosition(move.Start)
	finish := EncodePosition(move.Finish)
	return start + finish
}

// EncodePiece ...
//
// It converts the piece to FEN (only a kind and a color, not a position).
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
func EncodePieceStorage(storage models.PieceStorage) string {
	var rank string
	var shift int
	resetShift := func() {
		if shift != 0 {
			rank += strconv.Itoa(shift)
			shift = 0
		}
	}

	var ranks []string
	for _, position := range storage.Size().Positions() {
		if piece, ok := storage.Piece(position); ok {
			resetShift()

			rank += EncodePiece(piece)
		} else {
			shift++
		}

		// last file
		if position.File == storage.Size().Width-1 {
			resetShift()

			ranks = append(ranks, rank)
			rank = ""
		}
	}

	reverse(ranks)
	return strings.Join(ranks, "/")
}
