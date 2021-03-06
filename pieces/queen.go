package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

// Queen ...
type Queen struct{ Base }

// NewQueen ...
func NewQueen(color models.Color, position models.Position) Queen {
	base := NewBase(models.Queen, color, position)
	return Queen{base}
}

// ApplyPosition ...
func (piece Queen) ApplyPosition(position models.Position) models.Piece {
	base := piece.Base.ApplyPosition(position)
	return Queen{base}
}

// CheckMove ...
func (piece Queen) CheckMove(
	move models.Move,
	storage models.PieceStorage,
) bool {
	okForRook := Rook(piece).CheckMove(move, storage)
	okForBishop := Bishop(piece).CheckMove(move, storage)
	return okForRook || okForBishop
}
