package chessmodels

import (
	"errors"
	"reflect"
	"testing"
)

func TestMoveCheckerMovesForColor(
	test *testing.T,
) {
	type fields struct {
		size      Size
		pieces    []Piece
		checkMove func(move Move) error
	}
	type args struct {
		color Color
	}
	type data struct {
		fields fields
		args   args
		want   []Move
	}

	for _, data := range []data{
		data{
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
			want: []Move{
				Move{
					Start:  Position{0, 0},
					Finish: Position{0, 0},
				},
				Move{
					Start:  Position{0, 0},
					Finish: Position{1, 0},
				},
				Move{
					Start:  Position{0, 0},
					Finish: Position{0, 1},
				},
				Move{
					Start:  Position{0, 0},
					Finish: Position{1, 1},
				},

				Move{
					Start:  Position{0, 1},
					Finish: Position{0, 0},
				},
				Move{
					Start:  Position{0, 1},
					Finish: Position{1, 0},
				},
				Move{
					Start:  Position{0, 1},
					Finish: Position{0, 1},
				},
				Move{
					Start:  Position{0, 1},
					Finish: Position{1, 1},
				},
			},
		},
		data{
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
			want: []Move{
				Move{
					Start:  Position{1, 0},
					Finish: Position{0, 0},
				},
				Move{
					Start:  Position{1, 0},
					Finish: Position{1, 0},
				},
				Move{
					Start:  Position{1, 0},
					Finish: Position{0, 1},
				},
				Move{
					Start:  Position{1, 0},
					Finish: Position{1, 1},
				},

				Move{
					Start:  Position{1, 1},
					Finish: Position{0, 0},
				},
				Move{
					Start:  Position{1, 1},
					Finish: Position{1, 0},
				},
				Move{
					Start:  Position{1, 1},
					Finish: Position{0, 1},
				},
				Move{
					Start:  Position{1, 1},
					Finish: Position{1, 1},
				},
			},
		},
	} {
		storage := MockPieceStorage{
			size:      data.fields.size,
			pieces:    data.fields.pieces,
			checkMove: data.fields.checkMove,
		}
		generator := MoveGenerator{}
		got := generator.MovesForColor(
			storage,
			data.args.color,
		)

		if !reflect.DeepEqual(got, data.want) {
			test.Fail()
		}
	}
}

func TestMoveCheckerMovesForPosition(
	test *testing.T,
) {
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
		checkMove func(move Move) error
		wantMoves []Move
		wantErr   error
	}

	for _, data := range []data{
		data{
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
		data{
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
		data{
			fields: fields{
				size: Size{2, 2},
				checkMove: func(move Move) error {
					return nil
				},
			},
			args: args{Position{1, 1}},
			wantMoves: []Move{
				Move{
					Start:  Position{1, 1},
					Finish: Position{0, 0},
				},
				Move{
					Start:  Position{1, 1},
					Finish: Position{1, 0},
				},
				Move{
					Start:  Position{1, 1},
					Finish: Position{0, 1},
				},
				Move{
					Start:  Position{1, 1},
					Finish: Position{1, 1},
				},
			},
			wantErr: nil,
		},
		data{
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
				Move{
					Start:  Position{1, 1},
					Finish: Position{0, 0},
				},
				Move{
					Start:  Position{1, 1},
					Finish: Position{1, 0},
				},
			},
			wantErr: nil,
		},
	} {
		storage := MockPieceStorage{
			size:      data.fields.size,
			checkMove: data.fields.checkMove,
		}
		generator := MoveGenerator{}
		gotMoves, gotErr :=
			generator.MovesForPosition(
				storage,
				data.args.position,
			)

		if !reflect.DeepEqual(
			gotMoves,
			data.wantMoves,
		) {
			test.Fail()
		}
		if !reflect.DeepEqual(
			gotErr,
			data.wantErr,
		) {
			test.Fail()
		}
	}
}
