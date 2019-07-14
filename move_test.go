package chessmodels

import (
	"errors"
	"reflect"
	"testing"
)

type MockPieceStorage struct {
	size   Size
	pieces []Piece

	checkMove func(move Move) error
}

func (
	storage MockPieceStorage,
) Size() Size {
	return storage.size
}

func (
	storage MockPieceStorage,
) Piece(
	position Position,
) (piece Piece, ok bool) {
	panic("not implemented")
}

func (
	storage MockPieceStorage,
) Pieces() []Piece {
	return storage.pieces
}

func (storage MockPieceStorage) ApplyMove(
	move Move,
) PieceStorage {
	panic("not implemented")
}

func (storage MockPieceStorage) CheckMove(
	move Move,
) error {
	if storage.checkMove == nil {
		panic("not implemented")
	}

	return storage.checkMove(move)
}

func (
	storage MockPieceStorage,
) ToFEN() string {
	panic("not implemented")
}

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
		fields    fields
		args      args
		wantMoves []Move
		wantErr   error
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
			wantMoves: []Move{
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
			wantErr: nil,
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
			wantMoves: []Move{
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
			wantErr: nil,
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
					return errors.New("dummy")
				},
			},
			args:      args{Black},
			wantMoves: nil,
			wantErr:   nil,
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
					return ErrKingCapture
				},
			},
			args:      args{Black},
			wantMoves: nil,
			wantErr:   ErrKingCapture,
		},
	} {
		storage := MockPieceStorage{
			size:      data.fields.size,
			pieces:    data.fields.pieces,
			checkMove: data.fields.checkMove,
		}
		generator := MoveGenerator{}
		gotMoves, gotErr :=
			generator.MovesForColor(
				storage,
				data.args.color,
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
