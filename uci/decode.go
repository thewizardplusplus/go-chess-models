package uci

import (
	"errors"
	"unicode"

	models "github.com/thewizardplusplus/go-chess-models"
)

// PieceFactory ...
type PieceFactory func(
	kind models.Kind,
	color models.Color,
	position models.Position,
) models.Piece

// DecodePiece ...
func DecodePiece(
	fen rune,
	factory PieceFactory,
) (models.Piece, error) {
	var kind models.Kind
	switch unicode.ToLower(fen) {
	case 'k':
		kind = models.King
	case 'q':
		kind = models.Queen
	case 'r':
		kind = models.Rook
	case 'b':
		kind = models.Bishop
	case 'n':
		kind = models.Knight
	case 'p':
		kind = models.Pawn
	default:
		return nil, errors.New("unknown kind")
	}

	var color models.Color
	if unicode.IsLower(fen) {
		color = models.Black
	} else {
		color = models.White
	}

	var position models.Position
	piece := factory(kind, color, position)
	return piece, nil
}
