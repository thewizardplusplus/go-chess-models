package chessmodels

import (
	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
)

// Color ...
type Color = common.Color

// ...
const (
	Black = common.Black
	White = common.White
)

// Kind ...
type Kind = common.Kind

// ...
const (
	King   = common.King
	Queen  = common.Queen
	Rook   = common.Rook
	Bishop = common.Bishop
	Knight = common.Knight
	Pawn   = common.Pawn
)

// Position ...
type Position = common.Position

// Move ...
type Move = common.Move

// Size ...
type Size = common.Size

// Piece ...
type Piece = common.Piece

// PieceStorage ...
type PieceStorage = common.PieceStorage

// ...
var (
	ErrKingCapture = common.ErrKingCapture
)

// NewBoard ...
func NewBoard(size common.Size, pieces []common.Piece) common.PieceStorage {
	return boards.NewSliceBoard(size, pieces)
}
