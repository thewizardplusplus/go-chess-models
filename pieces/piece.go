package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

type factory func(
	color models.Color,
	position models.Position,
) models.Piece

var (
	factories = map[models.Kind]factory{
		models.King:   NewKing,
		models.Queen:  NewQueen,
		models.Rook:   NewRook,
		models.Bishop: NewBishop,
		models.Knight: NewKnight,
		models.Pawn:   NewPawn,
	}
)

// NewPiece ...
func NewPiece(
	kind models.Kind,
	color models.Color,
	position models.Position,
) models.Piece {
	factory := factories[kind]
	return factory(color, position)
}
