package chessmodels

import (
	"reflect"
	"testing"
)

func TestNewSliceBoard(test *testing.T) {
	board := NewSliceBoard(Size{5, 5}, []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})

	expectedBoard := SliceBoard{
		size: Size{5, 5},
		pieces: []Piece{
			14: MockPiece{position: Position{4, 2}},
			17: MockPiece{position: Position{2, 3}},
			24: nil,
		},
	}
	if !reflect.DeepEqual(board, expectedBoard) {
		test.Fail()
	}
}

func TestSliceBoardSize(test *testing.T) {
	board := NewSliceBoard(Size{5, 5}, nil)
	size := board.Size()

	if !reflect.DeepEqual(size, Size{5, 5}) {
		test.Fail()
	}
}

func TestSliceBoardPiece(test *testing.T) {
	type fields struct {
		size   Size
		pieces []Piece
	}
	type args struct {
		position Position
	}
	type data struct {
		fields    fields
		args      args
		wantPiece Piece
		wantOk    bool
	}

	for _, data := range []data{
		{
			fields: fields{
				size: Size{5, 5},
				pieces: []Piece{
					14: MockPiece{position: Position{4, 2}},
					17: MockPiece{position: Position{2, 3}},
					24: nil,
				},
			},
			args: args{Position{2, 3}},
			wantPiece: MockPiece{
				position: Position{2, 3},
			},
			wantOk: true,
		},
		{
			fields: fields{
				size: Size{5, 5},
				pieces: []Piece{
					14: MockPiece{position: Position{4, 2}},
					17: MockPiece{position: Position{2, 3}},
					24: nil,
				},
			},
			args:      args{Position{0, 0}},
			wantPiece: nil,
			wantOk:    false,
		},
	} {
		board := SliceBoard{
			size:   data.fields.size,
			pieces: data.fields.pieces,
		}
		gotPiece, gotOk := board.Piece(data.args.position)

		if !reflect.DeepEqual(gotPiece, data.wantPiece) {
			test.Fail()
		}
		if gotOk != data.wantOk {
			test.Fail()
		}
	}
}

func TestSliceBoardPieces(test *testing.T) {
	board := NewSliceBoard(Size{5, 5}, []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})
	pieces := board.Pieces()

	expectedPieces := []Piece{
		MockPiece{position: Position{4, 2}},
		MockPiece{position: Position{2, 3}},
	}
	if !reflect.DeepEqual(pieces, expectedPieces) {
		test.Fail()
	}
}

func TestSliceBoardApplyMove(test *testing.T) {
	board := NewSliceBoard(Size{5, 5}, []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})
	nextBoard := board.ApplyMove(Move{
		Start:  Position{4, 2},
		Finish: Position{1, 2},
	})

	expectedBoard := SliceBoard{
		size: Size{5, 5},
		pieces: []Piece{
			14: MockPiece{position: Position{4, 2}},
			17: MockPiece{position: Position{2, 3}},
			24: nil,
		},
	}
	if !reflect.DeepEqual(board, expectedBoard) {
		test.Fail()
	}

	expectedNextBoard := SliceBoard{
		size: Size{5, 5},
		pieces: []Piece{
			11: MockPiece{position: Position{1, 2}},
			17: MockPiece{position: Position{2, 3}},
			24: nil,
		},
	}
	if !reflect.DeepEqual(nextBoard, expectedNextBoard) {
		test.Fail()
	}
}

func TestSliceBoardCheckMove(test *testing.T) {
	type fields struct {
		size   Size
		pieces []Piece
	}
	type args struct {
		move Move
	}
	type data struct {
		fields fields
		args   args
		want   error
	}

	for _, data := range []data{
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					3: nil,
				},
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{0, 0},
				},
			},
			want: ErrNoMove,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					3: nil,
				},
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{-1, -1},
				},
			},
			want: ErrOutOfSize,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					3: nil,
				},
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{1, 1},
				},
			},
			want: ErrNoPiece,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					0: MockPiece{
						color:    Black,
						position: Position{0, 0},
					},
					3: MockPiece{
						color:    Black,
						position: Position{1, 1},
					},
				},
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{1, 1},
				},
			},
			want: ErrFriendlyTarget,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					0: MockPiece{
						position: Position{0, 0},
						checkMove: func(move Move, storage PieceStorage) bool {
							return false
						},
					},
					3: nil,
				},
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{1, 1},
				},
			},
			want: ErrIllegalMove,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					0: MockPiece{
						color:    Black,
						position: Position{0, 0},
						checkMove: func(move Move, storage PieceStorage) bool {
							return true
						},
					},
					3: MockPiece{
						kind:     King,
						color:    White,
						position: Position{1, 1},
					},
				},
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{1, 1},
				},
			},
			want: ErrKingCapture,
		},
		{
			fields: fields{
				size: Size{2, 2},
				pieces: []Piece{
					0: MockPiece{
						position: Position{0, 0},
						checkMove: func(move Move, storage PieceStorage) bool {
							return true
						},
					},
					3: nil,
				},
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{1, 1},
				},
			},
			want: nil,
		},
	} {
		board := SliceBoard{
			size:   data.fields.size,
			pieces: data.fields.pieces,
		}
		got := board.CheckMove(data.args.move)

		if got != data.want {
			test.Fail()
		}
	}
}