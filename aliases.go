package chessmodels

import (
	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
)

// Color ...
//
// Deprecated: use [common.Color] instead.
type Color = common.Color

// ...
const (
	// Deprecated: use common.Black instead.
	Black Color = common.Black
	// Deprecated: use common.White instead.
	White Color = common.White
)

// Kind ...
//
// Deprecated: use [common.Kind] instead.
type Kind = common.Kind

// ...
const (
	// Deprecated: use common.King instead.
	King Kind = common.King
	// Deprecated: use common.Queen instead.
	Queen Kind = common.Queen
	// Deprecated: use common.Rook instead.
	Rook Kind = common.Rook
	// Deprecated: use common.Bishop instead.
	Bishop Kind = common.Bishop
	// Deprecated: use common.Knight instead.
	Knight Kind = common.Knight
	// Deprecated: use common.Pawn instead.
	Pawn Kind = common.Pawn
)

// Position ...
//
// Deprecated: use [common.Position] instead.
type Position = common.Position

// Move ...
//
// Deprecated: use [common.Move] instead.
type Move = common.Move

// Size ...
//
// Deprecated: use [common.Size] instead.
type Size = common.Size

// Piece ...
//
// Deprecated: use [common.Piece] instead.
type Piece = common.Piece

// PieceStorage ...
//
// Deprecated: use [common.PieceStorage] instead.
type PieceStorage = common.PieceStorage

// ...
var (
	// Deprecated: use common.ErrKingCapture instead.
	ErrKingCapture = common.ErrKingCapture
)

// NewBoard ...
//
// Deprecated: use [boards.NewSliceBoard] instead.
func NewBoard(size common.Size, pieces []common.Piece) common.PieceStorage {
	return boards.NewSliceBoard(size, pieces)
}
