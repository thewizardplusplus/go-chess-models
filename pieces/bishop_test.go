package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
)

func TestNewBishop(test *testing.T) {
	piece := NewBishop(common.White, common.Position{
		File: 2,
		Rank: 3,
	})

	expectedPiece := Bishop{
		Base: Base{
			kind:  common.Bishop,
			color: common.White,
			position: common.Position{
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
	piece := NewBishop(common.White, common.Position{
		File: 2,
		Rank: 3,
	})
	nextPiece := piece.ApplyPosition(common.Position{
		File: 4,
		Rank: 2,
	})

	expectedPiece := Bishop{
		Base: Base{
			kind:  common.Bishop,
			color: common.White,
			position: common.Position{
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
			kind:  common.Bishop,
			color: common.White,
			position: common.Position{
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
		position   common.Position
	}
	type data struct {
		args      args
		wantMoves []common.Move
		wantErr   error
	}

	for _, data := range []data{
		{
			args: args{
				boardInFEN: "5/5/2B2/5/5",
				position: common.Position{
					File: 2,
					Rank: 2,
				},
			},
			wantMoves: []common.Move{
				{
					Start: common.Position{
						File: 2,
						Rank: 2,
					},
					Finish: common.Position{
						File: 0,
						Rank: 0,
					},
				},
				{
					Start: common.Position{
						File: 2,
						Rank: 2,
					},
					Finish: common.Position{
						File: 4,
						Rank: 0,
					},
				},
				{
					Start: common.Position{
						File: 2,
						Rank: 2,
					},
					Finish: common.Position{
						File: 1,
						Rank: 1,
					},
				},
				{
					Start: common.Position{
						File: 2,
						Rank: 2,
					},
					Finish: common.Position{
						File: 3,
						Rank: 1,
					},
				},
				{
					Start: common.Position{
						File: 2,
						Rank: 2,
					},
					Finish: common.Position{
						File: 1,
						Rank: 3,
					},
				},
				{
					Start: common.Position{
						File: 2,
						Rank: 2,
					},
					Finish: common.Position{
						File: 3,
						Rank: 3,
					},
				},
				{
					Start: common.Position{
						File: 2,
						Rank: 2,
					},
					Finish: common.Position{
						File: 0,
						Rank: 4,
					},
				},
				{
					Start: common.Position{
						File: 2,
						Rank: 2,
					},
					Finish: common.Position{
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
				position: common.Position{
					File: 2,
					Rank: 2,
				},
			},
			wantMoves: []common.Move{
				{
					Start: common.Position{
						File: 2,
						Rank: 2,
					},
					Finish: common.Position{
						File: 1,
						Rank: 1,
					},
				},
				{
					Start: common.Position{
						File: 2,
						Rank: 2,
					},
					Finish: common.Position{
						File: 3,
						Rank: 1,
					},
				},
				{
					Start: common.Position{
						File: 2,
						Rank: 2,
					},
					Finish: common.Position{
						File: 1,
						Rank: 3,
					},
				},
				{
					Start: common.Position{
						File: 2,
						Rank: 2,
					},
					Finish: common.Position{
						File: 3,
						Rank: 3,
					},
				},
				{
					Start: common.Position{
						File: 2,
						Rank: 2,
					},
					Finish: common.Position{
						File: 0,
						Rank: 4,
					},
				},
				{
					Start: common.Position{
						File: 2,
						Rank: 2,
					},
					Finish: common.Position{
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
				position: common.Position{
					File: 1,
					Rank: 3,
				},
			},
			wantMoves: []common.Move{
				{
					Start: common.Position{
						File: 1,
						Rank: 3,
					},
					Finish: common.Position{
						File: 3,
						Rank: 1,
					},
				},
				{
					Start: common.Position{
						File: 1,
						Rank: 3,
					},
					Finish: common.Position{
						File: 0,
						Rank: 2,
					},
				},
				{
					Start: common.Position{
						File: 1,
						Rank: 3,
					},
					Finish: common.Position{
						File: 2,
						Rank: 2,
					},
				},
				{
					Start: common.Position{
						File: 1,
						Rank: 3,
					},
					Finish: common.Position{
						File: 0,
						Rank: 4,
					},
				},
				{
					Start: common.Position{
						File: 1,
						Rank: 3,
					},
					Finish: common.Position{
						File: 2,
						Rank: 4,
					},
				},
			},
			wantErr: nil,
		},
	} {
		storage, err :=
			uci.DecodePieceStorage(data.args.boardInFEN, NewPiece, boards.NewMapBoard)
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
