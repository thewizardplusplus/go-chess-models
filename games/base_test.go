package games

import (
	"errors"
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

type MockPieceStorage struct {
	piece func(
		position models.Position,
	) (piece models.Piece, ok bool)
	applyMove func(
		move models.Move,
	) models.PieceStorage
	checkMove func(move models.Move) error
}

func (
	storage MockPieceStorage,
) Size() models.Size {
	panic("not implemented")
}

func (
	storage MockPieceStorage,
) Piece(
	position models.Position,
) (piece models.Piece, ok bool) {
	if storage.piece == nil {
		panic("not implemented")
	}

	return storage.piece(position)
}

func (
	storage MockPieceStorage,
) Pieces() []models.Piece {
	panic("not implemented")
}

func (storage MockPieceStorage) ApplyMove(
	move models.Move,
) models.PieceStorage {
	if storage.applyMove == nil {
		panic("not implemented")
	}

	return storage.applyMove(move)
}

func (storage MockPieceStorage) CheckMove(
	move models.Move,
) error {
	if storage.checkMove == nil {
		panic("not implemented")
	}

	return storage.checkMove(move)
}

type MockMoveSearcher struct {
	searchMove func(
		storage models.PieceStorage,
		color models.Color,
	) (models.Move, error)
}

func (searcher MockMoveSearcher) SearchMove(
	storage models.PieceStorage,
	color models.Color,
) (models.Move, error) {
	if searcher.searchMove == nil {
		panic("not implemented")
	}

	return searcher.searchMove(storage, color)
}

type MockPiece struct {
	color models.Color
}

func (piece MockPiece) Kind() models.Kind {
	panic("not implemented")
}

func (
	piece MockPiece,
) Color() models.Color {
	return piece.color
}

func (
	piece MockPiece,
) Position() models.Position {
	panic("not implemented")
}

func (piece MockPiece) ApplyPosition(
	position models.Position,
) models.Piece {
	panic("not implemented")
}

func (piece MockPiece) CheckMove(
	move models.Move,
	storage models.PieceStorage,
) bool {
	panic("not implemented")
}

func TestNewBase(test *testing.T) {
	type args struct {
		storage   models.PieceStorage
		checker   MoveSearcher
		nextColor models.Color
	}
	type data struct {
		args      args
		wantBase  bool
		wantState error
		wantErr   error
	}

	for _, data := range []data{
		data{
			args: args{
				storage: MockPieceStorage{},
				checker: MockMoveSearcher{
					searchMove: func(
						storage models.PieceStorage,
						color models.Color,
					) (models.Move, error) {
						_, ok :=
							storage.(MockPieceStorage)
						if !ok {
							test.Fail()
						}
						if color != models.White {
							test.Fail()
						}

						return models.Move{}, nil
					},
				},
				nextColor: models.White,
			},
			wantBase:  true,
			wantState: nil,
			wantErr:   nil,
		},
		data{
			args: args{
				storage: MockPieceStorage{},
				checker: MockMoveSearcher{
					searchMove: func(
						storage models.PieceStorage,
						color models.Color,
					) (models.Move, error) {
						_, ok :=
							storage.(MockPieceStorage)
						if !ok {
							test.Fail()
						}
						if color != models.White {
							test.Fail()
						}

						return models.Move{},
							errors.New("dummy")
					},
				},
				nextColor: models.White,
			},
			wantBase:  true,
			wantState: errors.New("dummy"),
			wantErr:   nil,
		},
		data{
			args: args{
				storage: MockPieceStorage{},
				checker: MockMoveSearcher{
					searchMove: func(
						storage models.PieceStorage,
						color models.Color,
					) (models.Move, error) {
						_, ok :=
							storage.(MockPieceStorage)
						if !ok {
							test.Fail()
						}
						if color != models.White {
							test.Fail()
						}

						return models.Move{},
							models.ErrKingCapture
					},
				},
				nextColor: models.White,
			},
			wantBase:  false,
			wantState: nil,
			wantErr:   ErrCheck,
		},
	} {
		gotBase, gotErr := NewBase(
			data.args.storage,
			data.args.checker,
			data.args.nextColor,
		)

		hasBase := gotBase != nil
		if hasBase != data.wantBase {
			test.Fail()
		}
		if gotErr != data.wantErr {
			test.Fail()
		}
		if data.wantBase == false {
			continue
		}

		if !reflect.DeepEqual(
			gotBase.storage,
			data.args.storage,
		) {
			test.Fail()
		}

		_, ok := gotBase.
			checker.(MockMoveSearcher)
		if !ok {
			test.Fail()
		}
		if !reflect.DeepEqual(
			gotBase.state,
			data.wantState,
		) {
			test.Fail()
		}
	}
}

func TestBaseStorage(test *testing.T) {
	var storage MockPieceStorage
	base := Base{
		storage: storage,
	}
	got := base.Storage()

	if !reflect.DeepEqual(got, storage) {
		test.Fail()
	}
}

func TestBaseState(test *testing.T) {
	state := errors.New("dummy")
	base := Base{
		state: state,
	}
	got := base.State()

	if got != state {
		test.Fail()
	}
}

func TestBaseApplyMove(test *testing.T) {
	type fields struct {
		storage models.PieceStorage
		checker MoveSearcher
		state   error
	}
	type args struct {
		move models.Move
	}
	type data struct {
		fields    fields
		args      args
		wantState error
		wantErr   error
	}

	for _, data := range []data{
		data{
			fields: fields{
				storage: MockPieceStorage{
					piece: func(
						position models.Position,
					) (piece models.Piece, ok bool) {
						expectedPosition :=
							models.Position{
								File: 1,
								Rank: 2,
							}
						if position !=
							expectedPosition {
							test.Fail()
						}

						piece = MockPiece{
							color: models.White,
						}
						return piece, true
					},
					applyMove: func(
						move models.Move,
					) models.PieceStorage {
						expectedMove := models.Move{
							Start: models.Position{
								File: 1,
								Rank: 2,
							},
							Finish: models.Position{
								File: 3,
								Rank: 4,
							},
						}
						if move != expectedMove {
							test.Fail()
						}

						return MockPieceStorage{}
					},
				},
				checker: MockMoveSearcher{
					searchMove: func(
						storage models.PieceStorage,
						color models.Color,
					) (models.Move, error) {
						_, ok :=
							storage.(MockPieceStorage)
						if !ok {
							test.Fail()
						}
						if color != models.Black {
							test.Fail()
						}

						return models.Move{}, nil
					},
				},
				state: nil,
			},
			args: args{
				move: models.Move{
					Start: models.Position{
						File: 1,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 4,
					},
				},
			},
			wantState: nil,
			wantErr:   nil,
		},
		data{
			fields: fields{
				storage: MockPieceStorage{
					piece: func(
						position models.Position,
					) (piece models.Piece, ok bool) {
						expectedPosition :=
							models.Position{
								File: 1,
								Rank: 2,
							}
						if position !=
							expectedPosition {
							test.Fail()
						}

						piece = MockPiece{
							color: models.White,
						}
						return piece, true
					},
					applyMove: func(
						move models.Move,
					) models.PieceStorage {
						expectedMove := models.Move{
							Start: models.Position{
								File: 1,
								Rank: 2,
							},
							Finish: models.Position{
								File: 3,
								Rank: 4,
							},
						}
						if move != expectedMove {
							test.Fail()
						}

						return MockPieceStorage{}
					},
				},
				checker: MockMoveSearcher{
					searchMove: func(
						storage models.PieceStorage,
						color models.Color,
					) (models.Move, error) {
						_, ok :=
							storage.(MockPieceStorage)
						if !ok {
							test.Fail()
						}
						if color != models.Black {
							test.Fail()
						}

						return models.Move{},
							errors.New("dummy")
					},
				},
				state: nil,
			},
			args: args{
				move: models.Move{
					Start: models.Position{
						File: 1,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 4,
					},
				},
			},
			wantState: errors.New("dummy"),
			wantErr:   nil,
		},
		data{
			fields: fields{
				storage: MockPieceStorage{
					piece: func(
						position models.Position,
					) (piece models.Piece, ok bool) {
						expectedPosition :=
							models.Position{
								File: 1,
								Rank: 2,
							}
						if position !=
							expectedPosition {
							test.Fail()
						}

						piece = MockPiece{
							color: models.White,
						}
						return piece, true
					},
					applyMove: func(
						move models.Move,
					) models.PieceStorage {
						expectedMove := models.Move{
							Start: models.Position{
								File: 1,
								Rank: 2,
							},
							Finish: models.Position{
								File: 3,
								Rank: 4,
							},
						}
						if move != expectedMove {
							test.Fail()
						}

						return MockPieceStorage{}
					},
				},
				checker: MockMoveSearcher{
					searchMove: func(
						storage models.PieceStorage,
						color models.Color,
					) (models.Move, error) {
						_, ok :=
							storage.(MockPieceStorage)
						if !ok {
							test.Fail()
						}
						if color != models.Black {
							test.Fail()
						}

						return models.Move{},
							models.ErrKingCapture
					},
				},
				state: nil,
			},
			args: args{
				move: models.Move{
					Start: models.Position{
						File: 1,
						Rank: 2,
					},
					Finish: models.Position{
						File: 3,
						Rank: 4,
					},
				},
			},
			wantState: nil,
			wantErr:   ErrCheck,
		},
	} {
		base := Base{
			storage: data.fields.storage,
			checker: data.fields.checker,
			state:   data.fields.state,
		}
		gotErr := base.ApplyMove(data.args.move)

		_, ok := base.
			storage.(MockPieceStorage)
		if !ok {
			test.Fail()
		}
		if !reflect.DeepEqual(
			base.state,
			data.wantState,
		) {
			test.Fail()
		}
		if !reflect.DeepEqual(
			gotErr,
			data.wantErr,
		) {
			test.Fail()
		}
	}
}
