package chessmodels

import (
	"strconv"
	"strings"
)

// PieceFactory ...
type PieceFactory func(
	kind Kind,
	color Color,
	position Position,
) (Piece, error)

// ParseBoard ...
func ParseBoard(
	boardInFEN string,
	pieceFactory PieceFactory,
) (PieceStorage, error) {
	ranks := strings.Split(boardInFEN, "/")
	// reverse ranks
	left, right := 0, len(ranks)-1
	for left < right {
		ranks[left], ranks[right] =
			ranks[right], ranks[left]
		left, right = left+1, right-1
	}

	var pieces []Piece
	var width int
	for index, rank := range ranks {
		rankPieces, rankWidth, err :=
			ParseRank(index, rank, pieceFactory)
		if err != nil {
			return nil, err
		}

		pieces = append(pieces, rankPieces...)
		if width < rankWidth {
			width = rankWidth
		}
	}

	size := Size{width, len(ranks)}
	board := NewBoard(size, pieces)
	return board, nil
}

// ParseRank ...
func ParseRank(
	rankIndex int,
	rankInFEN string,
	pieceFactory PieceFactory,
) (pieces []Piece, maxFile int, err error) {
	for _, symbol := range rankInFEN {
		kind, err := ParseKind(symbol)
		if err != nil {
			shift, err :=
				strconv.Atoi(string(symbol))
			if err != nil {
				return nil, 0, err
			}

			maxFile += shift
			continue
		}

		color := ParseColor(symbol)
		position :=
			Position{maxFile, rankIndex}
		piece, err := pieceFactory(
			kind,
			color,
			position,
		)
		if err != nil {
			return nil, 0, err
		}

		pieces = append(pieces, piece)
		maxFile++
	}

	return pieces, maxFile, nil
}
