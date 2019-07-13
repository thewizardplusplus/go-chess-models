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
) Piece

// ParseBoard ...
func ParseBoard(
	boardInFEN string,
	pieceFactory PieceFactory,
) (PieceStorage, error) {
	ranks := strings.Split(boardInFEN, "/")
	reverse(ranks)

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
		kind, color, err := ParsePiece(symbol)
		if err != nil {
			shift, err :=
				strconv.Atoi(string(symbol))
			if err != nil {
				return nil, 0, err
			}

			maxFile += shift
			continue
		}

		position := Position{maxFile, rankIndex}
		piece :=
			pieceFactory(kind, color, position)
		pieces = append(pieces, piece)
		maxFile++
	}

	return pieces, maxFile, nil
}

// ToFEN ...
func (board Board) ToFEN() (string, error) {
	var rank string
	var shift int
	resetShift := func() {
		if shift != 0 {
			rank += strconv.Itoa(shift)
			shift = 0
		}
	}

	var ranks []string
	positions := board.size.Positions()
	for _, position := range positions {
		piece, ok := board.Piece(position)
		if ok {
			kind, err := piece.Kind().
				ToFEN(piece.Color())
			if err != nil {
				return "", err
			}

			resetShift()
			rank += string(kind)
		} else {
			shift++
		}

		lastFile := board.size.Height - 1
		if position.File == lastFile {
			resetShift()
			ranks = append(ranks, rank)
			rank = ""
		}
	}

	reverse(ranks)
	return strings.Join(ranks, "/"), nil
}

func reverse(strings []string) {
	left, right := 0, len(strings)-1
	for left < right {
		strings[left], strings[right] =
			strings[right], strings[left]
		left, right = left+1, right-1
	}
}
