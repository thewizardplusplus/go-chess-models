package chessmodels

import (
	"errors"
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

type MockPiece struct {
	kind     common.Kind
	color    common.Color
	position common.Position

	checkMove func(move common.Move, storage common.PieceStorage) bool
}

func (piece MockPiece) Kind() common.Kind {
	return piece.kind
}

func (piece MockPiece) Color() common.Color {
	return piece.color
}

func (piece MockPiece) Position() common.Position {
	return piece.position
}

func (piece MockPiece) ApplyPosition(position common.Position) common.Piece {
	return MockPiece{
		kind:      piece.kind,
		color:     piece.color,
		position:  position,
		checkMove: piece.checkMove,
	}
}

func (piece MockPiece) CheckMove(
	move common.Move,
	storage common.PieceStorage,
) bool {
	if piece.checkMove == nil {
		panic("not implemented")
	}

	return piece.checkMove(move, storage)
}

type MockBasePieceStorage struct {
	size common.Size

	piece func(position common.Position) (piece common.Piece, ok bool)
}

func (storage MockBasePieceStorage) Size() common.Size {
	return storage.size
}

func (storage MockBasePieceStorage) Piece(position common.Position) (
	piece common.Piece,
	ok bool,
) {
	if storage.piece == nil {
		panic("not implemented")
	}

	return storage.piece(position)
}

func (storage MockBasePieceStorage) ApplyMove(
	move common.Move,
) common.PieceStorage {
	panic("not implemented")
}

type MockPieceGroupGetter struct {
	pieces []common.Piece
}

func (pieceGroupGetter MockPieceGroupGetter) Pieces() []common.Piece {
	return pieceGroupGetter.pieces
}

type MockMoveChecker struct {
	checkMove func(move common.Move) error
}

func (moveChecker MockMoveChecker) CheckMove(move common.Move) error {
	if moveChecker.checkMove == nil {
		panic("not implemented")
	}

	return moveChecker.checkMove(move)
}

type MockPieceStorage struct {
	MockBasePieceStorage
	MockPieceGroupGetter
	MockMoveChecker
}

func TestMoveCheckerMovesForColor(test *testing.T) {
	type fields struct {
		size      common.Size
		pieces    []common.Piece
		checkMove func(move common.Move) error
	}
	type args struct {
		color common.Color
	}
	type data struct {
		fields    fields
		args      args
		wantMoves []common.Move
		wantErr   error
	}

	for _, data := range []data{
		{
			fields: fields{
				size: common.Size{2, 2},
				pieces: []common.Piece{
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 0},
					},
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 1},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 0},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 1},
					},
				},
				checkMove: func(move common.Move) error {
					return nil
				},
			},
			args: args{
				color: common.Black,
			},
			wantMoves: []common.Move{
				{
					Start:  common.Position{0, 0},
					Finish: common.Position{0, 0},
				},
				{
					Start:  common.Position{0, 0},
					Finish: common.Position{1, 0},
				},
				{
					Start:  common.Position{0, 0},
					Finish: common.Position{0, 1},
				},
				{
					Start:  common.Position{0, 0},
					Finish: common.Position{1, 1},
				},

				{
					Start:  common.Position{0, 1},
					Finish: common.Position{0, 0},
				},
				{
					Start:  common.Position{0, 1},
					Finish: common.Position{1, 0},
				},
				{
					Start:  common.Position{0, 1},
					Finish: common.Position{0, 1},
				},
				{
					Start:  common.Position{0, 1},
					Finish: common.Position{1, 1},
				},
			},
			wantErr: nil,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				pieces: []common.Piece{
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 0},
					},
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 1},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 0},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 1},
					},
				},
				checkMove: func(move common.Move) error {
					return nil
				},
			},
			args: args{
				color: common.White,
			},
			wantMoves: []common.Move{
				{
					Start:  common.Position{1, 0},
					Finish: common.Position{0, 0},
				},
				{
					Start:  common.Position{1, 0},
					Finish: common.Position{1, 0},
				},
				{
					Start:  common.Position{1, 0},
					Finish: common.Position{0, 1},
				},
				{
					Start:  common.Position{1, 0},
					Finish: common.Position{1, 1},
				},

				{
					Start:  common.Position{1, 1},
					Finish: common.Position{0, 0},
				},
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{1, 0},
				},
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{0, 1},
				},
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{1, 1},
				},
			},
			wantErr: nil,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				pieces: []common.Piece{
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 0},
					},
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 1},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 0},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 1},
					},
				},
				checkMove: func(move common.Move) error {
					return errors.New("dummy")
				},
			},
			args: args{
				color: common.Black,
			},
			wantMoves: nil,
			wantErr:   nil,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				pieces: []common.Piece{
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 0},
					},
					MockPiece{
						color:    common.Black,
						position: common.Position{0, 1},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 0},
					},
					MockPiece{
						color:    common.White,
						position: common.Position{1, 1},
					},
				},
				checkMove: func(move common.Move) error {
					return common.ErrKingCapture
				},
			},
			args: args{
				color: common.Black,
			},
			wantMoves: nil,
			wantErr:   common.ErrKingCapture,
		},
	} {
		storage := MockPieceStorage{
			MockBasePieceStorage: MockBasePieceStorage{
				size: data.fields.size,
			},
			MockPieceGroupGetter: MockPieceGroupGetter{
				pieces: data.fields.pieces,
			},
			MockMoveChecker: MockMoveChecker{
				checkMove: data.fields.checkMove,
			},
		}

		var generator MoveGenerator
		gotMoves, gotErr := generator.MovesForColor(storage, data.args.color)

		if !reflect.DeepEqual(gotMoves, data.wantMoves) {
			test.Fail()
		}
		if !reflect.DeepEqual(gotErr, data.wantErr) {
			test.Fail()
		}
	}
}

