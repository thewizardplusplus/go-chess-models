package pieces

import (
	"errors"
	"unicode"

	models "github.com/thewizardplusplus/go-chess-models"
)

// PieceFactory ...
type PieceFactory func(
	kind models.Kind,
	color models.Color,
) models.Piece

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

// ParsePiece ...
func ParsePiece(
	fen rune,
	pieceFactory PieceFactory,
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

	piece := pieceFactory(kind, color)
	return piece, nil
}

// ParseDefaultPiece ...
func ParseDefaultPiece(
	fen rune,
) (models.Piece, error) {
	return ParsePiece(
		fen,
		func(
			kind models.Kind,
			color models.Color,
		) models.Piece {
			return NewPiece(
				kind,
				color,
				models.Position{},
			)
		},
	)
}
