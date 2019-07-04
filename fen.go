package chessmodels

import (
	"strconv"
	"strings"
	"unicode"
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

// String ...
//
// It converts the board to FEN.
func (board Board) String() string {
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

			kindsInFEN := map[Kind]rune{
				King:   'k',
				Queen:  'q',
				Rook:   'r',
				Bishop: 'b',
				Knight: 'n',
				Pawn:   'p',
			}
			kindInFEN := kindsInFEN[piece.Kind()]
			if piece.Color() == White {
				kindInFEN =
					unicode.ToUpper(kindInFEN)
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

	// reverse ranks
	left, right := 0, len(ranksInFEN)-1
	for left < right {
		ranksInFEN[left], ranksInFEN[right] =
			ranksInFEN[right], ranksInFEN[left]
		left, right = left+1, right-1
	}

	return strings.Join(ranksInFEN, "/")
}
