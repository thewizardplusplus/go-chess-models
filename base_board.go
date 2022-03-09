package chessmodels

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

// BaseBoard ...
type BaseBoard struct {
	size common.Size
}

// NewBaseBoard ...
func NewBaseBoard(size common.Size) BaseBoard {
	return BaseBoard{size}
}

// common.Size ...
func (board BaseBoard) Size() common.Size {
	return board.size
}
