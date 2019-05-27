package chessmodels

import (
	"errors"
	"reflect"
	"testing"
)

type MockMoveChecker struct {
	handler func(move Move) error
}

func (checker MockMoveChecker) CheckMove(
	move Move,
) error {
	return checker.handler(move)
}

func TestDefaultMoveCheckerMovesForColor(
	test *testing.T,
) {
	type args struct {
		color Color
	}
	type data struct {
		args    args
		checker func(move Move) error
		want    []Move
	}

	for _, data := range []data{
		data{
			args: args{Black},
			checker: func(move Move) error {
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
			checker: func(move Move) error {
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
		board := NewBoard(Size{2, 2}, []Piece{
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
		})
		checker := MockMoveChecker{data.checker}
		generator := DefaultMoveGenerator{
			Board:       board,
			MoveChecker: checker,
		}
		got := generator.MovesForColor(
			data.args.color,
		)
		if !reflect.DeepEqual(got, data.want) {
			test.Fail()
		}
	}
}

func TestDefaultMoveCheckerMovesForPosition(
	test *testing.T,
) {
	type args struct {
		start Position
	}
	type data struct {
		args    args
		checker func(move Move) error
		want    []Move
	}

	for _, data := range []data{
		data{
			args: args{Position{1, 1}},
			checker: func(move Move) error {
				return errors.New("dummy")
			},
			want: nil,
		},
		data{
			args: args{Position{1, 1}},
			checker: func(move Move) error {
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
			checker: func(move Move) error {
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
		board := NewBoard(Size{2, 2}, nil)
		checker := MockMoveChecker{data.checker}
		generator := DefaultMoveGenerator{
			Board:       board,
			MoveChecker: checker,
		}
		got := generator.MovesForPosition(
			data.args.start,
		)
		if !reflect.DeepEqual(got, data.want) {
			test.Fail()
		}
	}
}
