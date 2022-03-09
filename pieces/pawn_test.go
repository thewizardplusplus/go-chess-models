package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
)

func TestNewPawn(test *testing.T) {
	piece := NewPawn(common.White, models.Position{
		File: 2,
		Rank: 3,
	})

	expectedPiece := Pawn{
		Base: Base{
			kind:  common.Pawn,
			color: common.White,
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

func TestPawnApplyPosition(test *testing.T) {
	piece := NewPawn(common.White, models.Position{
		File: 2,
		Rank: 3,
	})
	nextPiece := piece.ApplyPosition(models.Position{
		File: 4,
		Rank: 2,
	})

	expectedPiece := Pawn{
		Base: Base{
			kind:  common.Pawn,
			color: common.White,
			position: models.Position{
				File: 2,
				Rank: 3,
			},
		},
	}
	if !reflect.DeepEqual(piece, expectedPiece) {
		test.Fail()
	}

	expectedNextPiece := Pawn{
		Base: Base{
			kind:  common.Pawn,
			color: common.White,
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

func TestPawnCheckMove(test *testing.T) {
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
				boardInFEN: "5/5/2p2/5/5",
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
						File: 2,
						Rank: 1,
					},
				},
			},
			wantErr: nil,
		},
		{
			args: args{
				boardInFEN: "5/5/2p2/1PPP1/5",
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
			},
			wantErr: nil,
		},
		{
			args: args{
				boardInFEN: "5/5/2P2/5/5",
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
						File: 2,
						Rank: 3,
					},
				},
			},
			wantErr: nil,
		},
		{
			args: args{
				boardInFEN: "5/1ppp1/2P2/5/5",
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
			},
			wantErr: nil,
		},
	} {
		storage, err :=
			uci.DecodePieceStorage(data.args.boardInFEN, NewPiece, models.NewBoard)
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
