package uci

import (
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

type MockPiece struct {
	kind  models.Kind
	color models.Color
}

func (piece MockPiece) Kind() models.Kind {
	return piece.kind
}

func (piece MockPiece) Color() models.Color {
	return piece.color
}

func (piece MockPiece) Position() models.Position {
	panic("not implemented")
}

func (piece MockPiece) ApplyPosition(position models.Position) models.Piece {
	panic("not implemented")
}

func (piece MockPiece) CheckMove(
	move models.Move,
	storage models.PieceStorage,
) bool {
	panic("not implemented")
}

type MockPieceStorage struct {
	size models.Size

	piece func(position models.Position) (piece models.Piece, ok bool)
}

func (storage MockPieceStorage) Size() models.Size {
	return storage.size
}

func (storage MockPieceStorage) Piece(
	position models.Position,
) (piece models.Piece, ok bool) {
	if storage.piece == nil {
		panic("not implemented")
	}

	return storage.piece(position)
}

func (storage MockPieceStorage) Pieces() []models.Piece {
	panic("not implemented")
}

func (storage MockPieceStorage) ApplyMove(
	move models.Move,
) models.PieceStorage {
	panic("not implemented")
}

func (storage MockPieceStorage) CheckMove(move models.Move) error {
	panic("not implemented")
}

func TestEncodePosition(test *testing.T) {
	type args struct {
		position models.Position
	}
	type data struct {
		args args
		want string
	}

	for _, data := range []data{
		{
			args: args{
				position: models.Position{
					File: 2,
					Rank: 1,
				},
			},
			want: "c2",
		},
		{
			args: args{
				position: models.Position{
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
		move models.Move
	}
	type data struct {
		args args
		want string
	}

	for _, data := range []data{
		{
			args: args{
				move: models.Move{
					Start: models.Position{
						File: 2,
						Rank: 1,
					},
					Finish: models.Position{
						File: 2,
						Rank: 3,
					},
				},
			},
			want: "c2c4",
		},
		{
			args: args{
				move: models.Move{
					Start: models.Position{
						File: 5,
						Rank: 6,
					},
					Finish: models.Position{
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
		piece models.Piece
	}
	type data struct {
		args args
		want string
	}

	for _, data := range []data{
		{
			args: args{
				piece: MockPiece{
					kind:  models.King,
					color: models.White,
				},
			},
			want: "K",
		},
		{
			args: args{
				piece: MockPiece{
					kind:  models.Queen,
					color: models.Black,
				},
			},
			want: "q",
		},
		{
			args: args{
				piece: MockPiece{
					kind:  models.Rook,
					color: models.White,
				},
			},
			want: "R",
		},
		{
			args: args{
				piece: MockPiece{
					kind:  models.Bishop,
					color: models.Black,
				},
			},
			want: "b",
		},
		{
			args: args{
				piece: MockPiece{
					kind:  models.Knight,
					color: models.White,
				},
			},
			want: "N",
		},
		{
			args: args{
				piece: MockPiece{
					kind:  models.Pawn,
					color: models.Black,
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
		storage models.PieceStorage
	}
	type data struct {
		args args
		want string
	}

	for _, data := range []data{
		{
			args: args{
				storage: MockPieceStorage{
					size: models.Size{
						Width:  5,
						Height: 5,
					},
					piece: func(position models.Position) (piece models.Piece, ok bool) {
						return nil, false
					},
				},
			},
			want: "5/5/5/5/5",
		},
		{
			args: args{
				storage: MockPieceStorage{
					size: models.Size{
						Width:  5,
						Height: 5,
					},
					piece: func(position models.Position) (piece models.Piece, ok bool) {
						switch position {
						case models.Position{File: 0, Rank: 2}:
							piece = pieces.NewKing(models.White, models.Position{
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
					size: models.Size{
						Width:  5,
						Height: 5,
					},
					piece: func(position models.Position) (piece models.Piece, ok bool) {
						switch position {
						case models.Position{File: 1, Rank: 2}:
							piece = pieces.NewKing(models.White, models.Position{
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
					size: models.Size{
						Width:  5,
						Height: 5,
					},
					piece: func(position models.Position) (piece models.Piece, ok bool) {
						switch position {
						case models.Position{File: 1, Rank: 2}:
							piece = pieces.NewKing(models.White, models.Position{
								File: 1,
								Rank: 2,
							})
						case models.Position{File: 2, Rank: 2}:
							piece = pieces.NewQueen(models.Black, models.Position{
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
					size: models.Size{
						Width:  5,
						Height: 5,
					},
					piece: func(position models.Position) (piece models.Piece, ok bool) {
						switch position {
						case models.Position{File: 1, Rank: 2}:
							piece = pieces.NewKing(models.White, models.Position{
								File: 1,
								Rank: 2,
							})
						case models.Position{File: 4, Rank: 2}:
							piece = pieces.NewQueen(models.Black, models.Position{
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
					size: models.Size{
						Width:  5,
						Height: 5,
					},
					piece: func(position models.Position) (piece models.Piece, ok bool) {
						switch position {
						case models.Position{File: 0, Rank: 3}:
							piece = pieces.NewKing(models.White, models.Position{
								File: 0,
								Rank: 3,
							})
						case models.Position{File: 1, Rank: 2}:
							piece = pieces.NewQueen(models.Black, models.Position{
								File: 1,
								Rank: 2,
							})
						case models.Position{File: 2, Rank: 2}:
							piece = pieces.NewQueen(models.White, models.Position{
								File: 2,
								Rank: 2,
							})
						case models.Position{File: 1, Rank: 1}:
							piece = pieces.NewRook(models.Black, models.Position{
								File: 1,
								Rank: 1,
							})
						case models.Position{File: 4, Rank: 1}:
							piece = pieces.NewRook(models.White, models.Position{
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
	} {
		got := EncodePieceStorage(data.args.storage)

		if got != data.want {
			test.Fail()
		}
	}
}
