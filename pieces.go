package chessmodels

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

// Pieces ...
func Pieces(storage common.PieceStorage) []common.Piece {
	var pieces []common.Piece
	storage.Size().IteratePositions(func(position common.Position) error { // nolint: errcheck, gosec, lll
		if piece, ok := storage.Piece(position); ok {
			pieces = append(pieces, piece)
		}

		return nil
	})

	return pieces
}
