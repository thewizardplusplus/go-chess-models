package chessmodels

import (
	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
)

// Board ...
type Board = boards.SliceBoard

// NewBoard ...
func NewBoard(size common.Size, pieces []common.Piece) common.PieceStorage {
	return boards.NewSliceBoard(size, pieces)
}
