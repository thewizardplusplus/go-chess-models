package chessmodels

import (
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

// PositionHandler ...
type PositionHandler = common.PositionHandler

// Size ...
type Size = common.Size

// Piece ...
type Piece = common.Piece

// PieceStorage ...
type PieceStorage = common.PieceStorage

// ...
var (
	ErrNoMove         = common.ErrNoMove
	ErrOutOfSize      = common.ErrOutOfSize
	ErrNoPiece        = common.ErrNoPiece
	ErrFriendlyTarget = common.ErrFriendlyTarget
	ErrIllegalMove    = common.ErrIllegalMove
	ErrKingCapture    = common.ErrKingCapture
)

// CheckMove ...
//
// It doesn't check for a check before or after the move.
func CheckMove(storage PieceStorage, move Move) error {
	return common.CheckMove(storage, move)
}

// Pieces ...
func Pieces(storage PieceStorage) []Piece {
	return common.Pieces(storage)
}
