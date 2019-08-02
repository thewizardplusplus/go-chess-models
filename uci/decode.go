package uci

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	models "github.com/thewizardplusplus/go-chess-models"
)

// PieceFactory ...
type PieceFactory func(
	kind models.Kind,
	color models.Color,
	position models.Position,
) models.Piece

// PieceStorageFactory ...
type PieceStorageFactory func(
	size models.Size,
	pieces []models.Piece,
) models.PieceStorage

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

// DecodePieceStorage ...
func DecodePieceStorage(
	fen string,
	pieceFactory PieceFactory,
	pieceStorageFactory PieceStorageFactory,
) (models.PieceStorage, error) {
	ranks := strings.Split(fen, "/")
	reverse(ranks)

	var pieces []models.Piece
	var width int
	for index, rank := range ranks {
		rankPieces, rankWidth, err :=
			decodeRank(index, rank, pieceFactory)
		if err != nil {
			return nil, err
		}

		pieces = append(pieces, rankPieces...)
		if width < rankWidth {
			width = rankWidth
		}
	}

	size := models.Size{width, len(ranks)}
	storage :=
		pieceStorageFactory(size, pieces)
	return storage, nil
}

func decodeRank(
	index int,
	fen string,
	pieceFactory PieceFactory,
) (
	pieces []models.Piece,
	maxFile int,
	err error,
) {
	for _, symbol := range fen {
		piece, err :=
			DecodePiece(symbol, pieceFactory)
		if err != nil {
			shift, err :=
				strconv.Atoi(string(symbol))
			if err != nil {
				return nil, 0, err
			}

			maxFile += shift
			continue
		}

		position :=
			models.Position{maxFile, index}
		piece = piece.ApplyPosition(position)
		pieces = append(pieces, piece)

		maxFile++
	}

	return pieces, maxFile, nil
}
