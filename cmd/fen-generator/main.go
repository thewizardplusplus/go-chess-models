package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

var (
	kindsInFEN = map[models.Kind]rune{
		models.King:   'k',
		models.Queen:  'q',
		models.Rook:   'r',
		models.Bishop: 'b',
		models.Knight: 'n',
		models.Pawn:   'p',
	}
)

func main() {
	storage := models.NewBoard(
		models.Size{8, 8},
		[]models.Piece{
			pieces.NewKing(
				models.Black,
				models.Position{4, 7},
			),
			pieces.NewKing(
				models.White,
				models.Position{4, 0},
			),
		},
	)

	var ranksInFEN []string
	width := storage.Size().Width
	height := storage.Size().Height
	for rank := 0; rank < height; rank++ {
		var rankInFEN string
		var number int
		for file := 0; file < width; file++ {
			position := models.Position{file, rank}
			piece, ok := storage.Piece(position)
			if !ok {
				number++
				continue
			}

			if number != 0 {
				rankInFEN += strconv.Itoa(number)
				number = 0
			}

			kindInFEN := kindsInFEN[piece.Kind()]
			if piece.Color() == models.White {
				kindInFEN =
					unicode.ToUpper(kindInFEN)
			}

			rankInFEN += string(kindInFEN)
		}

		if number != 0 {
			rankInFEN += strconv.Itoa(number)
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

	fen := strings.Join(ranksInFEN, "/")
	fmt.Println(fen)
}
