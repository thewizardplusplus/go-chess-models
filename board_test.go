package chessmodels

import (
	"reflect"
	"sort"
	"testing"
)

type ByPosition []Piece

func (group ByPosition) Len() int {
	return len(group)
}

func (group ByPosition) Swap(i, j int) {
	group[i], group[j] = group[j], group[i]
}

func (group ByPosition) Less(
	i, j int,
) bool {
	a := group[i].Position()
	b := group[j].Position()
	if a.File == b.File {
		return a.Rank < b.Rank
	}

	return a.File < b.File
}

func TestBoardInterface(test *testing.T) {
	boardType := reflect.TypeOf(Board{})
	storageType := reflect.
		TypeOf((*PieceStorage)(nil)).
		Elem()
	if !boardType.Implements(storageType) {
		test.Fail()
	}
}

func TestNewBoard(test *testing.T) {
	board := NewBoard(Size{5, 5}, []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})

	expectedBoard := Board{
		size: Size{5, 5},
		pieces: pieceGroup{
			Position{2, 3}: MockPiece{
				position: Position{2, 3},
			},
			Position{4, 2}: MockPiece{
				position: Position{4, 2},
			},
		},
	}
	if !reflect.DeepEqual(
		board,
		expectedBoard,
	) {
		test.Fail()
	}
}

func TestBoardSize(test *testing.T) {
	board := NewBoard(Size{5, 5}, nil)

	if !reflect.DeepEqual(
		board.size,
		Size{5, 5},
	) {
		test.Fail()
	}
}

func TestBoardPieces(test *testing.T) {
	board := NewBoard(Size{5, 5}, []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})
	pieces := board.Pieces()
	sort.Sort(ByPosition(pieces))

	expectedPieces := []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	}
	if !reflect.DeepEqual(
		pieces,
		expectedPieces,
	) {
		test.Fail()
	}
}

func TestBoardApplyMove(test *testing.T) {
	board := NewBoard(Size{5, 5}, []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})
	nextBoard := board.ApplyMove(Move{
		Start:  Position{4, 2},
		Finish: Position{6, 5},
	})

	expectedBoard := Board{
		size: Size{5, 5},
		pieces: pieceGroup{
			Position{2, 3}: MockPiece{
				position: Position{2, 3},
			},
			Position{4, 2}: MockPiece{
				position: Position{4, 2},
			},
		},
	}
	if !reflect.DeepEqual(
		board,
		expectedBoard,
	) {
		test.Fail()
	}

	expectedNextBoard := Board{
		size: Size{5, 5},
		pieces: pieceGroup{
			Position{2, 3}: MockPiece{
				position: Position{2, 3},
			},
			Position{6, 5}: MockPiece{
				position: Position{6, 5},
			},
		},
	}
	if !reflect.DeepEqual(
		nextBoard,
		expectedNextBoard,
	) {
		test.Fail()
	}
}

func TestBoardCheckMove(test *testing.T) {
	type fields struct {
		size   Size
		pieces pieceGroup
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
		data{
			fields: fields{
				size:   Size{2, 2},
				pieces: nil,
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{0, 0},
				},
			},
			want: ErrNoMove,
		},
		data{
			fields: fields{
				size:   Size{2, 2},
				pieces: nil,
			},
			args: args{
				move: Move{
					Start:  Position{0, 0},
					Finish: Position{1, 1},
				},
			},
			want: ErrNoPiece,
		},
		data{
			fields: fields{
				size: Size{2, 2},
				pieces: pieceGroup{
					Position{0, 0}: MockPiece{
						color:    Black,
						position: Position{0, 0},
					},
					Position{1, 1}: MockPiece{
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
		data{
			fields: fields{
				size: Size{2, 2},
				pieces: pieceGroup{
					Position{0, 0}: MockPiece{
						position: Position{0, 0},
						checkMove: func(
							move Move,
							board Board,
						) bool {
							return false
						},
					},
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
		data{
			fields: fields{
				size: Size{2, 2},
				pieces: pieceGroup{
					Position{0, 0}: MockPiece{
						position: Position{0, 0},
						checkMove: func(
							move Move,
							board Board,
						) bool {
							return true
						},
					},
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
		board := Board{
			size:   data.fields.size,
			pieces: data.fields.pieces,
		}
		got := board.CheckMove(data.args.move)

		if got != data.want {
			test.Fail()
		}
	}
}
