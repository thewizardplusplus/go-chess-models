package uci

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/thewizardplusplus/go-chess-models/common"
)

// PieceFactory ...
type PieceFactory func(
	kind common.Kind,
	color common.Color,
	position common.Position,
) common.Piece

// PieceStorageFactory ...
type PieceStorageFactory func(
	size common.Size,
	pieces []common.Piece,
) common.PieceStorage

const (
	minFileSymbol = 'a'
)

// DecodePosition ...
//
// It decodes a position from pure algebraic coordinate notation.
func DecodePosition(text string) (position common.Position, err error) {
	if len(text) != 2 {
		return common.Position{}, errors.New("incorrect length")
	}

	file := int(text[0]) - minFileSymbol
	if file < 0 {
		return common.Position{}, errors.New("incorrect file")
	}

	rank, err := strconv.Atoi(text[1:])
	if err != nil {
		return common.Position{}, fmt.Errorf("incorrect rank: %s", err)
	}
	rank--

	return common.Position{File: file, Rank: rank}, nil
}

// DecodeMove ...
//
// It decodes a move from pure algebraic coordinate notation.
func DecodeMove(text string) (move common.Move, err error) {
	if len(text) != 4 {
		return common.Move{}, errors.New("incorrect length")
	}

	start, err := DecodePosition(text[:2])
	if err != nil {
		return common.Move{}, fmt.Errorf("incorrect start: %s", err)
	}

	finish, err := DecodePosition(text[2:])
	if err != nil {
		return common.Move{}, fmt.Errorf("incorrect finish: %s", err)
	}

	return common.Move{Start: start, Finish: finish}, nil
}

// DecodePiece ...
//
// It decodes a piece from FEN (only a kind and a color, not a position).
func DecodePiece(fen rune, factory PieceFactory) (common.Piece, error) {
	var kind common.Kind
	switch unicode.ToLower(fen) {
	case 'k':
		kind = common.King
	case 'q':
		kind = common.Queen
	case 'r':
		kind = common.Rook
	case 'b':
		kind = common.Bishop
	case 'n':
		kind = common.Knight
	case 'p':
		kind = common.Pawn
	default:
		return nil, errors.New("unknown kind")
	}

	var color common.Color
	if unicode.IsLower(fen) {
		color = common.Black
	} else {
		color = common.White
	}

	piece := factory(kind, color, common.Position{})
	return piece, nil
}

// DecodePieceStorage ...
//
// It decodes a piece storage from FEN.
func DecodePieceStorage(
	fen string,
	pieceFactory PieceFactory,
	pieceStorageFactory PieceStorageFactory,
) (common.PieceStorage, error) {
	ranks := strings.Split(fen, "/")
	reverse(ranks)

	var pieces []common.Piece
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

	size := common.Size{Width: width, Height: len(ranks)}
	storage := pieceStorageFactory(size, pieces)
	return storage, nil
}

func decodeRank(
	index int,
	fen string,
	pieceFactory PieceFactory,
) (pieces []common.Piece, maxFile int, err error) {
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
			piece.ApplyPosition(common.Position{File: maxFile, Rank: index})
		pieces = append(pieces, placedPiece)

		maxFile++
	}

	return pieces, maxFile, nil
}
