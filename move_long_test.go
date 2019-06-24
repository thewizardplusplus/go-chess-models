// +build long

package chessmodels_test

import (
	"syscall"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

type PositionMap map[int]int

func TestPerft(test *testing.T) {
	type args struct {
		storage models.PieceStorage
		color   models.Color
		deep    int
	}
	type data struct {
		skip bool
		name string
		args args
		want int
	}

	_, longer := syscall.Getenv("LONGER")
	for index, data := range []data{
		data{
			name: "kings",
			args: args{
				storage: kings(),
				color:   models.White,
				deep:    0,
			},
			want: 1,
		},
		data{
			name: "kings",
			args: args{
				storage: kings(),
				color:   models.White,
				deep:    1,
			},
			want: 5,
		},
		data{
			name: "kings",
			args: args{
				storage: kings(),
				color:   models.White,
				deep:    2,
			},
			want: 25,
		},
		data{
			name: "kings",
			args: args{
				storage: kings(),
				color:   models.White,
				deep:    3,
			},
			want: 170,
		},
		data{
			name: "kings",
			args: args{
				storage: kings(),
				color:   models.White,
				deep:    4,
			},
			want: 1156,
		},
		data{
			name: "kings",
			args: args{
				storage: kings(),
				color:   models.White,
				deep:    5,
			},
			want: 7922,
		},
		data{
			name: "queens",
			args: args{
				storage: queens(),
				color:   models.White,
				deep:    0,
			},
			want: 1,
		},
		data{
			name: "queens",
			args: args{
				storage: queens(),
				color:   models.White,
				deep:    1,
			},
			want: 20,
		},
		data{
			name: "queens",
			args: args{
				storage: queens(),
				color:   models.White,
				deep:    2,
			},
			want: 301,
		},
		data{
			name: "queens",
			args: args{
				storage: queens(),
				color:   models.White,
				deep:    3,
			},
			want: 6063,
		},
		data{
			name: "rooks",
			args: args{
				storage: rooks(),
				color:   models.White,
				deep:    0,
			},
			want: 1,
		},
		data{
			name: "rooks",
			args: args{
				storage: rooks(),
				color:   models.White,
				deep:    1,
			},
			want: 24,
		},
		data{
			name: "rooks",
			args: args{
				storage: rooks(),
				color:   models.White,
				deep:    2,
			},
			want: 482,
		},
		data{
			name: "rooks",
			args: args{
				storage: rooks(),
				color:   models.White,
				deep:    3,
			},
			want: 11522,
		},
		data{
			name: "bishops",
			args: args{
				storage: bishops(),
				color:   models.White,
				deep:    0,
			},
			want: 1,
		},
		data{
			name: "bishops",
			args: args{
				storage: bishops(),
				color:   models.White,
				deep:    1,
			},
			want: 18,
		},
		data{
			name: "bishops",
			args: args{
				storage: bishops(),
				color:   models.White,
				deep:    2,
			},
			want: 305,
		},
		data{
			name: "bishops",
			args: args{
				storage: bishops(),
				color:   models.White,
				deep:    3,
			},
			want: 5575,
		},
		data{
			name: "initial",
			args: args{
				storage: initial(),
				color:   models.White,
				deep:    0,
			},
			want: 1,
		},
		data{
			name: "initial",
			args: args{
				storage: initial(),
				color:   models.White,
				deep:    1,
			},
			want: 12,
		},
		data{
			name: "initial",
			args: args{
				storage: initial(),
				color:   models.White,
				deep:    2,
			},
			want: 144,
		},
		data{
			name: "initial",
			args: args{
				storage: initial(),
				color:   models.White,
				deep:    3,
			},
			want: 2124,
		},
		data{
			name: "initial",
			skip: !longer,
			args: args{
				storage: initial(),
				color:   models.White,
				deep:    4,
			},
			want: 31250,
		},
		data{
			name: "kiwipete",
			args: args{
				storage: kiwipete(),
				color:   models.White,
				deep:    0,
			},
			want: 1,
		},
		data{
			name: "kiwipete",
			args: args{
				storage: kiwipete(),
				color:   models.White,
				deep:    1,
			},
			want: 44,
		},
		data{
			name: "kiwipete",
			args: args{
				storage: kiwipete(),
				color:   models.White,
				deep:    2,
			},
			want: 1740,
		},
		data{
			name: "kiwipete",
			skip: !longer,
			args: args{
				storage: kiwipete(),
				color:   models.White,
				deep:    3,
			},
			want: 77305,
		},
	} {
		if data.skip {
			continue
		}

		got := perft(
			data.args.storage,
			data.args.color,
			data.args.deep,
		)

		if got != data.want {
			test.Logf(
				"%s/#%d: %d/%d",
				data.name,
				index,
				got,
				data.want,
			)

			test.Fail()
		}
	}
}

func kings() models.Board {
	return models.NewBoard(
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
}

func queens() models.Board {
	return models.NewBoard(
		models.Size{8, 8},
		[]models.Piece{
			pieces.NewKing(
				models.Black,
				models.Position{4, 7},
			),
			pieces.NewQueen(
				models.Black,
				models.Position{3, 7},
			),
			pieces.NewKing(
				models.White,
				models.Position{4, 0},
			),
			pieces.NewQueen(
				models.White,
				models.Position{3, 0},
			),
		},
	)
}

func rooks() models.Board {
	return models.NewBoard(
		models.Size{8, 8},
		[]models.Piece{
			pieces.NewKing(
				models.Black,
				models.Position{4, 7},
			),
			pieces.NewRook(
				models.Black,
				models.Position{0, 7},
			),
			pieces.NewRook(
				models.Black,
				models.Position{7, 7},
			),
			pieces.NewKing(
				models.White,
				models.Position{4, 0},
			),
			pieces.NewRook(
				models.White,
				models.Position{0, 0},
			),
			pieces.NewRook(
				models.White,
				models.Position{7, 0},
			),
		},
	)
}

func bishops() models.Board {
	return models.NewBoard(
		models.Size{8, 8},
		[]models.Piece{
			pieces.NewKing(
				models.Black,
				models.Position{4, 7},
			),
			pieces.NewBishop(
				models.Black,
				models.Position{2, 7},
			),
			pieces.NewBishop(
				models.Black,
				models.Position{5, 7},
			),
			pieces.NewKing(
				models.White,
				models.Position{4, 0},
			),
			pieces.NewBishop(
				models.White,
				models.Position{2, 0},
			),
			pieces.NewBishop(
				models.White,
				models.Position{5, 0},
			),
		},
	)
}

func pawns(
	color models.Color,
	positions PositionMap,
) []models.Piece {
	var pawns []models.Piece
	for file, rank := range positions {
		pawns = append(pawns, pieces.NewPawn(
			color,
			models.Position{file, rank},
		))
	}

	return pawns
}

func initial() models.Board {
	restPieces := func(
		color models.Color,
		rank int,
	) []models.Piece {
		return []models.Piece{
			pieces.NewRook(
				color,
				models.Position{0, rank},
			),
			pieces.NewKnight(
				color,
				models.Position{1, rank},
			),
			pieces.NewBishop(
				color,
				models.Position{2, rank},
			),
			pieces.NewQueen(
				color,
				models.Position{3, rank},
			),
			pieces.NewKing(
				color,
				models.Position{4, rank},
			),
			pieces.NewBishop(
				color,
				models.Position{5, rank},
			),
			pieces.NewKnight(
				color,
				models.Position{6, rank},
			),
			pieces.NewRook(
				color,
				models.Position{7, rank},
			),
		}
	}

	var allPieces []models.Piece
	allPieces = append(allPieces, restPieces(
		models.Black,
		7,
	)...)
	allPieces = append(allPieces, pawns(
		models.Black,
		PositionMap{
			0: 6, 1: 6, 2: 6, 3: 6,
			4: 6, 5: 6, 6: 6, 7: 6,
		},
	)...)
	allPieces = append(allPieces, pawns(
		models.White,
		PositionMap{
			0: 1, 1: 1, 2: 1, 3: 1,
			4: 1, 5: 1, 6: 1, 7: 1,
		},
	)...)
	allPieces = append(allPieces, restPieces(
		models.White,
		0,
	)...)

	return models.NewBoard(
		models.Size{8, 8},
		allPieces,
	)
}

func kiwipete() models.Board {
	var allPieces []models.Piece
	allPieces = append(
		allPieces,
		[]models.Piece{
			// kings
			pieces.NewKing(
				models.Black,
				models.Position{4, 7},
			),
			pieces.NewKing(
				models.White,
				models.Position{4, 0},
			),

			// queens
			pieces.NewQueen(
				models.Black,
				models.Position{4, 6},
			),
			pieces.NewQueen(
				models.White,
				models.Position{5, 2},
			),

			// rooks
			pieces.NewRook(
				models.Black,
				models.Position{0, 7},
			),
			pieces.NewRook(
				models.Black,
				models.Position{7, 7},
			),
			pieces.NewRook(
				models.White,
				models.Position{0, 0},
			),
			pieces.NewRook(
				models.White,
				models.Position{7, 0},
			),

			// bishops
			pieces.NewBishop(
				models.Black,
				models.Position{0, 5},
			),
			pieces.NewBishop(
				models.Black,
				models.Position{6, 6},
			),
			pieces.NewBishop(
				models.White,
				models.Position{3, 1},
			),
			pieces.NewBishop(
				models.White,
				models.Position{4, 1},
			),

			// knights
			pieces.NewKnight(
				models.Black,
				models.Position{1, 5},
			),
			pieces.NewKnight(
				models.Black,
				models.Position{5, 5},
			),
			pieces.NewKnight(
				models.White,
				models.Position{2, 2},
			),
			pieces.NewKnight(
				models.White,
				models.Position{4, 4},
			),
		}...,
	)
	allPieces = append(allPieces, pawns(
		models.Black,
		PositionMap{
			0: 6, 1: 3, 2: 6, 3: 6,
			4: 5, 5: 6, 6: 5, 7: 2,
		},
	)...)
	allPieces = append(allPieces, pawns(
		models.White,
		PositionMap{
			0: 1, 1: 1, 2: 1, 3: 4,
			4: 3, 5: 1, 6: 1, 7: 1,
		},
	)...)

	return models.NewBoard(
		models.Size{8, 8},
		allPieces,
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
