package chessmodels

import (
	"reflect"
	"testing"
)

type MockPiece struct {
	kind     Kind
	color    Color
	position Position

	checkMove func(move Move, storage PieceStorage) bool
}

func (piece MockPiece) Kind() Kind {
	return piece.kind
}

func (piece MockPiece) Color() Color {
	return piece.color
}

func (piece MockPiece) Position() Position {
	return piece.position
}

func (piece MockPiece) ApplyPosition(position Position) Piece {
	return MockPiece{
		kind:      piece.kind,
		color:     piece.color,
		position:  position,
		checkMove: piece.checkMove,
	}
}

func (piece MockPiece) CheckMove(move Move, storage PieceStorage) bool {
	if piece.checkMove == nil {
		panic("not implemented")
	}

	return piece.checkMove(move, storage)
}

func TestNewMapBoard(test *testing.T) {
	board := NewMapBoard(Size{5, 5}, []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})

	expectedBoard := DefaultBoardWrapper{
		BasePieceStorage: MapBoard{
			BaseBoard: BaseBoard{
				size: Size{5, 5},
			},

			pieces: pieceGroup{
				Position{2, 3}: MockPiece{
					position: Position{2, 3},
				},
				Position{4, 2}: MockPiece{
					position: Position{4, 2},
				},
			},
		},
	}
	if !reflect.DeepEqual(board, expectedBoard) {
		test.Fail()
	}
}

func TestMapBoardPiece(test *testing.T) {
	type fields struct {
		size   Size
		pieces pieceGroup
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
				pieces: pieceGroup{
					Position{2, 3}: MockPiece{
						position: Position{2, 3},
					},
					Position{4, 2}: MockPiece{
						position: Position{4, 2},
					},
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
				pieces: pieceGroup{
					Position{2, 3}: MockPiece{
						position: Position{2, 3},
					},
					Position{4, 2}: MockPiece{
						position: Position{4, 2},
					},
				},
			},
			args:      args{Position{0, 0}},
			wantPiece: nil,
			wantOk:    false,
		},
	} {
		board := MapBoard{
			BaseBoard: BaseBoard{
				size: data.fields.size,
			},

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

func TestMapBoardApplyMove(test *testing.T) {
	board := NewMapBoard(Size{5, 5}, []Piece{
		MockPiece{position: Position{2, 3}},
		MockPiece{position: Position{4, 2}},
	})
	nextBoard := board.ApplyMove(Move{
		Start:  Position{4, 2},
		Finish: Position{1, 2},
	})

	expectedBoard := DefaultBoardWrapper{
		BasePieceStorage: MapBoard{
			BaseBoard: BaseBoard{
				size: Size{5, 5},
			},

			pieces: pieceGroup{
				Position{2, 3}: MockPiece{
					position: Position{2, 3},
				},
				Position{4, 2}: MockPiece{
					position: Position{4, 2},
				},
			},
		},
	}
	if !reflect.DeepEqual(board, expectedBoard) {
		test.Fail()
	}

	expectedNextBoard := DefaultBoardWrapper{
		BasePieceStorage: MapBoard{
			BaseBoard: BaseBoard{
				size: Size{5, 5},
			},

			pieces: pieceGroup{
				Position{1, 2}: MockPiece{
					position: Position{1, 2},
				},
				Position{2, 3}: MockPiece{
					position: Position{2, 3},
				},
			},
		},
	}
	if !reflect.DeepEqual(nextBoard, expectedNextBoard) {
		test.Fail()
	}
}
