package common

import (
	"errors"
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

func TestSizePositionIndex(test *testing.T) {
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
		want   int
	}

	for _, data := range []data{
		{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{0, 0},
			},
			want: 0,
		},
		{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{2, 3},
			},
			want: 26,
		},
		{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{7, 7},
			},
			want: 63,
		},
	} {
		size := Size{
			Width:  data.fields.Width,
			Height: data.fields.Height,
		}
		got := size.PositionIndex(data.args.position)

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

func TestSizeIteratePositions(test *testing.T) {
	type fields struct {
		Width  int
		Height int
	}
	type args struct {
		makeHandler func(positions *[]Position) PositionHandler
	}
	type data struct {
		fields        fields
		args          args
		wantPositions []Position
		wantErr       error
	}

	for _, data := range []data{
		{
			fields: fields{
				Width:  3,
				Height: 3,
			},
			args: args{
				makeHandler: func(positions *[]Position) PositionHandler {
					return func(position Position) error {
						*positions = append(*positions, position)
						return nil
					}
				},
			},
			wantPositions: []Position{
				{0, 0},
				{1, 0},
				{2, 0},
				{0, 1},
				{1, 1},
				{2, 1},
				{0, 2},
				{1, 2},
				{2, 2},
			},
			wantErr: nil,
		},
		{
			fields: fields{
				Width:  3,
				Height: 3,
			},
			args: args{
				makeHandler: func(positions *[]Position) PositionHandler {
					return func(position Position) error {
						if position.Rank > 1 {
							return errors.New("dummy")
						}

						*positions = append(*positions, position)
						return nil
					}
				},
			},
			wantPositions: []Position{
				{0, 0},
				{1, 0},
				{2, 0},
				{0, 1},
				{1, 1},
				{2, 1},
			},
			wantErr: errors.New("dummy"),
		},
	} {
		size := Size{
			Width:  data.fields.Width,
			Height: data.fields.Height,
		}

		var gotPositions []Position
		gotErr := size.IteratePositions(data.args.makeHandler(&gotPositions))

		if !reflect.DeepEqual(gotPositions, data.wantPositions) {
			test.Fail()
		}
		if !reflect.DeepEqual(gotErr, data.wantErr) {
			test.Fail()
		}
	}
}
