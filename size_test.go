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
		data{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{4, 1},
			},
			want: true,
		},
		data{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{-1, 1},
			},
			want: false,
		},
		data{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{4, -1},
			},
			want: false,
		},
		data{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{-1, -1},
			},
			want: false,
		},
		data{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{10, 1},
			},
			want: false,
		},
		data{
			fields: fields{
				Width:  8,
				Height: 8,
			},
			args: args{
				position: Position{4, 10},
			},
			want: false,
		},
		data{
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
		got :=
			size.HasPosition(data.args.position)

		if got != data.want {
			test.Fail()
		}
	}
}

func TestSizePositions(test *testing.T) {
	positions := Size{3, 3}.Positions()

	expectedPositions := []Position{
		Position{0, 0},
		Position{1, 0},
		Position{2, 0},
		Position{0, 1},
		Position{1, 1},
		Position{2, 1},
		Position{0, 2},
		Position{1, 2},
		Position{2, 2},
	}
	if !reflect.DeepEqual(
		positions,
		expectedPositions,
	) {
		test.Fail()
	}
}
