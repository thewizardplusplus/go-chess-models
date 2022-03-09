package pieces

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

// NewPiece ...
func NewPiece(
	kind common.Kind,
	color common.Color,
	position common.Position,
) common.Piece {
	var piece common.Piece
	switch kind {
	case common.King:
		piece = NewKing(color, position)
	case common.Queen:
		piece = NewQueen(color, position)
	case common.Rook:
		piece = NewRook(color, position)
	case common.Bishop:
		piece = NewBishop(color, position)
	case common.Knight:
		piece = NewKnight(color, position)
	case common.Pawn:
		piece = NewPawn(color, position)
	}

	return piece
}
