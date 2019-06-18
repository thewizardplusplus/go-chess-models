package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestPawnInterface(test *testing.T) {
	pawnType := reflect.TypeOf(Pawn{})
	pieceType := reflect.
		TypeOf((*models.Piece)(nil)).
		Elem()
	if !pawnType.Implements(pieceType) {
		test.Fail()
	}
}

func TestNewPawn(test *testing.T) {
	piece := NewPawn(
		models.White,
		models.Position{File: 2, Rank: 3},
	)

	expectedPiece := Pawn{
		Base: Base{
			kind:  models.Pawn,
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

func TestPawnApplyPosition(
	test *testing.T,
) {
	piece := NewPawn(
		models.White,
		models.Position{File: 2, Rank: 3},
	)
	nextPiece := piece.ApplyPosition(
		models.Position{
			File: 4,
			Rank: 2,
		},
	)

	expectedPiece := Pawn{
		Base: Base{
			kind:  models.Pawn,
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

	expectedNextPiece := Pawn{
		Base: Base{
			kind:  models.Pawn,
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

func TestPawnCheckMove(test *testing.T) {
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
					NewPawn(
						models.Black,
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
						File: 2,
						Rank: 1,
					},
				},
			},
		},
		data{
			fields: fields{
				size: models.Size{5, 5},
				pieces: []models.Piece{
					NewPawn(
						models.Black,
						models.Position{
							File: 2,
							Rank: 2,
						},
					),
					NewPawn(
						models.White,
						models.Position{
							File: 1,
							Rank: 1,
						},
					),
					NewPawn(
						models.White,
						models.Position{
							File: 2,
							Rank: 1,
						},
					),
					NewPawn(
						models.White,
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
			},
		},
		data{
			fields: fields{
				size: models.Size{5, 5},
				pieces: []models.Piece{
					NewPawn(
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
						File: 2,
						Rank: 3,
					},
				},
			},
		},
		data{
			fields: fields{
				size: models.Size{5, 5},
				pieces: []models.Piece{
					NewPawn(
						models.White,
						models.Position{
							File: 2,
							Rank: 2,
						},
					),
					NewPawn(
						models.Black,
						models.Position{
							File: 1,
							Rank: 3,
						},
					),
					NewPawn(
						models.Black,
						models.Position{
							File: 2,
							Rank: 3,
						},
					),
					NewPawn(
						models.Black,
						models.Position{
							File: 3,
							Rank: 3,
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
