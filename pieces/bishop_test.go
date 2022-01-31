package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
)

func TestNewBishop(test *testing.T) {
	piece := NewBishop(models.White, models.Position{
		File: 2,
		Rank: 3,
	})

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
	if !reflect.DeepEqual(piece, expectedPiece) {
		test.Fail()
	}
}

func TestBishopApplyPosition(test *testing.T) {
	piece := NewBishop(models.White, models.Position{
		File: 2,
		Rank: 3,
	})
	nextPiece := piece.ApplyPosition(models.Position{
		File: 4,
		Rank: 2,
	})

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
	if !reflect.DeepEqual(piece, expectedPiece) {
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
	if !reflect.DeepEqual(nextPiece, expectedNextPiece) {
		test.Fail()
	}
}

func TestBishopCheckMove(test *testing.T) {
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
				boardInFEN: "5/5/2B2/5/5",
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
						File: 4,
						Rank: 4,
					},
				},
			},
			wantErr: nil,
		},
		{
			args: args{
				boardInFEN: "5/5/2B2/1p1p1/5",
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
						File: 4,
						Rank: 4,
					},
				},
			},
			wantErr: nil,
		},
		// specific test for the bug with a path scanning
		{
			args: args{
				boardInFEN: "5/1B3/5/3p1/5",
				position: models.Position{
					File: 1,
					Rank: 3,
				},
			},
			wantMoves: []models.Move{
				{
					Start: models.Position{
						File: 1,
						Rank: 3,
					},
					Finish: models.Position{
						File: 3,
						Rank: 1,
					},
				},
				{
					Start: models.Position{
						File: 1,
						Rank: 3,
					},
					Finish: models.Position{
						File: 0,
						Rank: 2,
					},
				},
				{
					Start: models.Position{
						File: 1,
						Rank: 3,
					},
					Finish: models.Position{
						File: 2,
						Rank: 2,
					},
				},
				{
					Start: models.Position{
						File: 1,
						Rank: 3,
					},
					Finish: models.Position{
						File: 0,
						Rank: 4,
					},
				},
				{
					Start: models.Position{
						File: 1,
						Rank: 3,
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
