package uci

import (
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func TestDecodePosition(test *testing.T) {
	type args struct {
		text string
	}
	type data struct {
		args         args
		wantPosition common.Position
		wantErr      bool
	}

	for _, data := range []data{
		{
			args: args{"e2"},
			wantPosition: common.Position{
				File: 4,
				Rank: 1,
			},
			wantErr: false,
		},
		{
			args:         args{"e"},
			wantPosition: common.Position{},
			wantErr:      true,
		},
		{
			args:         args{"e23"},
			wantPosition: common.Position{},
			wantErr:      true,
		},
		{
			args:         args{"\n2"},
			wantPosition: common.Position{},
			wantErr:      true,
		},
		{
			args:         args{"e\n"},
			wantPosition: common.Position{},
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
		wantMove common.Move
		wantErr  bool
	}

	for _, data := range []data{
		{
			args: args{"e2e4"},
			wantMove: common.Move{
				Start: common.Position{
					File: 4,
					Rank: 1,
				},
				Finish: common.Position{
					File: 4,
					Rank: 3,
				},
			},
			wantErr: false,
		},
		{
			args:     args{"e2e"},
			wantMove: common.Move{},
			wantErr:  true,
		},
		{
			args:     args{"e2e42"},
			wantMove: common.Move{},
			wantErr:  true,
		},
		{
			args:     args{"e\ne4"},
			wantMove: common.Move{},
			wantErr:  true,
		},
		{
			args:     args{"e2e\n"},
			wantMove: common.Move{},
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
		wantPiece common.Piece
		wantErr   bool
	}

	for _, data := range []data{
		{
			args: args{'K'},
			wantPiece: pieces.NewKing(
				common.White,
				common.Position{},
			),
			wantErr: false,
		},
		{
			args: args{'q'},
			wantPiece: pieces.NewQueen(
				common.Black,
				common.Position{},
			),
			wantErr: false,
		},
		{
			args: args{'R'},
			wantPiece: pieces.NewRook(
				common.White,
				common.Position{},
			),
			wantErr: false,
		},
		{
			args: args{'b'},
			wantPiece: pieces.NewBishop(
				common.Black,
				common.Position{},
			),
			wantErr: false,
		},
		{
			args: args{'N'},
			wantPiece: pieces.NewKnight(
				common.White,
				common.Position{},
			),
			wantErr: false,
		},
		{
			args: args{'p'},
			wantPiece: pieces.NewPawn(
				common.Black,
				common.Position{},
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
		wantStorage common.PieceStorage
		wantErr     bool
	}

	for _, data := range []data{
		{
			args: args{
				fen: "2K3q/8/pp1R",
			},
			wantStorage: boards.NewMapBoard(
				common.Size{
					Width:  8,
					Height: 3,
				},
				[]common.Piece{
					pieces.NewPawn(common.Black, common.Position{
						File: 0,
						Rank: 0,
					}),
					pieces.NewPawn(common.Black, common.Position{
						File: 1,
						Rank: 0,
					}),
					pieces.NewRook(common.White, common.Position{
						File: 3,
						Rank: 0,
					}),
					pieces.NewKing(common.White, common.Position{
						File: 2,
						Rank: 2,
					}),
					pieces.NewQueen(common.Black, common.Position{
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
			wantStorage: boards.NewMapBoard(
				common.Size{
					Width:  8,
					Height: 4,
				},
				[]common.Piece{
					pieces.NewPawn(common.Black, common.Position{
						File: 0,
						Rank: 0,
					}),
					pieces.NewPawn(common.Black, common.Position{
						File: 1,
						Rank: 0,
					}),
					pieces.NewRook(common.White, common.Position{
						File: 3,
						Rank: 0,
					}),
					pieces.NewKing(common.White, common.Position{
						File: 2,
						Rank: 2,
					}),
					pieces.NewQueen(common.Black, common.Position{
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
			wantStorage: boards.NewMapBoard(
				common.Size{
					Width:  8,
					Height: 4,
				},
				[]common.Piece{
					pieces.NewPawn(common.Black, common.Position{
						File: 0,
						Rank: 1,
					}),
					pieces.NewPawn(common.Black, common.Position{
						File: 1,
						Rank: 1,
					}),
					pieces.NewRook(common.White, common.Position{
						File: 3,
						Rank: 1,
					}),
					pieces.NewKing(common.White, common.Position{
						File: 2,
						Rank: 3,
					}),
					pieces.NewQueen(common.Black, common.Position{
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
			DecodePieceStorage(data.args.fen, pieces.NewPiece, boards.NewMapBoard)

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
		wantPieces  []common.Piece
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
			wantPieces: []common.Piece{
				pieces.NewKing(common.White, common.Position{
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
			wantPieces: []common.Piece{
				pieces.NewKing(common.White, common.Position{
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
			wantPieces: []common.Piece{
				pieces.NewKing(common.White, common.Position{
					File: 2,
					Rank: 7,
				}),
				pieces.NewQueen(common.Black, common.Position{
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
			wantPieces: []common.Piece{
				pieces.NewKing(common.White, common.Position{
					File: 2,
					Rank: 7,
				}),
				pieces.NewQueen(common.Black, common.Position{
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
			wantPieces: []common.Piece{
				pieces.NewKing(common.White, common.Position{
					File: 2,
					Rank: 7,
				}),
				pieces.NewQueen(common.Black, common.Position{
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
