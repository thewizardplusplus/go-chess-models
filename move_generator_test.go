package chessmodels

import (
	"errors"
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

type MockPieceStorage struct {
	MockBasePieceStorage
	MockPieceGroupGetter
	MockMoveChecker
}

func TestMoveCheckerMovesForColor(test *testing.T) {
	type fields struct {
		size      Size
		pieces    []Piece
		checkMove func(move common.Move) error
	}
	type args struct {
		color common.Color
	}
	type data struct {
		fields    fields
		args      args
		wantMoves []common.Move
		wantErr   error
	}

	for _, data := range []data{
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 0},
					},
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 1},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 0},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 1},
					},
				},
				checkMove: func(move common.Move) error {
					return nil
				},
			},
			args: args{common.Black},
			wantMoves: []common.Move{
				{
					Start:  common.Position{0, 0},
					Finish: common.Position{0, 0},
				},
				{
					Start:  common.Position{0, 0},
					Finish: common.Position{1, 0},
				},
				{
					Start:  common.Position{0, 0},
					Finish: common.Position{0, 1},
				},
				{
					Start:  common.Position{0, 0},
					Finish: common.Position{1, 1},
				},

				{
					Start:  common.Position{0, 1},
					Finish: common.Position{0, 0},
				},
				{
					Start:  common.Position{0, 1},
					Finish: common.Position{1, 0},
				},
				{
					Start:  common.Position{0, 1},
					Finish: common.Position{0, 1},
				},
				{
					Start:  common.Position{0, 1},
					Finish: common.Position{1, 1},
				},
			},
			wantErr: nil,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 0},
					},
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 1},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 0},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 1},
					},
				},
				checkMove: func(move common.Move) error {
					return nil
				},
			},
			args: args{common.White},
			wantMoves: []common.Move{
				{
					Start:  common.Position{1, 0},
					Finish: common.Position{0, 0},
				},
				{
					Start:  common.Position{1, 0},
					Finish: common.Position{1, 0},
				},
				{
					Start:  common.Position{1, 0},
					Finish: common.Position{0, 1},
				},
				{
					Start:  common.Position{1, 0},
					Finish: common.Position{1, 1},
				},

				{
					Start:  common.Position{1, 1},
					Finish: common.Position{0, 0},
				},
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{1, 0},
				},
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{0, 1},
				},
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{1, 1},
				},
			},
			wantErr: nil,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 0},
					},
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 1},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 0},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 1},
					},
				},
				checkMove: func(move common.Move) error {
					return errors.New("dummy")
				},
			},
			args:      args{common.Black},
			wantMoves: nil,
			wantErr:   nil,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 0},
					},
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 1},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 0},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 1},
					},
				},
				checkMove: func(move common.Move) error {
					return ErrKingCapture
				},
			},
			args:      args{common.Black},
			wantMoves: nil,
			wantErr:   ErrKingCapture,
		},
	} {
		storage := MockPieceStorage{
			MockBasePieceStorage: MockBasePieceStorage{
				size: data.fields.size,
			},
			MockPieceGroupGetter: MockPieceGroupGetter{
				pieces: data.fields.pieces,
			},
			MockMoveChecker: MockMoveChecker{
				checkMove: data.fields.checkMove,
			},
		}
		var generator MoveGenerator
		gotMoves, gotErr := generator.MovesForColor(storage, data.args.color)

		if !reflect.DeepEqual(gotMoves, data.wantMoves) {
			test.Fail()
		}
		if !reflect.DeepEqual(gotErr, data.wantErr) {
			test.Fail()
		}
	}
}

func TestMoveCheckerMovesForPosition(test *testing.T) {
	type fields struct {
		size      Size
		checkMove func(move common.Move) error
	}
	type args struct {
		position common.Position
	}
	type data struct {
		fields    fields
		args      args
		wantMoves []common.Move
		wantErr   error
	}

	for _, data := range []data{
		{
			fields: fields{
				size: Size{2, 2},
				checkMove: func(move common.Move) error {
					return errors.New("dummy")
				},
			},
			args:      args{common.Position{1, 1}},
			wantMoves: nil,
			wantErr:   nil,
		},
		{
			fields: fields{
				size: Size{2, 2},
				checkMove: func(move common.Move) error {
					return ErrKingCapture
				},
			},
			args:      args{common.Position{1, 1}},
			wantMoves: nil,
			wantErr:   ErrKingCapture,
		},
		{
			fields: fields{
				size: Size{2, 2},
				checkMove: func(move common.Move) error {
					return nil
				},
			},
			args: args{common.Position{1, 1}},
			wantMoves: []common.Move{
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{0, 0},
				},
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{1, 0},
				},
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{0, 1},
				},
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{1, 1},
				},
			},
			wantErr: nil,
		},
		{
			fields: fields{
				size: Size{2, 2},
				checkMove: func(move common.Move) error {
					if move.Finish.Rank == 1 {
						return errors.New("dummy")
					}

					return nil
				},
			},
			args: args{common.Position{1, 1}},
			wantMoves: []common.Move{
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{0, 0},
				},
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{1, 0},
				},
			},
			wantErr: nil,
		},
	} {
		storage := MockPieceStorage{
			MockBasePieceStorage: MockBasePieceStorage{
				size: data.fields.size,
			},
			MockMoveChecker: MockMoveChecker{
				checkMove: data.fields.checkMove,
			},
		}
		var generator MoveGenerator
		gotMoves, gotErr := generator.MovesForPosition(storage, data.args.position)

		if !reflect.DeepEqual(gotMoves, data.wantMoves) {
			test.Fail()
		}
		if !reflect.DeepEqual(gotErr, data.wantErr) {
			test.Fail()
		}
	}
}
