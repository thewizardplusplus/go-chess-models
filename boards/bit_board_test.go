package boards

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

func TestNewBitBoard(test *testing.T) {
	board := NewBitBoard(
		common.Size{5, 5},
		[]common.Piece{
			MockPiece{
				kind:     common.King,
				color:    common.Black,
				position: common.Position{2, 3},
			},
			MockPiece{
				kind:     common.Queen,
				color:    common.White,
				position: common.Position{4, 2},
			},
		},
		func(
			kind common.Kind,
			color common.Color,
			position common.Position,
		) common.Piece {
			return MockPiece{kind: kind, color: color, position: position}
		},
	)

	expectedBitBoard := BitBoard{
		BaseBoard: NewBaseBoard(common.Size{5, 5}),

		pieces: new(bitBoardPieceGroup),
	}
	expectedBitBoard.pieces.AddPiece(common.Size{5, 5}, MockPiece{
		kind:     common.King,
		color:    common.Black,
		position: common.Position{2, 3},
	})
	expectedBitBoard.pieces.AddPiece(common.Size{5, 5}, MockPiece{
		kind:     common.Queen,
		color:    common.White,
		position: common.Position{4, 2},
	})

	if !isCorrectBitBoard(test, board, expectedBitBoard) {
		test.Fail()
	}
}

func TestBitBoardPiece(test *testing.T) {
	type fields struct {
		size         common.Size
		pieces       *bitBoardPieceGroup
		pieceFactory common.PieceFactory
	}
	type args struct {
		position common.Position
	}
	type data struct {
		fields    fields
		args      args
		wantPiece common.Piece
		wantOk    bool
	}

	for _, data := range []data{
		{
			fields: fields{
				size: common.Size{5, 5},
				pieces: func() *bitBoardPieceGroup {
					pieceGroup := new(bitBoardPieceGroup)
					pieceGroup.AddPiece(common.Size{5, 5}, MockPiece{
						kind:     common.King,
						color:    common.Black,
						position: common.Position{2, 3},
					})
					pieceGroup.AddPiece(common.Size{5, 5}, MockPiece{
						kind:     common.Queen,
						color:    common.White,
						position: common.Position{4, 2},
					})

					return pieceGroup
				}(),
				pieceFactory: func(
					kind common.Kind,
					color common.Color,
					position common.Position,
				) common.Piece {
					return MockPiece{kind: kind, color: color, position: position}
				},
			},
			args: args{
				position: common.Position{2, 3},
			},
			wantPiece: MockPiece{
				kind:     common.King,
				color:    common.Black,
				position: common.Position{2, 3},
			},
			wantOk: true,
		},
		{
			fields: fields{
				size: common.Size{5, 5},
				pieces: func() *bitBoardPieceGroup {
					pieceGroup := new(bitBoardPieceGroup)
					pieceGroup.AddPiece(common.Size{5, 5}, MockPiece{
						kind:     common.King,
						color:    common.Black,
						position: common.Position{2, 3},
					})
					pieceGroup.AddPiece(common.Size{5, 5}, MockPiece{
						kind:     common.Queen,
						color:    common.White,
						position: common.Position{4, 2},
					})

					return pieceGroup
				}(),
				pieceFactory: func(
					kind common.Kind,
					color common.Color,
					position common.Position,
				) common.Piece {
					return MockPiece{kind: kind, color: color, position: position}
				},
			},
			args: args{
				position: common.Position{0, 0},
			},
			wantPiece: nil,
			wantOk:    false,
		},
	} {
		board := BitBoard{
			BaseBoard: NewBaseBoard(data.fields.size),

			pieces:       data.fields.pieces,
			pieceFactory: data.fields.pieceFactory,
		}
		gotPiece, gotOk := board.Piece(data.args.position)

		if !reflect.DeepEqual(gotPiece, data.wantPiece) {
			test.Fail()
		}
		if gotOk != data.wantOk {
			test.Fail()
		}
	}
}

