package chessmodels

import (
	"errors"
	"reflect"
	"testing"
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
		checkMove func(move Move) error
	}
	type args struct {
		color Color
	}
	type data struct {
		fields    fields
		args      args
		wantMoves []Move
		wantErr   error
	}

	for _, data := range []data{
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					MockPiece{
						color:    Black,
						position: Position{0, 0},
					},
					MockPiece{
						color:    Black,
						position: Position{0, 1},
					},
					MockPiece{
						color:    White,
						position: Position{1, 0},
					},
					MockPiece{
						color:    White,
						position: Position{1, 1},
					},
				},
				checkMove: func(move Move) error {
					return nil
				},
			},
			args: args{Black},
			wantMoves: []Move{
				{
					Start:  Position{0, 0},
					Finish: Position{0, 0},
				},
				{
					Start:  Position{0, 0},
					Finish: Position{1, 0},
				},
				{
					Start:  Position{0, 0},
					Finish: Position{0, 1},
				},
				{
					Start:  Position{0, 0},
					Finish: Position{1, 1},
				},

				{
					Start:  Position{0, 1},
					Finish: Position{0, 0},
				},
				{
					Start:  Position{0, 1},
					Finish: Position{1, 0},
				},
				{
					Start:  Position{0, 1},
					Finish: Position{0, 1},
				},
				{
					Start:  Position{0, 1},
					Finish: Position{1, 1},
				},
			},
			wantErr: nil,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					MockPiece{
						color:    Black,
						position: Position{0, 0},
					},
					MockPiece{
						color:    Black,
						position: Position{0, 1},
					},
					MockPiece{
						color:    White,
						position: Position{1, 0},
					},
					MockPiece{
						color:    White,
						position: Position{1, 1},
					},
				},
				checkMove: func(move Move) error {
					return nil
				},
			},
			args: args{White},
			wantMoves: []Move{
				{
					Start:  Position{1, 0},
					Finish: Position{0, 0},
				},
				{
					Start:  Position{1, 0},
					Finish: Position{1, 0},
				},
				{
					Start:  Position{1, 0},
					Finish: Position{0, 1},
				},
				{
					Start:  Position{1, 0},
					Finish: Position{1, 1},
				},

				{
					Start:  Position{1, 1},
					Finish: Position{0, 0},
				},
				{
					Start:  Position{1, 1},
					Finish: Position{1, 0},
				},
				{
					Start:  Position{1, 1},
					Finish: Position{0, 1},
				},
				{
					Start:  Position{1, 1},
					Finish: Position{1, 1},
				},
			},
			wantErr: nil,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					MockPiece{
						color:    Black,
						position: Position{0, 0},
					},
					MockPiece{
						color:    Black,
						position: Position{0, 1},
					},
					MockPiece{
						color:    White,
						position: Position{1, 0},
					},
					MockPiece{
						color:    White,
						position: Position{1, 1},
					},
				},
				checkMove: func(move Move) error {
					return errors.New("dummy")
				},
			},
			args:      args{Black},
			wantMoves: nil,
			wantErr:   nil,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					MockPiece{
						color:    Black,
						position: Position{0, 0},
					},
					MockPiece{
						color:    Black,
						position: Position{0, 1},
					},
					MockPiece{
						color:    White,
						position: Position{1, 0},
					},
					MockPiece{
						color:    White,
						position: Position{1, 1},
					},
				},
				checkMove: func(move Move) error {
					return ErrKingCapture
				},
			},
			args:      args{Black},
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
		checkMove func(move Move) error
	}
	type args struct {
		position Position
	}
	type data struct {
		fields    fields
		args      args
		wantMoves []Move
		wantErr   error
	}

	for _, data := range []data{
		{
			fields: fields{
				size: Size{2, 2},
				checkMove: func(move Move) error {
					return errors.New("dummy")
				},
			},
			args:      args{Position{1, 1}},
			wantMoves: nil,
			wantErr:   nil,
		},
		{
			fields: fields{
				size: Size{2, 2},
				checkMove: func(move Move) error {
					return ErrKingCapture
				},
			},
			args:      args{Position{1, 1}},
			wantMoves: nil,
			wantErr:   ErrKingCapture,
		},
		{
			fields: fields{
				size: Size{2, 2},
				checkMove: func(move Move) error {
					return nil
				},
			},
			args: args{Position{1, 1}},
			wantMoves: []Move{
				{
					Start:  Position{1, 1},
					Finish: Position{0, 0},
				},
				{
					Start:  Position{1, 1},
					Finish: Position{1, 0},
				},
				{
					Start:  Position{1, 1},
					Finish: Position{0, 1},
				},
				{
					Start:  Position{1, 1},
					Finish: Position{1, 1},
				},
			},
			wantErr: nil,
		},
		{
			fields: fields{
				size: Size{2, 2},
				checkMove: func(move Move) error {
					if move.Finish.Rank == 1 {
						return errors.New("dummy")
					}

					return nil
				},
			},
			args: args{Position{1, 1}},
			wantMoves: []Move{
				{
					Start:  Position{1, 1},
					Finish: Position{0, 0},
				},
				{
					Start:  Position{1, 1},
					Finish: Position{1, 0},
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
