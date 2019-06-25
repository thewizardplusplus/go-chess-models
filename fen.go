package chessmodels

import (
	"strconv"
)

// PieceFactory ...
type PieceFactory func(
	kind Kind,
	color Color,
	position Position,
) (Piece, error)

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
