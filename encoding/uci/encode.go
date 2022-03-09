package uci

import (
	"strconv"
	"strings"
	"unicode"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/common"
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
	case common.Black:
		kindCase = unicode.LowerCase
	case common.White:
		kindCase = unicode.UpperCase
	}

	var kindInFEN rune
	switch piece.Kind() {
	case common.King:
		kindInFEN = 'k'
	case common.Queen:
		kindInFEN = 'q'
	case common.Rook:
		kindInFEN = 'r'
	case common.Bishop:
		kindInFEN = 'b'
	case common.Knight:
		kindInFEN = 'n'
	case common.Pawn:
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

	ranks := make([]string, 0, storage.Size().Height)
	storage.Size().IteratePositions(func(position models.Position) error { // nolint: errcheck, gosec, lll
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

		return nil
	})

	reverse(ranks)
	return strings.Join(ranks, "/")
}
