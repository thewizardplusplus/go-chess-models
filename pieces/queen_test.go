package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
)

func TestNewQueen(test *testing.T) {
	piece := NewQueen(models.White, models.Position{
		File: 2,
		Rank: 3,
	})

	expectedPiece := Queen{
		Base: Base{
			kind:  models.Queen,
			color: models.White,
			position: models.Position{
				File: 2,
				Rank: 3,
			},
		},
	}
	if !reflect.DeepEqual(piece, expectedPiece) {
		test.Fail()
	}
}

func TestQueenApplyPosition(test *testing.T) {
	piece := NewQueen(models.White, models.Position{
		File: 2,
		Rank: 3,
	})
	nextPiece := piece.ApplyPosition(models.Position{
		File: 4,
		Rank: 2,
	})

	expectedPiece := Queen{
		Base: Base{
			kind:  models.Queen,
			color: models.White,
			position: models.Position{
				File: 2,
				Rank: 3,
			},
		},
	}
	if !reflect.DeepEqual(piece, expectedPiece) {
		test.Fail()
	}

	expectedNextPiece := Queen{
		Base: Base{
			kind:  models.Queen,
			color: models.White,
			position: models.Position{
				File: 4,
				Rank: 2,
			},
		},
	}
	if !reflect.DeepEqual(nextPiece, expectedNextPiece) {
		test.Fail()
	}
}

func TestQueenCheckMove(test *testing.T) {
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
		{
			args: args{
				boardInFEN: "5/5/2Q2/5/5",
				position: models.Position{
					File: 2,
					Rank: 2,
				},
			},
			wantMoves: []models.Move{
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 0,
						Rank: 0,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 2,
						Rank: 0,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 4,
						Rank: 0,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 1,
						Rank: 1,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 2,
						Rank: 1,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 1,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 0,
						Rank: 2,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 1,
						Rank: 2,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 2,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 4,
						Rank: 2,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 1,
						Rank: 3,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 2,
						Rank: 3,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 3,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 0,
						Rank: 4,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 2,
						Rank: 4,
					},
				},
				{
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
			wantErr: nil,
		},
		{
			args: args{
				boardInFEN: "5/5/1pQ2/1pp2/5",
				position: models.Position{
					File: 2,
					Rank: 2,
				},
			},
			wantMoves: []models.Move{
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 4,
						Rank: 0,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 1,
						Rank: 1,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 2,
						Rank: 1,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 1,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 1,
						Rank: 2,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 2,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 4,
						Rank: 2,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 1,
						Rank: 3,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 2,
						Rank: 3,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 3,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 0,
						Rank: 4,
					},
				},
				{
					Start: models.Position{
						File: 2,
						Rank: 2,
					},
					Finish: models.Position{
						File: 2,
						Rank: 4,
					},
				},
				{
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
			wantErr: nil,
		},
	} {
		storage, err :=
			uci.DecodePieceStorage(data.args.boardInFEN, NewPiece, models.NewMapBoard)
		if err != nil {
			test.Fail()
			continue
		}

		var generator models.MoveGenerator
		gotMoves, gotErr := generator.MovesForPosition(storage, data.args.position)

		if !reflect.DeepEqual(gotMoves, data.wantMoves) {
			test.Fail()
		}
		if !reflect.DeepEqual(gotErr, data.wantErr) {
			test.Fail()
		}
	}
}
