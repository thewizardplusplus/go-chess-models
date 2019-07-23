package chessmodels_test

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func TestParseBoard(test *testing.T) {
	type args struct {
		boardInFEN   string
		pieceFactory models.PieceFactory
	}
	type data struct {
		args        args
		wantStorage models.PieceStorage
		wantErr     bool
	}

	for _, data := range []data{
		data{
			args: args{
				boardInFEN:   "2K3q/8/pp1R",
				pieceFactory: pieces.NewPiece,
			},
			wantStorage: models.NewBoard(
				models.Size{8, 3},
				[]models.Piece{
					pieces.NewPawn(
						models.Black,
						models.Position{0, 0},
					),
					pieces.NewPawn(
						models.Black,
						models.Position{1, 0},
					),
					pieces.NewRook(
						models.White,
						models.Position{3, 0},
					),
					pieces.NewKing(
						models.White,
						models.Position{2, 2},
					),
					pieces.NewQueen(
						models.Black,
						models.Position{6, 2},
					),
				},
			),
			wantErr: false,
		},
		data{
			args: args{
				boardInFEN:   "1/2K3q/8/pp1R",
				pieceFactory: pieces.NewPiece,
			},
			wantStorage: models.NewBoard(
				models.Size{8, 4},
				[]models.Piece{
					pieces.NewPawn(
						models.Black,
						models.Position{0, 0},
					),
					pieces.NewPawn(
						models.Black,
						models.Position{1, 0},
					),
					pieces.NewRook(
						models.White,
						models.Position{3, 0},
					),
					pieces.NewKing(
						models.White,
						models.Position{2, 2},
					),
					pieces.NewQueen(
						models.Black,
						models.Position{6, 2},
					),
				},
			),
			wantErr: false,
		},
		data{
			args: args{
				boardInFEN:   "2K3q/8/pp1R/1",
				pieceFactory: pieces.NewPiece,
			},
			wantStorage: models.NewBoard(
				models.Size{8, 4},
				[]models.Piece{
					pieces.NewPawn(
						models.Black,
						models.Position{0, 1},
					),
					pieces.NewPawn(
						models.Black,
						models.Position{1, 1},
					),
					pieces.NewRook(
						models.White,
						models.Position{3, 1},
					),
					pieces.NewKing(
						models.White,
						models.Position{2, 3},
					),
					pieces.NewQueen(
						models.Black,
						models.Position{6, 3},
					),
				},
			),
			wantErr: false,
		},
		data{
			args: args{
				boardInFEN:   "2K3q/#/pp1R",
				pieceFactory: pieces.NewPiece,
			},
			wantStorage: models.Board{},
			wantErr:     true,
		},
	} {
		gotStorage, gotErr :=
			models.ParseBoard(
				data.args.boardInFEN,
				data.args.pieceFactory,
			)

		if !reflect.DeepEqual(
			gotStorage,
			data.wantStorage,
		) {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}

func TestParseRank(test *testing.T) {
	type args struct {
		rankIndex    int
		rankInFEN    string
		pieceFactory models.PieceFactory
	}
	type data struct {
		args        args
		wantPieces  []models.Piece
		wantMaxFile int
		wantErr     bool
	}

	for _, data := range []data{
		data{
			args: args{
				rankIndex:    7,
				rankInFEN:    "2",
				pieceFactory: pieces.NewPiece,
			},
			wantPieces:  nil,
			wantMaxFile: 2,
			wantErr:     false,
		},
		data{
			args: args{
				rankIndex:    7,
				rankInFEN:    "K",
				pieceFactory: pieces.NewPiece,
			},
			wantPieces: []models.Piece{
				pieces.NewKing(
					models.White,
					models.Position{0, 7},
				),
			},
			wantMaxFile: 1,
			wantErr:     false,
		},
		data{
			args: args{
				rankIndex:    7,
				rankInFEN:    "2K",
				pieceFactory: pieces.NewPiece,
			},
			wantPieces: []models.Piece{
				pieces.NewKing(
					models.White,
					models.Position{2, 7},
				),
			},
			wantMaxFile: 3,
			wantErr:     false,
		},
		data{
			args: args{
				rankIndex:    7,
				rankInFEN:    "2Kq",
				pieceFactory: pieces.NewPiece,
			},
			wantPieces: []models.Piece{
				pieces.NewKing(
					models.White,
					models.Position{2, 7},
				),
				pieces.NewQueen(
					models.Black,
					models.Position{3, 7},
				),
			},
			wantMaxFile: 4,
			wantErr:     false,
		},
		data{
			args: args{
				rankIndex:    7,
				rankInFEN:    "2K3q",
				pieceFactory: pieces.NewPiece,
			},
			wantPieces: []models.Piece{
				pieces.NewKing(
					models.White,
					models.Position{2, 7},
				),
				pieces.NewQueen(
					models.Black,
					models.Position{6, 7},
				),
			},
			wantMaxFile: 7,
			wantErr:     false,
		},
		data{
			args: args{
				rankIndex:    7,
				rankInFEN:    "2K3q4",
				pieceFactory: pieces.NewPiece,
			},
			wantPieces: []models.Piece{
				pieces.NewKing(
					models.White,
					models.Position{2, 7},
				),
				pieces.NewQueen(
					models.Black,
					models.Position{6, 7},
				),
			},
			wantMaxFile: 11,
			wantErr:     false,
		},
		data{
			args: args{
				rankIndex:    7,
				rankInFEN:    "2K#q4",
				pieceFactory: pieces.NewPiece,
			},
			wantPieces:  nil,
			wantMaxFile: 0,
			wantErr:     true,
		},
	} {
		gotPieces, gotMaxFile, gotErr :=
			models.ParseRank(
				data.args.rankIndex,
				data.args.rankInFEN,
				data.args.pieceFactory,
			)

		if !reflect.DeepEqual(
			gotPieces,
			data.wantPieces,
		) {
			test.Fail()
		}
		if gotMaxFile != data.wantMaxFile {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}

func TestBoardString(test *testing.T) {
	type fields struct {
		size   models.Size
		pieces []models.Piece
	}
	type data struct {
		fields fields
		want   string
	}

	for _, data := range []data{
		data{
			fields: fields{
				size:   models.Size{5, 5},
				pieces: nil,
			},
			want: "5/5/5/5/5",
		},
		data{
			fields: fields{
				size: models.Size{5, 5},
				pieces: []models.Piece{
					pieces.NewKing(
						models.White,
						models.Position{0, 2},
					),
				},
			},
			want: "5/5/K4/5/5",
		},
		data{
			fields: fields{
				size: models.Size{5, 5},
				pieces: []models.Piece{
					pieces.NewKing(
						models.White,
						models.Position{1, 2},
					),
				},
			},
			want: "5/5/1K3/5/5",
		},
		data{
			fields: fields{
				size: models.Size{5, 5},
				pieces: []models.Piece{
					pieces.NewKing(
						models.White,
						models.Position{1, 2},
					),
					pieces.NewQueen(
						models.Black,
						models.Position{2, 2},
					),
				},
			},
			want: "5/5/1Kq2/5/5",
		},
		data{
			fields: fields{
				size: models.Size{5, 5},
				pieces: []models.Piece{
					pieces.NewKing(
						models.White,
						models.Position{1, 2},
					),
					pieces.NewQueen(
						models.Black,
						models.Position{4, 2},
					),
				},
			},
			want: "5/5/1K2q/5/5",
		},
		data{
			fields: fields{
				size: models.Size{5, 5},
				pieces: []models.Piece{
					pieces.NewKing(
						models.White,
						models.Position{0, 3},
					),
					pieces.NewQueen(
						models.Black,
						models.Position{1, 2},
					),
					pieces.NewQueen(
						models.White,
						models.Position{2, 2},
					),
					pieces.NewRook(
						models.Black,
						models.Position{1, 1},
					),
					pieces.NewRook(
						models.White,
						models.Position{4, 1},
					),
				},
			},
			want: "5/K4/1qQ2/1r2R/5",
		},
	} {
		storage := models.NewBoard(
			data.fields.size,
			data.fields.pieces,
		)
		got := storage.String()

		if got != data.want {
			test.Fail()
		}
	}
}
