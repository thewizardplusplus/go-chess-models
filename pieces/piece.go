package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

// NewPiece ...
func NewPiece(
	kind models.Kind,
	color models.Color,
	position models.Position,
) models.Piece {
	var piece models.Piece
	switch kind {
	case models.King:
		piece = NewKing(color, position)
	case models.Queen:
		piece = NewQueen(color, position)
	case models.Rook:
		piece = NewRook(color, position)
	case models.Bishop:
		piece = NewBishop(color, position)
	case models.Knight:
		piece = NewKnight(color, position)
	case models.Pawn:
		piece = NewPawn(color, position)
	}

	return piece
}
