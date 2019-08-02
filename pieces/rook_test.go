package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/uci"
)

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
	type args struct {
		boardInFEN string
		position   models.Position
	}
	type data struct {
		args      args
		wantMoves []models.Move
		wantErr   error
	}

	for _, data := range []data{
		data{
			args: args{
				boardInFEN: "5/5/2R2/5/5",
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
			args: args{
				boardInFEN: "5/5/1pR2/2p2/5",
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
		storage, err := uci.DecodePieceStorage(
			data.args.boardInFEN,
			NewPiece,
			models.NewBoard,
		)
		if err != nil {
			test.Fail()
			continue
		}

		generator := models.MoveGenerator{}
		gotMoves, gotErr :=
			generator.MovesForPosition(
				storage,
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
