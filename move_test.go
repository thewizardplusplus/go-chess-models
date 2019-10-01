package chessmodels

import (
	"testing"
)

func TestMoveIsZero(test *testing.T) {
	type fields struct {
		start  Position
		finish Position
	}
	type data struct {
		fields fields
		want   bool
	}

	for _, data := range []data{
		data{
			fields: fields{
				start:  Position{1, 2},
				finish: Position{3, 4},
			},
			want: false,
		},
		data{
			fields: fields{
				start:  Position{1, 0},
				finish: Position{3, 0},
			},
			want: false,
		},
		data{
			fields: fields{
				start:  Position{1, 2},
				finish: Position{0, 0},
			},
			want: false,
		},
		data{
			fields: fields{
				start:  Position{0, 0},
				finish: Position{0, 0},
			},
			want: true,
		},
	} {
		move := Move{
			Start:  data.fields.start,
			Finish: data.fields.finish,
		}
		got := move.IsZero()

		if got != data.want {
			test.Fail()
		}
	}
}

func TestMoveIsEmpty(test *testing.T) {
	type fields struct {
		start  Position
		finish Position
	}
	type data struct {
		fields fields
		want   bool
	}

	for _, data := range []data{
		data{
			fields: fields{
				start:  Position{1, 2},
				finish: Position{3, 4},
			},
			want: false,
		},
		data{
			fields: fields{
				start:  Position{1, 2},
				finish: Position{1, 4},
			},
			want: false,
		},
		data{
			fields: fields{
				start:  Position{1, 2},
				finish: Position{3, 2},
			},
			want: false,
		},
		data{
			fields: fields{
				start:  Position{1, 2},
				finish: Position{1, 2},
			},
			want: true,
		},
	} {
		move := Move{
			Start:  data.fields.start,
			Finish: data.fields.finish,
		}
		got := move.IsEmpty()

		if got != data.want {
			test.Fail()
		}
	}
}
