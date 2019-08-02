package uci

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func TestDecodePiece(test *testing.T) {
	type args struct {
		fen rune
	}
	type data struct {
		args      args
		wantPiece models.Piece
		wantErr   bool
	}

	for _, data := range []data{
		data{
			args: args{'K'},
			wantPiece: pieces.NewKing(
				models.White,
				models.Position{},
			),
			wantErr: false,
		},
		data{
			args: args{'q'},
			wantPiece: pieces.NewQueen(
				models.Black,
				models.Position{},
			),
			wantErr: false,
		},
		data{
			args:      args{'a'},
			wantPiece: nil,
			wantErr:   true,
		},
	} {
		gotPiece, gotErr := DecodePiece(
			data.args.fen,
			pieces.NewPiece,
		)

		if !reflect.DeepEqual(
			gotPiece,
			data.wantPiece,
		) {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}

func TestDecodeRank(test *testing.T) {
	type args struct {
		index int
		fen   string
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
				index: 7,
				fen:   "2",
			},
			wantPieces:  nil,
			wantMaxFile: 2,
			wantErr:     false,
		},
		data{
			args: args{
				index: 7,
				fen:   "K",
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
				index: 7,
				fen:   "2K",
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
				index: 7,
				fen:   "2Kq",
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
				index: 7,
				fen:   "2K3q",
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
				index: 7,
				fen:   "2K3q4",
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
				index: 7,
				fen:   "2K#q4",
			},
			wantPieces:  nil,
			wantMaxFile: 0,
			wantErr:     true,
		},
	} {
		gotPieces, gotMaxFile, gotErr :=
			decodeRank(
				data.args.index,
				data.args.fen,
				pieces.NewPiece,
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
