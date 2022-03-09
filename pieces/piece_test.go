package pieces

import (
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/common"
)

func TestNewPiece(test *testing.T) {
	type args struct {
		kind     common.Kind
		color    common.Color
		position models.Position
	}
	type data struct {
		args args
		want models.Piece
	}

	for _, data := range []data{
		{
			args: args{
				kind:  common.King,
				color: common.White,
				position: models.Position{
					File: 2,
					Rank: 3,
				},
			},
			want: NewKing(common.White, models.Position{
				File: 2,
				Rank: 3,
			}),
		},
		{
			args: args{
				kind:  common.Queen,
				color: common.Black,
				position: models.Position{
					File: 4,
					Rank: 2,
				},
			},
			want: NewQueen(common.Black, models.Position{
				File: 4,
				Rank: 2,
			}),
		},
		{
			args: args{
				kind:  common.Rook,
				color: common.White,
				position: models.Position{
					File: 2,
					Rank: 3,
				},
			},
			want: NewRook(common.White, models.Position{
				File: 2,
				Rank: 3,
			}),
		},
		{
			args: args{
				kind:  common.Bishop,
				color: common.Black,
				position: models.Position{
					File: 4,
					Rank: 2,
				},
			},
			want: NewBishop(common.Black, models.Position{
				File: 4,
				Rank: 2,
			}),
		},
		{
			args: args{
				kind:  common.Knight,
				color: common.White,
				position: models.Position{
					File: 2,
					Rank: 3,
				},
			},
			want: NewKnight(common.White, models.Position{
				File: 2,
				Rank: 3,
			}),
		},
		{
			args: args{
				kind:  common.Pawn,
				color: common.Black,
				position: models.Position{
					File: 4,
					Rank: 2,
				},
			},
			want: NewPawn(common.Black, models.Position{
				File: 4,
				Rank: 2,
			}),
		},
	} {
		got := NewPiece(data.args.kind, data.args.color, data.args.position)

		if !reflect.DeepEqual(got, data.want) {
			test.Fail()
		}
	}
}
