package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestBishopInterface(test *testing.T) {
	bishopType := reflect.TypeOf(Bishop{})
	pieceType := reflect.
		TypeOf((*models.Piece)(nil)).
		Elem()
	if !bishopType.Implements(pieceType) {
		test.Fail()
	}
}

func TestNewBishop(test *testing.T) {
	piece := NewBishop(
		models.White,
		models.Position{File: 2, Rank: 3},
	)

	expectedPiece := Bishop{
		Base: Base{
			kind:  models.Bishop,
			color: models.White,
			position: models.Position{
				File: 2,
				Rank: 3,
			},
		},
	}
	if !reflect.DeepEqual(
		piece,
		expectedPiece,
	) {
		test.Fail()
	}
}

func TestBishopApplyPosition(
	test *testing.T,
) {
	piece := NewBishop(
		models.White,
		models.Position{File: 2, Rank: 3},
	)
	nextPiece := piece.ApplyPosition(
		models.Position{
			File: 4,
			Rank: 2,
		},
	)

	expectedPiece := Bishop{
		Base: Base{
			kind:  models.Bishop,
			color: models.White,
			position: models.Position{
				File: 2,
				Rank: 3,
			},
		},
	}
	if !reflect.DeepEqual(
		piece,
		expectedPiece,
	) {
		test.Fail()
	}

	expectedNextPiece := Bishop{
		Base: Base{
			kind:  models.Bishop,
			color: models.White,
			position: models.Position{
				File: 4,
				Rank: 2,
			},
		},
	}
	if !reflect.DeepEqual(
		nextPiece,
		expectedNextPiece,
	) {
		test.Fail()
	}
}

func TestBishopCheckMove(test *testing.T) {
	type fields struct {
		size   models.Size
		pieces []models.Piece
	}
	type args struct {
		position models.Position
	}
	type data struct {
		fields fields
		args   args
		want   []models.Move
	}

	for _, data := range []data{
		data{
			fields: fields{
				size: models.Size{5, 5},
				pieces: []models.Piece{
					NewBishop(
						models.White,
						models.Position{
							File: 2,
							Rank: 2,
						},
					),
				},
			},
			args: args{
				position: models.Position{
					File: 2,
					Rank: 2,
				},
			},
			want: []models.Move{
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 0,
						Rank: 0,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 4,
						Rank: 0,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 1,
						Rank: 1,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 1,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 1,
						Rank: 3,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 3,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 0,
						Rank: 4,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 4,
						Rank: 4,
					},
				},
			},
		},
		data{
			fields: fields{
				size: models.Size{5, 5},
				pieces: []models.Piece{
					NewBishop(
						models.White,
						models.Position{
							File: 2,
							Rank: 2,
						},
					),
					NewKing(
						models.Black,
						models.Position{
							File: 1,
							Rank: 1,
						},
					),
					NewKing(
						models.Black,
						models.Position{
							File: 3,
							Rank: 1,
						},
					),
				},
			},
			args: args{
				position: models.Position{
					File: 2,
					Rank: 2,
				},
			},
			want: []models.Move{
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 1,
						Rank: 1,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 1,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 1,
						Rank: 3,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 3,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 0,
						Rank: 4,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 4,
						Rank: 4,
					},
				},
			},
		},
	} {
		board := models.NewBoard(
			data.fields.size,
			data.fields.pieces,
		)
		generator := models.MoveGenerator{}
		got := generator.MovesForPosition(
			board,
			data.args.position,
		)

		if !reflect.DeepEqual(
			got,
			data.want,
		) {
			test.Fail()
		}
	}
}
