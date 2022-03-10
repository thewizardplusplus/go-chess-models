package chessmodels

import (
	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
)

// BaseBoard ...
type BaseBoard = boards.BaseBoard

// DefaultBoardWrapper ...
type DefaultBoardWrapper = boards.DefaultBoardWrapper

// MapBoard ...
type MapBoard = boards.MapBoard

// NewMapBoard ...
func NewMapBoard(size common.Size, pieces []common.Piece) common.PieceStorage {
	return boards.NewMapBoard(size, pieces)
}

// SliceBoard ...
type SliceBoard = boards.SliceBoard

// NewSliceBoard ...
func NewSliceBoard(size common.Size, pieces []common.Piece) common.PieceStorage {
	return boards.NewSliceBoard(size, pieces)
}
