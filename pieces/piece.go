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
		models.King: func(
			color models.Color,
			position models.Position,
		) models.Piece {
			return NewKing(color, position)
		},
		models.Queen: func(
			color models.Color,
			position models.Position,
		) models.Piece {
			return NewQueen(color, position)
		},
		models.Rook: func(
			color models.Color,
			position models.Position,
		) models.Piece {
			return NewRook(color, position)
		},
		models.Bishop: func(
			color models.Color,
			position models.Position,
		) models.Piece {
			return NewBishop(color, position)
		},
		models.Knight: func(
			color models.Color,
			position models.Position,
		) models.Piece {
			return NewKnight(color, position)
		},
		models.Pawn: func(
			color models.Color,
			position models.Position,
		) models.Piece {
			return NewPawn(color, position)
		},
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
