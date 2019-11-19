package uci

import (
	"errors"
	"fmt"
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

const (
	minFileCount = 97
)

// DecodePosition ...
//
// It decodes a position from pure algebraic coordinate notation.
func DecodePosition(text string) (position models.Position, err error) {
	if len(text) != 2 {
		return models.Position{}, errors.New("incorrect length")
	}

	file := int(text[0]) - minFileCount
	if file < 0 {
		return models.Position{}, errors.New("incorrect file")
	}

	rank, err := strconv.Atoi(text[1:])
	if err != nil {
		return models.Position{}, fmt.Errorf("incorrect rank: %s", err)
	}
	rank--

	return models.Position{File: file, Rank: rank}, nil
}

// DecodeMove ...
//
// It decodes a move from pure algebraic coordinate notation.
func DecodeMove(text string) (move models.Move, err error) {
	if len(text) != 4 {
		return models.Move{}, errors.New("incorrect length")
	}

	start, err := DecodePosition(text[:2])
	if err != nil {
		return models.Move{}, fmt.Errorf("incorrect start: %s", err)
	}

	finish, err := DecodePosition(text[2:])
	if err != nil {
		return models.Move{}, fmt.Errorf("incorrect finish: %s", err)
	}

	return models.Move{Start: start, Finish: finish}, nil
}

// DecodePiece ...
//
// It decodes a piece from FEN (only a kind and a color, not a position).
func DecodePiece(fen rune, factory PieceFactory) (models.Piece, error) {
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

	piece := factory(kind, color, models.Position{})
	return piece, nil
}

// DecodePieceStorage ...
//
// It decodes a piece storage from FEN.
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
		rankPieces, rankWidth, err := decodeRank(index, rank, pieceFactory)
		if err != nil {
			return nil, err
		}

		pieces = append(pieces, rankPieces...)
		if width < rankWidth {
			width = rankWidth
		}
	}

	size := models.Size{Width: width, Height: len(ranks)}
	storage := pieceStorageFactory(size, pieces)
	return storage, nil
}

func decodeRank(
	index int,
	fen string,
	pieceFactory PieceFactory,
) (pieces []models.Piece, maxFile int, err error) {
	for _, symbol := range fen {
		piece, err := DecodePiece(symbol, pieceFactory)
		if err != nil {
			shift, err := strconv.Atoi(string(symbol))
			if err != nil {
				return nil, 0, err
			}

			maxFile += shift
			continue
		}

		placedPiece :=
			piece.ApplyPosition(models.Position{File: maxFile, Rank: index})
		pieces = append(pieces, placedPiece)

		maxFile++
	}

	return pieces, maxFile, nil
}
