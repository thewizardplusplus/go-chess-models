package chessmodels

import (
	"errors"
	"reflect"
	"testing"
)

type MockPieceStorage struct {
	size      Size
	pieces    []Piece
	checkMove func(move Move) error
}

func (
	storage MockPieceStorage,
) Size() Size {
	return storage.size
}

func (
	storage MockPieceStorage,
) Pieces() []Piece {
	return storage.pieces
}

func (storage MockPieceStorage) CheckMove(
	move Move,
) error {
	return storage.checkMove(move)
}

func TestMoveCheckerMovesForColor(
	test *testing.T,
) {
	type args struct {
		color Color
	}
	type data struct {
		args      args
		checkMove func(move Move) error
		want      []Move
	}

	for _, data := range []data{
		data{
			args: args{Black},
			checkMove: func(move Move) error {
				return nil
			},
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
			args: args{White},
			checkMove: func(move Move) error {
				return nil
			},
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
			checkMove: data.checkMove,
		}
		got := MoveGenerator{storage}.
			MovesForColor(data.args.color)

		if !reflect.DeepEqual(got, data.want) {
			test.Fail()
		}
	}
}

func TestMoveCheckerMovesForPosition(
	test *testing.T,
) {
	type args struct {
		start Position
	}
	type data struct {
		args      args
		checkMove func(move Move) error
		want      []Move
	}

	for _, data := range []data{
		data{
			args: args{Position{1, 1}},
			checkMove: func(move Move) error {
				return errors.New("dummy")
			},
			want: nil,
		},
		data{
			args: args{Position{1, 1}},
			checkMove: func(move Move) error {
				return nil
			},
			want: []Move{
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
		data{
			args: args{Position{1, 1}},
			checkMove: func(move Move) error {
				if move.Finish.Rank == 1 {
					return errors.New("dummy")
				}

				return nil
			},
			want: []Move{
				Move{
					Start:  Position{1, 1},
					Finish: Position{0, 0},
				},
				Move{
					Start:  Position{1, 1},
					Finish: Position{1, 0},
				},
			},
		},
	} {
		storage := MockPieceStorage{
			size:      Size{2, 2},
			checkMove: data.checkMove,
		}
		got := MoveGenerator{storage}.
			MovesForPosition(data.args.start)

		if !reflect.DeepEqual(got, data.want) {
			test.Fail()
		}
	}
}
