package chessmodels

import (
	"strconv"
	"strings"
)

// PieceParser ...
type PieceParser func(
	fen rune,
) (Piece, error)

// PieceStorageFactory ...
type PieceStorageFactory func(
	size Size,
	pieces []Piece,
) PieceStorage

// ParseBoard ...
func ParseBoard(
	fen string,
	pieceParser PieceParser,
	pieceStorageFactory PieceStorageFactory,
) (PieceStorage, error) {
	ranks := strings.Split(fen, "/")
	reverse(ranks)

	var pieces []Piece
	var width int
	for index, rank := range ranks {
		rankPieces, rankWidth, err :=
			parseRank(index, rank, pieceParser)
		if err != nil {
			return Board{}, err
		}

		pieces = append(pieces, rankPieces...)
		if width < rankWidth {
			width = rankWidth
		}
	}

	size := Size{width, len(ranks)}
	storage :=
		pieceStorageFactory(size, pieces)
	return storage, nil
}

// ParseDefaultBoard ...
func ParseDefaultBoard(
	fen string,
	pieceParser PieceParser,
) (PieceStorage, error) {
	return ParseBoard(
		fen,
		pieceParser,
		func(
			size Size,
			pieces []Piece,
		) PieceStorage {
			return NewBoard(size, pieces)
		},
	)
}

// String ...
//
// It converts the board to FEN.
func (board Board) String() string {
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
			resetShift()

			rank += piece.String()
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
	return strings.Join(ranks, "/")
}

func reverse(strings []string) {
	left, right := 0, len(strings)-1
	for left < right {
		strings[left], strings[right] =
			strings[right], strings[left]
		left, right = left+1, right-1
	}
}

func parseRank(
	index int,
	fen string,
	pieceParser PieceParser,
) (pieces []Piece, maxFile int, err error) {
	for _, symbol := range fen {
		piece, err := pieceParser(symbol)
		if err != nil {
			shift, err :=
				strconv.Atoi(string(symbol))
			if err != nil {
				return nil, 0, err
			}

			maxFile += shift
			continue
		}

		position := Position{maxFile, index}
		piece = piece.ApplyPosition(position)
		pieces = append(pieces, piece)

		maxFile++
	}

	return pieces, maxFile, nil
}
