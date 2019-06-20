// +build long

package chessmodels_test

import (
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func TestPerft(test *testing.T) {
	type args struct {
		storage models.PieceStorage
		color   models.Color
		deep    int
	}
	type data struct {
		args args
		want int
	}

	for _, data := range []data{
		data{
			args: args{
				storage: initial(),
				color:   models.White,
				deep:    0,
			},
			want: 1,
		},
		data{
			args: args{
				storage: initial(),
				color:   models.White,
				deep:    1,
			},
			want: 12,
		},
		data{
			args: args{
				storage: initial(),
				color:   models.White,
				deep:    2,
			},
			want: 144,
		},
		data{
			args: args{
				storage: initial(),
				color:   models.White,
				deep:    3,
			},
			want: 2124,
		},
		/*data{
		  args: args{
		    storage: initial(),
		    color:   models.White,
		    deep:    4,
		  },
		  want: 31250,
		},*/
	} {
		got := perft(
			data.args.storage,
			data.args.color,
			data.args.deep,
		)

		if got != data.want {
			test.Log(got, data.want)
			test.Fail()
		}
	}
}

func initial() models.Board {
	pawns := func(
		color models.Color,
		rank int,
	) []models.Piece {
		var pawns []models.Piece
		for file := 0; file < 8; file++ {
			pawns = append(pawns, pieces.NewPawn(
				color,
				models.Position{file, rank},
			))
		}

		return pawns
	}

	minorPieces := func(
		color models.Color,
		files []int,
		rank int,
	) []models.Piece {
		return []models.Piece{
			pieces.NewRook(
				color,
				models.Position{files[0], rank},
			),
			pieces.NewKnight(
				color,
				models.Position{files[1], rank},
			),
			pieces.NewBishop(
				color,
				models.Position{files[2], rank},
			),
		}
	}

	restPieces := func(
		color models.Color,
		rank int,
	) []models.Piece {
		var restPieces []models.Piece
		restPieces = append(
			restPieces,
			minorPieces(
				color,
				[]int{0, 1, 2},
				rank,
			)...,
		)
		restPieces = append(
			restPieces,
			pieces.NewQueen(
				color,
				models.Position{3, rank},
			),
		)
		restPieces = append(
			restPieces,
			pieces.NewKing(
				color,
				models.Position{4, rank},
			),
		)
		restPieces = append(
			restPieces,
			minorPieces(
				color,
				[]int{7, 6, 5},
				rank,
			)...,
		)

		return restPieces
	}

	var pieces []models.Piece
	pieces = append(pieces, restPieces(
		models.Black,
		7,
	)...)
	pieces = append(pieces, pawns(
		models.Black,
		6,
	)...)
	pieces = append(pieces, pawns(
		models.White,
		1,
	)...)
	pieces = append(pieces, restPieces(
		models.White,
		0,
	)...)

	return models.NewBoard(
		models.Size{8, 8},
		pieces,
	)
}

func kiwipete() models.Board {
	type cons func(
		models.Color,
		models.Position,
	) models.Piece
	type positions map[int]int
	type colors map[models.Color]positions
	type kinds map[cons]colors

	makePieces := func(
		kinds kinds,
	) []models.Piece {
		var pieces []models.Piece
		for cons, colors := range kinds {
			for color, positions := range colors {
				for file, rank := range positions {
					pieces = append(pieces, cons(
						color,
						models.Position{file, rank},
					))
				}
			}
		}

		return pieces
	}

	return models.NewBoard(
		model.Size{8, 8},
		makePieces(kinds{
			models.NewKing: colors{
				models.Black: positions{4: 7},
				models.White: positions{4: 0},
			},
			models.NewQueen: colors{
				models.Black: positions{4: 6},
				models.White: positions{5: 2},
			},
			models.NewRook: colors{
				models.Black: positions{0: 7, 7: 7},
				models.White: positions{0: 0, 7: 0},
			},
			models.NewBishop: colors{
				models.Black: positions{0: 5, 6: 6},
				models.White: positions{3: 1, 4: 1},
			},
			models.NewKnight: colors{
				models.Black: positions{0: 7, 7: 7},
				models.White: positions{0: 0, 7: 0},
			},
			models.NewPawn: colors{
				models.Black: positions{0: 7, 7: 7},
				models.White: positions{0: 0, 7: 0},
			},
		}),
	)
}

func perft(
	storage models.PieceStorage,
	color models.Color,
	deep int,
) int {
	// check for a check should be first,
	// including before a termination check,
	// because a terminated evaluation
	// doesn't make sense for a check position
	moves, err := models.MoveGenerator{}.
		MovesForColor(storage, color)
	if err != nil {
		return 0
	}

	if deep == 0 {
		return 1
	}

	var count int
	for _, move := range moves {
		nextStorage := storage.ApplyMove(move)
		nextColor := color.Negative()
		count += perft(
			nextStorage,
			nextColor,
			deep-1,
		)
	}

	return count
}