func TestMoveCheckerMovesForPosition(test *testing.T) {
	type fields struct {
		size      common.Size
		checkMove func(move common.Move) error
	}
	type args struct {
		position common.Position
	}
	type data struct {
		fields    fields
		args      args
		wantMoves []common.Move
		wantErr   error
	}

	for _, data := range []data{
		{
			fields: fields{
				size: common.Size{2, 2},
				checkMove: func(move common.Move) error {
					return errors.New("dummy")
				},
			},
			args: args{
				position: common.Position{1, 1},
			},
			wantMoves: nil,
			wantErr:   nil,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				checkMove: func(move common.Move) error {
					return common.ErrKingCapture
				},
			},
			args: args{
				position: common.Position{1, 1},
			},
			wantMoves: nil,
			wantErr:   common.ErrKingCapture,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				checkMove: func(move common.Move) error {
					return nil
				},
			},
			args: args{
				position: common.Position{1, 1},
			},
			wantMoves: []common.Move{
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{0, 0},
				},
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{1, 0},
				},
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{0, 1},
				},
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{1, 1},
				},
			},
			wantErr: nil,
		},
		{
			fields: fields{
				size: common.Size{2, 2},
				checkMove: func(move common.Move) error {
					if move.Finish.Rank == 1 {
						return errors.New("dummy")
					}

					return nil
				},
			},
			args: args{
				position: common.Position{1, 1},
			},
			wantMoves: []common.Move{
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{0, 0},
				},
				{
					Start:  common.Position{1, 1},
					Finish: common.Position{1, 0},
				},
			},
			wantErr: nil,
		},
	} {
		storage := MockPieceStorage{
			MockBasePieceStorage: MockBasePieceStorage{
				size: data.fields.size,
			},
			MockMoveChecker: MockMoveChecker{
				checkMove: data.fields.checkMove,
			},
		}

		var generator MoveGenerator
		gotMoves, gotErr := generator.MovesForPosition(storage, data.args.position)

		if !reflect.DeepEqual(gotMoves, data.wantMoves) {
			test.Fail()
		}
		if !reflect.DeepEqual(gotErr, data.wantErr) {
			test.Fail()
		}
	}
}
