package boards

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

func TestBitBoardPieceGroupByColorAndKindToBigInt(test *testing.T) {
	piecesByColorAndKind := new(bitBoardPieceGroupByColorAndKind)
	bigInt := piecesByColorAndKind.ToBigInt()

	expectedBigInt := big.NewInt(0)
	if !reflect.DeepEqual(bigInt, expectedBigInt) {
		test.Fail()
	}
}

func TestBitBoardPieceGroupByColorAndKindIsPositionOccupied(test *testing.T) {
	type args struct {
		size     common.Size
		position common.Position
	}
	type data struct {
		piecesByColorAndKind *bitBoardPieceGroupByColorAndKind
		args                 args
		want                 bool
	}

	for _, data := range []data{
		{
			piecesByColorAndKind: func() *bitBoardPieceGroupByColorAndKind {
				piecesByColorAndKind := new(bitBoardPieceGroupByColorAndKind)
				piecesByColorAndKind.SetPositionStatus(
					common.Size{5, 5},
					common.Position{4, 2},
					occupiedPositionStatus,
				)

				return piecesByColorAndKind
			}(),
			args: args{
				size:     common.Size{5, 5},
				position: common.Position{2, 3},
			},
			want: false,
		},
		{
			piecesByColorAndKind: func() *bitBoardPieceGroupByColorAndKind {
				piecesByColorAndKind := new(bitBoardPieceGroupByColorAndKind)
				piecesByColorAndKind.SetPositionStatus(
					common.Size{5, 5},
					common.Position{4, 2},
					occupiedPositionStatus,
				)

				return piecesByColorAndKind
			}(),
			args: args{
				size:     common.Size{5, 5},
				position: common.Position{4, 2},
			},
			want: true,
		},
	} {
		got := data.piecesByColorAndKind.
			IsPositionOccupied(data.args.size, data.args.position)

		if got != data.want {
			test.Fail()
		}
	}
}

func TestBitBoardPieceGroupByColorAndKindSetValue(test *testing.T) {
	expectedPiecesByColorAndKind := new(bitBoardPieceGroupByColorAndKind)
	expectedPiecesByColorAndKind.SetPositionStatus(
		common.Size{5, 5},
		common.Position{2, 3},
		occupiedPositionStatus,
	)

	piecesByColorAndKind := new(bitBoardPieceGroupByColorAndKind)
	piecesByColorAndKind.SetValue(expectedPiecesByColorAndKind)

	if !reflect.DeepEqual(piecesByColorAndKind, expectedPiecesByColorAndKind) {
		test.Fail()
	}
}

func TestBitBoardPieceGroupByColorAndKindSetPositionStatus(test *testing.T) {
	type args struct {
		size           common.Size
		position       common.Position
		positionStatus positionStatus
	}
	type data struct {
		piecesByColorAndKind     *bitBoardPieceGroupByColorAndKind
		args                     args
		wantPiecesByColorAndKind *bitBoardPieceGroupByColorAndKind
	}

	for _, data := range []data{
		{
			piecesByColorAndKind: func() *bitBoardPieceGroupByColorAndKind {
				piecesByColorAndKind := new(bitBoardPieceGroupByColorAndKind)
				piecesByColorAndKind.ToBigInt().
					SetBit(piecesByColorAndKind.ToBigInt(), 5*3+2, 1)

				return piecesByColorAndKind
			}(),
			args: args{
				size:           common.Size{5, 5},
				position:       common.Position{2, 3},
				positionStatus: freePositionStatus,
			},
			wantPiecesByColorAndKind: func() *bitBoardPieceGroupByColorAndKind {
				piecesByColorAndKind := new(bitBoardPieceGroupByColorAndKind)
				piecesByColorAndKind.ToBigInt().SetBits([]big.Word{})

				return piecesByColorAndKind
			}(),
		},
		{
			piecesByColorAndKind: new(bitBoardPieceGroupByColorAndKind),
			args: args{
				size:           common.Size{5, 5},
				position:       common.Position{2, 3},
				positionStatus: occupiedPositionStatus,
			},
			wantPiecesByColorAndKind: func() *bitBoardPieceGroupByColorAndKind {
				piecesByColorAndKind := new(bitBoardPieceGroupByColorAndKind)
				piecesByColorAndKind.ToBigInt().
					SetBit(piecesByColorAndKind.ToBigInt(), 5*3+2, 1)

				return piecesByColorAndKind
			}(),
		},
	} {
		data.piecesByColorAndKind.SetPositionStatus(
			data.args.size,
			data.args.position,
			data.args.positionStatus,
		)

		if !reflect.DeepEqual(
			data.piecesByColorAndKind,
			data.wantPiecesByColorAndKind,
		) {
			test.Fail()
		}
	}
}