func TestBitBoardApplyMove(test *testing.T) {
	type fields struct {
		size         common.Size
		pieces       *bitBoardPieceGroup
		pieceFactory common.PieceFactory
	}
	type args struct {
		move common.Move
	}
	type data struct {
		fields        fields
		args          args
		wantNextBoard BitBoard
	}

	for _, data := range []data{
		{
			fields: fields{
				size: common.Size{5, 5},
				pieces: func() *bitBoardPieceGroup {
					pieceGroup := new(bitBoardPieceGroup)
					pieceGroup.AddPiece(common.Size{5, 5}, MockPiece{
						kind:     common.King,
						color:    common.Black,
						position: common.Position{2, 3},
					})
					pieceGroup.AddPiece(common.Size{5, 5}, MockPiece{
						kind:     common.Queen,
						color:    common.White,
						position: common.Position{4, 2},
					})

					return pieceGroup
				}(),
				pieceFactory: func(
					kind common.Kind,
					color common.Color,
					position common.Position,
				) common.Piece {
					return MockPiece{kind: kind, color: color, position: position}
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{4, 2},
					Finish: common.Position{1, 2},
				},
			},
			wantNextBoard: func() BitBoard {
				wantNextBoard := BitBoard{
					BaseBoard: NewBaseBoard(common.Size{5, 5}),

					pieces: new(bitBoardPieceGroup),
				}
				wantNextBoard.pieces.AddPiece(common.Size{5, 5}, MockPiece{
					kind:     common.King,
					color:    common.Black,
					position: common.Position{2, 3},
				})
				wantNextBoard.pieces.AddPiece(common.Size{5, 5}, MockPiece{
					kind:     common.Queen,
					color:    common.White,
					position: common.Position{1, 2},
				})

				return wantNextBoard
			}(),
		},
		{
			fields: fields{
				size: common.Size{5, 5},
				pieces: func() *bitBoardPieceGroup {
					pieceGroup := new(bitBoardPieceGroup)
					pieceGroup.AddPiece(common.Size{5, 5}, MockPiece{
						kind:     common.King,
						color:    common.Black,
						position: common.Position{2, 3},
					})
					pieceGroup.AddPiece(common.Size{5, 5}, MockPiece{
						kind:     common.Queen,
						color:    common.White,
						position: common.Position{4, 2},
					})

					return pieceGroup
				}(),
				pieceFactory: func(
					kind common.Kind,
					color common.Color,
					position common.Position,
				) common.Piece {
					return MockPiece{kind: kind, color: color, position: position}
				},
			},
			args: args{
				move: common.Move{
					Start:  common.Position{4, 2},
					Finish: common.Position{2, 3},
				},
			},
			wantNextBoard: func() BitBoard {
				wantNextBoard := BitBoard{
					BaseBoard: NewBaseBoard(common.Size{5, 5}),

					pieces: new(bitBoardPieceGroup),
				}
				wantNextBoard.pieces[common.Black][common.King].
					ToBigInt().SetBits([]big.Word{})
				wantNextBoard.pieces.AddPiece(common.Size{5, 5}, MockPiece{
					kind:     common.Queen,
					color:    common.White,
					position: common.Position{2, 3},
				})

				return wantNextBoard
			}(),
		},
	} {
		board := BitBoard{
			BaseBoard: NewBaseBoard(data.fields.size),

			pieces:       data.fields.pieces,
			pieceFactory: data.fields.pieceFactory,
		}
		gotNextBoard := board.ApplyMove(data.args.move)

		if !isCorrectBitBoard(test, gotNextBoard, data.wantNextBoard) {
			test.Fail()
		}
	}
}

func isCorrectBitBoard(
	test *testing.T,
	actualStorage common.PieceStorage,
	expectedBitBoard BitBoard,
) bool {
	actualStorageWrapper, ok := actualStorage.(pieceStorageWrapper)
	if !ok {
		return false
	}

	actualBitBoard, ok := actualStorageWrapper.BasePieceStorage.(BitBoard)
	if !ok {
		return false
	}

	if actualBitBoard.pieceFactory == nil {
		return false
	}
	actualBitBoard.pieceFactory = nil

	return reflect.DeepEqual(actualBitBoard, expectedBitBoard)
}
