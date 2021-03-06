package uci

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func TestDecodePosition(test *testing.T) {
	type args struct {
		text string
	}
	type data struct {
		args         args
		wantPosition models.Position
		wantErr      bool
	}

	for _, data := range []data{
		{
			args: args{"e2"},
			wantPosition: models.Position{
				File: 4,
				Rank: 1,
			},
			wantErr: false,
		},
		{
			args:         args{"e"},
			wantPosition: models.Position{},
			wantErr:      true,
		},
		{
			args:         args{"e23"},
			wantPosition: models.Position{},
			wantErr:      true,
		},
		{
			args:         args{"\n2"},
			wantPosition: models.Position{},
			wantErr:      true,
		},
		{
			args:         args{"e\n"},
			wantPosition: models.Position{},
			wantErr:      true,
		},
	} {
		gotPosition, gotErr := DecodePosition(data.args.text)

		if !reflect.DeepEqual(gotPosition, data.wantPosition) {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}

func TestDecodeMove(test *testing.T) {
	type args struct {
		text string
	}
	type data struct {
		args     args
		wantMove models.Move
		wantErr  bool
	}

	for _, data := range []data{
		{
			args: args{"e2e4"},
			wantMove: models.Move{
				Start: models.Position{
					File: 4,
					Rank: 1,
				},
				Finish: models.Position{
					File: 4,
					Rank: 3,
				},
			},
			wantErr: false,
		},
		{
			args:     args{"e2e"},
			wantMove: models.Move{},
			wantErr:  true,
		},
		{
			args:     args{"e2e42"},
			wantMove: models.Move{},
			wantErr:  true,
		},
		{
			args:     args{"e\ne4"},
			wantMove: models.Move{},
			wantErr:  true,
		},
		{
			args:     args{"e2e\n"},
			wantMove: models.Move{},
			wantErr:  true,
		},
	} {
		gotMove, gotErr := DecodeMove(data.args.text)

		if !reflect.DeepEqual(gotMove, data.wantMove) {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}

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
		{
			args: args{'K'},
			wantPiece: pieces.NewKing(
				models.White,
				models.Position{},
			),
			wantErr: false,
		},
		{
			args: args{'q'},
			wantPiece: pieces.NewQueen(
				models.Black,
				models.Position{},
			),
			wantErr: false,
		},
		{
			args:      args{'a'},
			wantPiece: nil,
			wantErr:   true,
		},
	} {
		gotPiece, gotErr := DecodePiece(data.args.fen, pieces.NewPiece)

		if !reflect.DeepEqual(gotPiece, data.wantPiece) {
			test.Fail()
		}

		hasErr := gotErr != nil
		if hasErr != data.wantErr {
			test.Fail()
		}
	}
}

func TestDecodePieceStorage(test *testing.T) {
	type args struct {
		fen string
	}
	type data struct {
		args        args
		wantStorage models.PieceStorage
		wantErr     bool
	}

	for _, data := range []data{
		{
			args: args{
				fen: "2K3q/8/pp1R",
			},
			wantStorage: models.NewBoard(
				models.Size{
					Width:  8,
					Height: 3,
				},
				[]models.Piece{
					pieces.NewPawn(models.Black, models.Position{
						File: 0,
						Rank: 0,
					}),
					pieces.NewPawn(models.Black, models.Position{
						File: 1,
						Rank: 0,
					}),
					pieces.NewRook(models.White, models.Position{
						File: 3,
						Rank: 0,
					}),
					pieces.NewKing(models.White, models.Position{
						File: 2,
						Rank: 2,
					}),
					pieces.NewQueen(models.Black, models.Position{
						File: 6,
						Rank: 2,
					}),
				},
			),
			wantErr: false,
		},
		{
			args: args{
				fen: "1/2K3q/8/pp1R",
			},
			wantStorage: models.NewBoard(
				models.Size{
					Width:  8,
					Height: 4,
				},
				[]models.Piece{
					pieces.NewPawn(models.Black, models.Position{
						File: 0,
						Rank: 0,
					}),
					pieces.NewPawn(models.Black, models.Position{
						File: 1,
						Rank: 0,
					}),
					pieces.NewRook(models.White, models.Position{
						File: 3,
						Rank: 0,
					}),
					pieces.NewKing(models.White, models.Position{
						File: 2,
						Rank: 2,
					}),
					pieces.NewQueen(models.Black, models.Position{
						File: 6,
						Rank: 2,
					}),
				},
			),
			wantErr: false,
		},
		{
			args: args{
				fen: "2K3q/8/pp1R/1",
			},
			wantStorage: models.NewBoard(
				models.Size{
					Width:  8,
					Height: 4,
				},
				[]models.Piece{
					pieces.NewPawn(models.Black, models.Position{
						File: 0,
						Rank: 1,
					}),
					pieces.NewPawn(models.Black, models.Position{
						File: 1,
						Rank: 1,
					}),
					pieces.NewRook(models.White, models.Position{
						File: 3,
						Rank: 1,
					}),
					pieces.NewKing(models.White, models.Position{
						File: 2,
						Rank: 3,
					}),
					pieces.NewQueen(models.Black, models.Position{
						File: 6,
						Rank: 3,
					}),
				},
			),
			wantErr: false,
		},
		{
			args: args{
				fen: "2K3q/#/pp1R",
			},
			wantStorage: nil,
			wantErr:     true,
		},
	} {
		gotStorage, gotErr :=
			DecodePieceStorage(data.args.fen, pieces.NewPiece, models.NewBoard)

		if !reflect.DeepEqual(gotStorage, data.wantStorage) {
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
		{
			args: args{
				index: 7,
				fen:   "2",
			},
			wantPieces:  nil,
			wantMaxFile: 2,
			wantErr:     false,
		},
		{
			args: args{
				index: 7,
				fen:   "K",
			},
			wantPieces: []models.Piece{
				pieces.NewKing(models.White, models.Position{
					File: 0,
					Rank: 7,
				}),
			},
			wantMaxFile: 1,
			wantErr:     false,
		},
		{
			args: args{
				index: 7,
				fen:   "2K",
			},
			wantPieces: []models.Piece{
				pieces.NewKing(models.White, models.Position{
					File: 2,
					Rank: 7,
				}),
			},
			wantMaxFile: 3,
			wantErr:     false,
		},
		{
			args: args{
				index: 7,
				fen:   "2Kq",
			},
			wantPieces: []models.Piece{
				pieces.NewKing(models.White, models.Position{
					File: 2,
					Rank: 7,
				}),
				pieces.NewQueen(models.Black, models.Position{
					File: 3,
					Rank: 7,
				}),
			},
			wantMaxFile: 4,
			wantErr:     false,
		},
		{
			args: args{
				index: 7,
				fen:   "2K3q",
			},
			wantPieces: []models.Piece{
				pieces.NewKing(models.White, models.Position{
					File: 2,
					Rank: 7,
				}),
				pieces.NewQueen(models.Black, models.Position{
					File: 6,
					Rank: 7,
				}),
			},
			wantMaxFile: 7,
			wantErr:     false,
		},
		{
			args: args{
				index: 7,
				fen:   "2K3q4",
			},
			wantPieces: []models.Piece{
				pieces.NewKing(models.White, models.Position{
					File: 2,
					Rank: 7,
				}),
				pieces.NewQueen(models.Black, models.Position{
					File: 6,
					Rank: 7,
				}),
			},
			wantMaxFile: 11,
			wantErr:     false,
		},
		{
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
			decodeRank(data.args.index, data.args.fen, pieces.NewPiece)

		if !reflect.DeepEqual(gotPieces, data.wantPieces) {
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
