package chessmodels

import (
	"reflect"
	"testing"
)

func TestSizeHasPosition(test *testing.T) {
	type fields struct {
		Width  int
		Height int
	}
	type args struct {
		position Position
	}
	type data struct {
		fields fields
		args   args
		want   bool
	}

	for _, data := range []data{
		{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{4, 1},
			},
			want: true,
		},
		{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{-1, 1},
			},
			want: false,
		},
		{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{4, -1},
			},
			want: false,
		},
		{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{-1, -1},
			},
			want: false,
		},
		{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{10, 1},
			},
			want: false,
		},
		{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{4, 10},
			},
			want: false,
		},
		{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{10, 10},
			},
			want: false,
		},
	} {
		size := Size{
			Width:  data.fields.Width,
			Height: data.fields.Height,
		}
		got := size.HasPosition(data.args.position)

		if got != data.want {
			test.Fail()
		}
	}
}

func TestSizeHasMove(test *testing.T) {
	type fields struct {
		Width  int
		Height int
	}
	type args struct {
		move Move
	}
	type data struct {
		fields fields
		args   args
		want   bool
	}

	for _, data := range []data{
		{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				move: Move{
					Start:  Position{4, 1},
					Finish: Position{4, 3},
				},
			},
			want: true,
		},
		{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				move: Move{
					Start:  Position{-1, 1},
					Finish: Position{4, 3},
				},
			},
			want: false,
		},
		{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				move: Move{
					Start:  Position{4, 1},
					Finish: Position{-1, 3},
				},
			},
			want: false,
		},
		{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				move: Move{
					Start:  Position{-1, 1},
					Finish: Position{-1, 3},
				},
			},
			want: false,
		},
	} {
		size := Size{
			Width:  data.fields.Width,
			Height: data.fields.Height,
		}
		got := size.HasMove(data.args.move)

		if got != data.want {
			test.Fail()
		}
	}
}

func TestSizePositionCount(test *testing.T) {
	positionCount := Size{3, 3}.PositionCount()

	if positionCount != 9 {
		test.Fail()
	}
}

func TestSizePositions(test *testing.T) {
	positions := Size{3, 3}.Positions()

	expectedPositions := []Position{
		{0, 0},
		{1, 0},
		{2, 0},
		{0, 1},
		{1, 1},
		{2, 1},
		{0, 2},
		{1, 2},
		{2, 2},
	}
	if !reflect.DeepEqual(positions, expectedPositions) {
		test.Fail()
	}
}
