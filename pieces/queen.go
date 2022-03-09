package pieces

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

// Queen ...
type Queen struct{ Base }

// NewQueen ...
func NewQueen(color common.Color, position common.Position) Queen {
	base := NewBase(common.Queen, color, position)
	return Queen{base}
}

// ApplyPosition ...
func (piece Queen) ApplyPosition(position common.Position) common.Piece {
	base := piece.Base.ApplyPosition(position)
	return Queen{base}
}

// CheckMove ...
func (piece Queen) CheckMove(
	move common.Move,
	storage common.PieceStorage,
) bool {
	okForRook := Rook(piece).CheckMove(move, storage)
	okForBishop := Bishop(piece).CheckMove(move, storage)
	return okForRook || okForBishop
}
