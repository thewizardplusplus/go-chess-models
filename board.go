package chessmodels

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

// Board ...
type Board = SliceBoard

// NewBoard ...
func NewBoard(size common.Size, pieces []common.Piece) common.PieceStorage {
	return NewSliceBoard(size, pieces)
}
