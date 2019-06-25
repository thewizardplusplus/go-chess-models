package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestRookInterface(test *testing.T) {
	rookType := reflect.TypeOf(Rook{})
	pieceType := reflect.
		TypeOf((*models.Piece)(nil)).
		Elem()
	if !rookType.Implements(pieceType) {
		test.Fail()
	}
}

func TestRookFactoryType(test *testing.T) {
	gotType := reflect.TypeOf(NewRook)
	wantType := reflect.
		TypeOf((*Factory)(nil)).
		Elem()
	if !gotType.AssignableTo(wantType) {
		test.Fail()
	}
}

func TestNewRook(test *testing.T) {
	piece := NewRook(
		models.White,
		models.Position{File: 2, Rank: 3},
	)

	expectedPiece := Rook{
		Base: Base{
			kind:  models.Rook,
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

func TestRookApplyPosition(
	test *testing.T,
) {
	piece := NewRook(
		models.White,
		models.Position{File: 2, Rank: 3},
	)
	nextPiece := piece.ApplyPosition(
		models.Position{
			File: 4,
			Rank: 2,
		},
	)

	expectedPiece := Rook{
		Base: Base{
			kind:  models.Rook,
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

	expectedNextPiece := Rook{
		Base: Base{
			kind:  models.Rook,
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

func TestRookCheckMove(test *testing.T) {
	type fields struct {
		size   models.Size
		pieces []models.Piece
	}
	type args struct {
		position models.Position
	}
	type data struct {
		fields    fields
		args      args
		wantMoves []models.Move
		wantErr   error
	}

	for _, data := range []data{
		data{
			fields: fields{
				size: models.Size{5, 5},
				pieces: []models.Piece{
					NewRook(
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
			wantMoves: []models.Move{
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 2,
						Rank: 0,
					},
				},
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
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 0,
						Rank: 2,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 1,
						Rank: 2,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 2,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 4,
						Rank: 2,
					},
				},
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
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 2,
						Rank: 4,
					},
				},
			},
			wantErr: nil,
		},
		data{
			fields: fields{
				size: models.Size{5, 5},
				pieces: []models.Piece{
					NewRook(
						models.White,
						models.Position{
							File: 2,
							Rank: 2,
						},
					),
					NewPawn(
						models.Black,
						models.Position{
							File: 2,
							Rank: 1,
						},
					),
					NewPawn(
						models.Black,
						models.Position{
							File: 1,
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
			wantMoves: []models.Move{
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
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 1,
						Rank: 2,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 2,
					},
				},
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 4,
						Rank: 2,
					},
				},
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
				models.Move{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 2,
						Rank: 4,
					},
				},
			},
			wantErr: nil,
		},
	} {
		board := models.NewBoard(
			data.fields.size,
			data.fields.pieces,
		)
		generator := models.MoveGenerator{}
		gotMoves, gotErr :=
			generator.MovesForPosition(
				board,
				data.args.position,
			)

		if !reflect.DeepEqual(
			gotMoves,
			data.wantMoves,
		) {
			test.Fail()
		}
		if !reflect.DeepEqual(
			gotErr,
			data.wantErr,
		) {
			test.Fail()
		}
	}
}
