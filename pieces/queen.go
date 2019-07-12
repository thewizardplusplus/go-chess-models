package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

// Queen ...
type Queen struct{ base }

// NewQueen ...
func NewQueen(
	color models.Color,
	position models.Position,
) models.Piece {
	kind := models.Queen
	base := base{kind, color, position}
	return Queen{base}
}

// ApplyPosition ...
func (piece Queen) ApplyPosition(
	position models.Position,
) models.Piece {
	base := piece.base.ApplyPosition(position)
	return Queen{base}
}

// CheckMove ...
func (piece Queen) CheckMove(
	move models.Move,
	storage models.PieceStorage,
) bool {
	okForRook := Rook{piece.base}.
		CheckMove(move, storage)
	okForBishop := Bishop{piece.base}.
		CheckMove(move, storage)
	return okForRook || okForBishop
}
