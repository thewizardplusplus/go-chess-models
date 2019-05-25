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
	allowedCheck bool,
) error {
	return checker.handler(move)
}

func TestMoveCheckerLegalMovesForPosition(
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
				Move{Position{1, 1}, Position{0, 0}},
				Move{Position{1, 1}, Position{1, 0}},
				Move{Position{1, 1}, Position{0, 1}},
				Move{Position{1, 1}, Position{1, 1}},
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
				Move{Position{1, 1}, Position{0, 0}},
				Move{Position{1, 1}, Position{1, 0}},
			},
		},
	} {
		board := NewBoard(Size{2, 2}, nil)
		checker := MockMoveChecker{data.checker}
		got := MoveGenerator{board, checker}.
			LegalMovesForPosition(
			data.args.start,
			false,
		)
		if !reflect.DeepEqual(got, data.want) {
			test.Log(got)
			test.Log(data.want)
			test.Fail()
		}
	}
}
