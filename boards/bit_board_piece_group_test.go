package boards

import (
	"errors"
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

func TestBitBoardPieceGroupIteratePiecesByColorAndKind(test *testing.T) {
	type handlerArgs struct {
		color                common.Color
		kind                 common.Kind
		piecesByColorAndKind *bitBoardPieceGroupByColorAndKind
	}
	type args struct {
		makeHandler func(
			handlerArgsGroup *[]handlerArgs,
		) bitBoardPieceGroupByColorAndKindHandler
	}
	type data struct {
		pieceGroup           *bitBoardPieceGroup
		args                 args
		wantHandlerArgsGroup []handlerArgs
		wantErr              error
	}

	for _, data := range []data{
		{
			pieceGroup: func() *bitBoardPieceGroup {
				pieceGroup := new(bitBoardPieceGroup)
				positionIndex := 0
				for colorAsInt := 0; colorAsInt < int(common.ColorCount); colorAsInt++ {
					for kindAsInt := 0; kindAsInt < int(common.KindCount); kindAsInt++ {
						piecesByColorAndKind := &pieceGroup[colorAsInt][kindAsInt]
						piecesByColorAndKind.ToBigInt().
							SetBit(piecesByColorAndKind.ToBigInt(), positionIndex, 1)

						positionIndex++
					}
				}

				return pieceGroup
			}(),
			args: args{
				makeHandler: func(
					handlerArgsGroup *[]handlerArgs,
				) bitBoardPieceGroupByColorAndKindHandler {
					return func(
						color common.Color,
						kind common.Kind,
						piecesByColorAndKind *bitBoardPieceGroupByColorAndKind,
					) error {
						handlerArgs := handlerArgs{color, kind, piecesByColorAndKind}
						*handlerArgsGroup = append(*handlerArgsGroup, handlerArgs)

						return nil
					}
				},
			},
			wantHandlerArgsGroup: func() []handlerArgs {
				var handlerArgsGroup []handlerArgs
				var positionIndex int
				for colorAsInt := 0; colorAsInt < int(common.ColorCount); colorAsInt++ {
					for kindAsInt := 0; kindAsInt < int(common.KindCount); kindAsInt++ {
						color, kind := common.Color(colorAsInt), common.Kind(kindAsInt)

						piecesByColorAndKind := new(bitBoardPieceGroupByColorAndKind)
						piecesByColorAndKind.ToBigInt().
							SetBit(piecesByColorAndKind.ToBigInt(), positionIndex, 1)

						handlerArgs := handlerArgs{color, kind, piecesByColorAndKind}
						handlerArgsGroup = append(handlerArgsGroup, handlerArgs)

						positionIndex++
					}
				}

				return handlerArgsGroup
			}(),
			wantErr: nil,
		},
		{
			pieceGroup: func() *bitBoardPieceGroup {
				pieceGroup := new(bitBoardPieceGroup)
				positionIndex := 0
				for colorAsInt := 0; colorAsInt < int(common.ColorCount); colorAsInt++ {
					for kindAsInt := 0; kindAsInt < int(common.KindCount); kindAsInt++ {
						piecesByColorAndKind := &pieceGroup[colorAsInt][kindAsInt]
						piecesByColorAndKind.ToBigInt().
							SetBit(piecesByColorAndKind.ToBigInt(), positionIndex, 1)

						positionIndex++
					}
				}

				return pieceGroup
			}(),
			args: args{
				makeHandler: func(
					handlerArgsGroup *[]handlerArgs,
				) bitBoardPieceGroupByColorAndKindHandler {
					return func(
						color common.Color,
						kind common.Kind,
						piecesByColorAndKind *bitBoardPieceGroupByColorAndKind,
					) error {
						if color == common.White {
							return errors.New("dummy")
						}

						handlerArgs := handlerArgs{color, kind, piecesByColorAndKind}
						*handlerArgsGroup = append(*handlerArgsGroup, handlerArgs)

						return nil
					}
				},
			},
			wantHandlerArgsGroup: func() []handlerArgs {
				var handlerArgsGroup []handlerArgs
				var positionIndex int
				for kindAsInt := 0; kindAsInt < int(common.KindCount); kindAsInt++ {
					color, kind := common.Black, common.Kind(kindAsInt)

					piecesByColorAndKind := new(bitBoardPieceGroupByColorAndKind)
					piecesByColorAndKind.ToBigInt().
						SetBit(piecesByColorAndKind.ToBigInt(), positionIndex, 1)

					handlerArgs := handlerArgs{color, kind, piecesByColorAndKind}
					handlerArgsGroup = append(handlerArgsGroup, handlerArgs)

					positionIndex++
				}

				return handlerArgsGroup
			}(),
			wantErr: errors.New("dummy"),
		},
	} {
		var gotHandlerArgsGroup []handlerArgs
		gotErr := data.pieceGroup.
			IteratePiecesByColorAndKind(data.args.makeHandler(&gotHandlerArgsGroup))

		if !reflect.DeepEqual(gotHandlerArgsGroup, data.wantHandlerArgsGroup) {
			test.Fail()
		}
		if !reflect.DeepEqual(gotErr, data.wantErr) {
			test.Fail()
		}
	}
}

func TestBitBoardPieceGroupPieceByPosition(test *testing.T) {
	type args struct {
		size         common.Size
		position     common.Position
		pieceFactory common.PieceFactory
	}
	type data struct {
		pieceGroup               *bitBoardPieceGroup
		args                     args
		wantPiece                common.Piece
		wantPiecesByColorAndKind *bitBoardPieceGroupByColorAndKind
		wantOk                   bool
	}

	for _, data := range []data{
		{
			pieceGroup: new(bitBoardPieceGroup),
			args: args{
				size:     common.Size{5, 5},
				position: common.Position{2, 3},
				pieceFactory: func(
					kind common.Kind,
					color common.Color,
					position common.Position,
				) common.Piece {
					return MockPiece{kind: kind, color: color, position: position}
				},
			},
			wantPiece:                nil,
			wantPiecesByColorAndKind: nil,
			wantOk:                   false,
		},
		{
			pieceGroup: func() *bitBoardPieceGroup {
				pieceGroup := new(bitBoardPieceGroup)
				pieceGroup.AddPiece(common.Size{5, 5}, MockPiece{
					kind:     common.Queen,
					color:    common.White,
					position: common.Position{2, 3},
				})

				return pieceGroup
			}(),
			args: args{
				size:     common.Size{5, 5},
				position: common.Position{2, 3},
				pieceFactory: func(
					kind common.Kind,
					color common.Color,
					position common.Position,
				) common.Piece {
					return MockPiece{kind: kind, color: color, position: position}
				},
			},
			wantPiece: MockPiece{
				kind:     common.Queen,
				color:    common.White,
				position: common.Position{2, 3},
			},
			wantPiecesByColorAndKind: func() *bitBoardPieceGroupByColorAndKind {
				piecesByColorAndKind := new(bitBoardPieceGroupByColorAndKind)
				piecesByColorAndKind.SetPositionStatus(
					common.Size{5, 5},
					common.Position{2, 3},
					occupiedPositionStatus,
				)

				return piecesByColorAndKind
			}(),
			wantOk: true,
		},
	} {
		gotPiece, gotPiecesByColorAndKind, gotOk := data.pieceGroup.PieceByPosition(
			data.args.size,
			data.args.position,
			data.args.pieceFactory,
		)

		if !reflect.DeepEqual(gotPiece, data.wantPiece) {
			test.Fail()
		}
		if !reflect.DeepEqual(
			gotPiecesByColorAndKind,
			data.wantPiecesByColorAndKind,
		) {
			test.Fail()
		}
		if gotOk != data.wantOk {
			test.Fail()
		}
	}
}

func TestBitBoardPieceGroupSetValue(test *testing.T) {
	expectedPieceGroup := new(bitBoardPieceGroup)
	expectedPieceGroup.AddPiece(common.Size{5, 5}, MockPiece{
		kind:     common.Queen,
		color:    common.White,
		position: common.Position{2, 3},
	})

	pieceGroup := new(bitBoardPieceGroup)
	pieceGroup.SetValue(expectedPieceGroup)

	if !reflect.DeepEqual(pieceGroup, expectedPieceGroup) {
		test.Fail()
	}
}

func TestBitBoardPieceGroupAddPiece(test *testing.T) {
	pieceGroup := new(bitBoardPieceGroup)
	pieceGroup.AddPiece(common.Size{5, 5}, MockPiece{
		kind:     common.Queen,
		color:    common.White,
		position: common.Position{2, 3},
	})

	expectedPieceGroup := new(bitBoardPieceGroup)
	expectedPieceGroup.IteratePiecesByColorAndKind(func( // nolint: errcheck
		color common.Color,
		kind common.Kind,
		piecesByColorAndKind *bitBoardPieceGroupByColorAndKind,
	) error {
		if color == common.White && kind == common.Queen {
			piecesByColorAndKind.SetPositionStatus(
				common.Size{5, 5},
				common.Position{2, 3},
				occupiedPositionStatus,
			)
		}

		return nil
	})
	if !reflect.DeepEqual(pieceGroup, expectedPieceGroup) {
		test.Fail()
	}
}

func TestBitBoardPieceGroupClearPosition(test *testing.T) {
	type args struct {
		size         common.Size
		position     common.Position
		pieceFactory common.PieceFactory
	}
	type data struct {
		pieceGroup     *bitBoardPieceGroup
		args           args
		wantPieceGroup *bitBoardPieceGroup
		wantPiece      common.Piece
		wantOk         bool
	}

	for _, data := range []data{
		{
			pieceGroup: new(bitBoardPieceGroup),
			args: args{
				size:     common.Size{5, 5},
				position: common.Position{2, 3},
				pieceFactory: func(
					kind common.Kind,
					color common.Color,
					position common.Position,
				) common.Piece {
					return MockPiece{kind: kind, color: color, position: position}
				},
			},
			wantPieceGroup: new(bitBoardPieceGroup),
			wantPiece:      nil,
			wantOk:         false,
		},
		{
			pieceGroup: func() *bitBoardPieceGroup {
				pieceGroup := new(bitBoardPieceGroup)
				pieceGroup.AddPiece(common.Size{5, 5}, MockPiece{
					kind:     common.Queen,
					color:    common.White,
					position: common.Position{2, 3},
				})

				return pieceGroup
			}(),
			args: args{
				size:     common.Size{5, 5},
				position: common.Position{2, 3},
				pieceFactory: func(
					kind common.Kind,
					color common.Color,
					position common.Position,
				) common.Piece {
					return MockPiece{kind: kind, color: color, position: position}
				},
			},
			wantPieceGroup: new(bitBoardPieceGroup),
			wantPiece: MockPiece{
				kind:     common.Queen,
				color:    common.White,
				position: common.Position{2, 3},
			},
			wantOk: true,
		},
	} {
		gotPiece, gotOk := data.pieceGroup.ClearPosition(
			data.args.size,
			data.args.position,
			data.args.pieceFactory,
		)

		if !reflect.DeepEqual(gotPiece, data.wantPiece) {
			test.Fail()
		}
		if gotOk != data.wantOk {
			test.Fail()
		}
	}
}
