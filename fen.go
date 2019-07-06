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

// ToFEN ...
func (board Board) ToFEN() (string, error) {
	var ranksInFEN []string
	width := board.size.Width
	height := board.size.Height
	for rank := 0; rank < height; rank++ {
		var rankInFEN string
		var shift int
		for file := 0; file < width; file++ {
			position := Position{file, rank}
			piece, ok := board.Piece(position)
			if !ok {
				shift++
				continue
			}

			if shift != 0 {
				rankInFEN += strconv.Itoa(shift)
				shift = 0
			}

			kindInFEN, err := piece.Kind().
				ToFEN(piece.Color())
			if err != nil {
				return "", err
			}

			rankInFEN += string(kindInFEN)
		}

		if shift != 0 {
			rankInFEN += strconv.Itoa(shift)
		}

		ranksInFEN = append(
			ranksInFEN,
			rankInFEN,
		)
	}

	reverse(ranksInFEN)
	return strings.Join(ranksInFEN, "/"), nil
}

func reverse(strings []string) {
	left, right := 0, len(strings)-1
	for left < right {
		strings[left], strings[right] =
			strings[right], strings[left]
		left, right = left+1, right-1
	}
}
