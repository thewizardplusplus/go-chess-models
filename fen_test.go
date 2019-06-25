package chessmodels_test

import (
	"errors"
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

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
		data{
			args: args{
				rankIndex: 7,
				rankInFEN: "2K3q4",
				pieceFactory: func(
					kind models.Kind,
					color models.Color,
					position models.Position,
				) (models.Piece, error) {
					return nil,
						errors.New("dummy error")
				},
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
