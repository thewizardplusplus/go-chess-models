package chessmodels

import (
	"errors"
	"reflect"
	"sort"
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

type MoveGroup []Move

func (group MoveGroup) Len() int {
	return len(group)
}

func (group MoveGroup) Swap(i, j int) {
	group[i], group[j] = group[j], group[i]
}

func (group MoveGroup) Less(i, j int) bool {
	a, b := group[i], group[j]
	return positionLess(a.Start, b.Start) &&
		positionLess(a.Finish, b.Finish)
}

func TestMoveCheckerMovesForColor(
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
		got := MoveGenerator{board, checker}.
			MovesForColor(data.args.color)
		sort.Sort(MoveGroup(got))

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
		got := MoveGenerator{board, checker}.
			MovesForPosition(data.args.start)

		if !reflect.DeepEqual(got, data.want) {
			test.Fail()
		}
	}
}
