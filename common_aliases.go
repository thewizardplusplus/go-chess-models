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
