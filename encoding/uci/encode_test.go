package uci

import (
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

type MockPiece struct {
	kind  common.Kind
	color common.Color
}

func (piece MockPiece) Kind() common.Kind {
	return piece.kind
}

func (piece MockPiece) Color() common.Color {
	return piece.color
}

func (piece MockPiece) Position() common.Position {
	panic("not implemented")
}

func (piece MockPiece) ApplyPosition(position common.Position) common.Piece {
	panic("not implemented")
}

func (piece MockPiece) CheckMove(
	move common.Move,
	storage common.PieceStorage,
) bool {
	panic("not implemented")
}

type MockPieceStorage struct {
	size common.Size

	piece func(position common.Position) (piece common.Piece, ok bool)
}

func (storage MockPieceStorage) Size() common.Size {
	return storage.size
}

func (storage MockPieceStorage) Piece(position common.Position) (
	piece common.Piece,
	ok bool,
) {
	if storage.piece == nil {
		panic("not implemented")
	}

	return storage.piece(position)
}

func (storage MockPieceStorage) Pieces() []common.Piece {
	panic("not implemented")
}

func (storage MockPieceStorage) ApplyMove(
	move common.Move,
) common.PieceStorage {
	panic("not implemented")
}

func (storage MockPieceStorage) CheckMove(move common.Move) error {
	panic("not implemented")
}

func TestEncodePosition(test *testing.T) {
	type args struct {
		position common.Position
	}
	type data struct {
		args args
		want string
	}

	for _, data := range []data{
		{
			args: args{
				position: common.Position{
					File: 2,
					Rank: 1,
				},
			},
			want: "c2",
		},
		{
			args: args{
				position: common.Position{
					File: 5,
					Rank: 6,
				},
			},
			want: "f7",
		},
	} {
		got := EncodePosition(data.args.position)

		if got != data.want {
			test.Fail()
		}
	}
}

func TestEncodeMove(test *testing.T) {
	type args struct {
		move common.Move
	}
	type data struct {
		args args
		want string
	}

	for _, data := range []data{
		{
			args: args{
				move: common.Move{
					Start: common.Position{
						File: 2,
						Rank: 1,
					},
					Finish: common.Position{
						File: 2,
						Rank: 3,
					},
				},
			},
			want: "c2c4",
		},
		{
			args: args{
				move: common.Move{
					Start: common.Position{
						File: 5,
						Rank: 6,
					},
					Finish: common.Position{
						File: 5,
						Rank: 4,
					},
				},
			},
			want: "f7f5",
		},
	} {
		got := EncodeMove(data.args.move)

		if got != data.want {
			test.Fail()
		}
	}
}

func TestEncodePiece(test *testing.T) {
	type args struct {
		piece common.Piece
	}
	type data struct {
		args args
		want string
	}

	for _, data := range []data{
		{
			args: args{
				piece: MockPiece{
					kind:  common.King,
					color: common.White,
				},
			},
			want: "K",
		},
		{
			args: args{
				piece: MockPiece{
					kind:  common.Queen,
					color: common.Black,
				},
			},
			want: "q",
		},
		{
			args: args{
				piece: MockPiece{
					kind:  common.Rook,
					color: common.White,
				},
			},
			want: "R",
		},
		{
			args: args{
				piece: MockPiece{
					kind:  common.Bishop,
					color: common.Black,
				},
			},
			want: "b",
		},
		{
			args: args{
				piece: MockPiece{
					kind:  common.Knight,
					color: common.White,
				},
			},
			want: "N",
		},
		{
			args: args{
				piece: MockPiece{
					kind:  common.Pawn,
					color: common.Black,
				},
			},
			want: "p",
		},
	} {
		got := EncodePiece(data.args.piece)

		if got != data.want {
			test.Fail()
		}
	}
}

func TestEncodePieceStorage(test *testing.T) {
	type args struct {
		storage common.PieceStorage
	}
	type data struct {
		args args
		want string
	}

	for _, data := range []data{
		{
			args: args{
				storage: MockPieceStorage{
					size: common.Size{
						Width:  5,
						Height: 5,
					},
					piece: func(position common.Position) (piece common.Piece, ok bool) {
						return nil, false
					},
				},
			},
			want: "5/5/5/5/5",
		},
		{
			args: args{
				storage: MockPieceStorage{
					size: common.Size{
						Width:  5,
						Height: 5,
					},
					piece: func(position common.Position) (piece common.Piece, ok bool) {
						switch position {
						case common.Position{File: 0, Rank: 2}:
							piece = pieces.NewKing(common.White, common.Position{
								File: 0,
								Rank: 2,
							})
						}

						ok = piece != nil
						return piece, ok
					},
				},
			},
			want: "5/5/K4/5/5",
		},
		{
			args: args{
				storage: MockPieceStorage{
					size: common.Size{
						Width:  5,
						Height: 5,
					},
					piece: func(position common.Position) (piece common.Piece, ok bool) {
						switch position {
						case common.Position{File: 1, Rank: 2}:
							piece = pieces.NewKing(common.White, common.Position{
								File: 1,
								Rank: 2,
							})
						}

						ok = piece != nil
						return piece, ok
					},
				},
			},
			want: "5/5/1K3/5/5",
		},
		{
			args: args{
				storage: MockPieceStorage{
					size: common.Size{
						Width:  5,
						Height: 5,
					},
					piece: func(position common.Position) (piece common.Piece, ok bool) {
						switch position {
						case common.Position{File: 1, Rank: 2}:
							piece = pieces.NewKing(common.White, common.Position{
								File: 1,
								Rank: 2,
							})
						case common.Position{File: 2, Rank: 2}:
							piece = pieces.NewQueen(common.Black, common.Position{
								File: 2,
								Rank: 2,
							})
						}

						ok = piece != nil
						return piece, ok
					},
				},
			},
			want: "5/5/1Kq2/5/5",
		},
		{
			args: args{
				storage: MockPieceStorage{
					size: common.Size{
						Width:  5,
						Height: 5,
					},
					piece: func(position common.Position) (piece common.Piece, ok bool) {
						switch position {
						case common.Position{File: 1, Rank: 2}:
							piece = pieces.NewKing(common.White, common.Position{
								File: 1,
								Rank: 2,
							})
						case common.Position{File: 4, Rank: 2}:
							piece = pieces.NewQueen(common.Black, common.Position{
								File: 4,
								Rank: 2,
							})
						}

						ok = piece != nil
						return piece, ok
					},
				},
			},
			want: "5/5/1K2q/5/5",
		},
		{
			args: args{
				storage: MockPieceStorage{
					size: common.Size{
						Width:  5,
						Height: 5,
					},
					piece: func(position common.Position) (piece common.Piece, ok bool) {
						switch position {
						case common.Position{File: 0, Rank: 3}:
							piece = pieces.NewKing(common.White, common.Position{
								File: 0,
								Rank: 3,
							})
						case common.Position{File: 1, Rank: 2}:
							piece = pieces.NewQueen(common.Black, common.Position{
								File: 1,
								Rank: 2,
							})
						case common.Position{File: 2, Rank: 2}:
							piece = pieces.NewQueen(common.White, common.Position{
								File: 2,
								Rank: 2,
							})
						case common.Position{File: 1, Rank: 1}:
							piece = pieces.NewRook(common.Black, common.Position{
								File: 1,
								Rank: 1,
							})
						case common.Position{File: 4, Rank: 1}:
							piece = pieces.NewRook(common.White, common.Position{
								File: 4,
								Rank: 1,
							})
						}

						ok = piece != nil
						return piece, ok
					},
				},
			},
			want: "5/K4/1qQ2/1r2R/5",
		},
		// specific test for the bug with calculation of the last file
		{
			args: args{
				storage: MockPieceStorage{
					size: common.Size{
						Width:  5,
						Height: 3,
					},
					piece: func(position common.Position) (piece common.Piece, ok bool) {
						switch position {
						case common.Position{File: 0, Rank: 2}:
							piece = pieces.NewKing(common.White, common.Position{
								File: 0,
								Rank: 2,
							})
						case common.Position{File: 1, Rank: 1}:
							piece = pieces.NewQueen(common.Black, common.Position{
								File: 1,
								Rank: 1,
							})
						case common.Position{File: 2, Rank: 1}:
							piece = pieces.NewQueen(common.White, common.Position{
								File: 2,
								Rank: 1,
							})
						case common.Position{File: 1, Rank: 0}:
							piece = pieces.NewRook(common.Black, common.Position{
								File: 1,
								Rank: 0,
							})
						case common.Position{File: 4, Rank: 0}:
							piece = pieces.NewRook(common.White, common.Position{
								File: 4,
								Rank: 0,
							})
						}

						ok = piece != nil
						return piece, ok
					},
				},
			},
			want: "K4/1qQ2/1r2R",
		},
	} {
		got := EncodePieceStorage(data.args.storage)

		if got != data.want {
			test.Fail()
		}
	}
}
