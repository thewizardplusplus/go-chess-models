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

func (
	piece MockPiece,
) Position() models.Position {
	panic("not implemented")
}

func (piece MockPiece) ApplyPosition(
	position models.Position,
) models.Piece {
	panic("not implemented")
}

func (piece MockPiece) CheckMove(
	move models.Move,
	storage models.PieceStorage,
) bool {
	panic("not implemented")
}

func (piece MockPiece) String() string {
	panic("not implemented")
}

type MockPieceStorage struct {
	size models.Size

	piece func(
		position models.Position,
	) (piece models.Piece, ok bool)
}

func (
	storage MockPieceStorage,
) Size() models.Size {
	return storage.size
}

func (
	storage MockPieceStorage,
) Piece(
	position models.Position,
) (piece models.Piece, ok bool) {
	if storage.piece == nil {
		panic("not implemented")
	}

	return storage.piece(position)
}

func (
	storage MockPieceStorage,
) Pieces() []models.Piece {
	panic("not implemented")
}

func (storage MockPieceStorage) ApplyMove(
	move models.Move,
) models.PieceStorage {
	panic("not implemented")
}

func (storage MockPieceStorage) CheckMove(
	move models.Move,
) error {
	panic("not implemented")
}

func (
	storage MockPieceStorage,
) String() string {
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
		data{
			args: args{
				position: models.Position{2, 1},
			},
			want: "c2",
		},
		data{
			args: args{
				position: models.Position{5, 6},
			},
			want: "f7",
		},
	} {
		got :=
			EncodePosition(data.args.position)

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
		data{
			args: args{
				piece: MockPiece{
					kind:  models.King,
					color: models.White,
				},
			},
			want: "K",
		},
		data{
			args: args{
				piece: MockPiece{
					kind:  models.Queen,
					color: models.Black,
				},
			},
			want: "q",
		},
	} {
		got := EncodePiece(data.args.piece)

		if got != data.want {
			test.Fail()
		}
	}
}

func TestEncodePieceStorage(
	test *testing.T,
) {
	type args struct {
		storage models.PieceStorage
	}
	type data struct {
		args args
		want string
	}

	for _, data := range []data{
		data{
			args: args{
				storage: MockPieceStorage{
					size: models.Size{5, 5},
					piece: func(
						position models.Position,
					) (piece models.Piece, ok bool) {
						return nil, false
					},
				},
			},
			want: "5/5/5/5/5",
		},
		data{
			args: args{
				storage: MockPieceStorage{
					size: models.Size{5, 5},
					piece: func(
						position models.Position,
					) (piece models.Piece, ok bool) {
						switch position {
						case models.Position{0, 2}:
							piece = pieces.NewKing(
								models.White,
								models.Position{0, 2},
							)
						}

						ok = piece != nil
						return piece, ok
					},
				},
			},
			want: "5/5/K4/5/5",
		},
		data{
			args: args{
				storage: MockPieceStorage{
					size: models.Size{5, 5},
					piece: func(
						position models.Position,
					) (piece models.Piece, ok bool) {
						switch position {
						case models.Position{1, 2}:
							piece = pieces.NewKing(
								models.White,
								models.Position{1, 2},
							)
						}

						ok = piece != nil
						return piece, ok
					},
				},
			},
			want: "5/5/1K3/5/5",
		},
		data{
			args: args{
				storage: MockPieceStorage{
					size: models.Size{5, 5},
					piece: func(
						position models.Position,
					) (piece models.Piece, ok bool) {
						switch position {
						case models.Position{1, 2}:
							piece = pieces.NewKing(
								models.White,
								models.Position{1, 2},
							)
						case models.Position{2, 2}:
							piece = pieces.NewQueen(
								models.Black,
								models.Position{2, 2},
							)
						}

						ok = piece != nil
						return piece, ok
					},
				},
			},
			want: "5/5/1Kq2/5/5",
		},
		data{
			args: args{
				storage: MockPieceStorage{
					size: models.Size{5, 5},
					piece: func(
						position models.Position,
					) (piece models.Piece, ok bool) {
						switch position {
						case models.Position{1, 2}:
							piece = pieces.NewKing(
								models.White,
								models.Position{1, 2},
							)
						case models.Position{4, 2}:
							piece = pieces.NewQueen(
								models.Black,
								models.Position{4, 2},
							)
						}

						ok = piece != nil
						return piece, ok
					},
				},
			},
			want: "5/5/1K2q/5/5",
		},
		data{
			args: args{
				storage: MockPieceStorage{
					size: models.Size{5, 5},
					piece: func(
						position models.Position,
					) (piece models.Piece, ok bool) {
						switch position {
						case models.Position{0, 3}:
							piece = pieces.NewKing(
								models.White,
								models.Position{0, 3},
							)
						case models.Position{1, 2}:
							piece = pieces.NewQueen(
								models.Black,
								models.Position{1, 2},
							)
						case models.Position{2, 2}:
							piece = pieces.NewQueen(
								models.White,
								models.Position{2, 2},
							)
						case models.Position{1, 1}:
							piece = pieces.NewRook(
								models.Black,
								models.Position{1, 1},
							)
						case models.Position{4, 1}:
							piece = pieces.NewRook(
								models.White,
								models.Position{4, 1},
							)
						}

						ok = piece != nil
						return piece, ok
					},
				},
			},
			want: "5/K4/1qQ2/1r2R/5",
		},
	} {
		got := EncodePieceStorage(
			data.args.storage,
		)

		if got != data.want {
			test.Fail()
		}
	}
}
