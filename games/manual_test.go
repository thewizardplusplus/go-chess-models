package games

import (
	"errors"
	"reflect"
	"testing"

	models "github.com/thewizardplusplus/go-chess-models"
)

func TestNewManual(test *testing.T) {
	type args struct {
		storage       models.PieceStorage
		checker       MoveSearcher
		searcher      MoveSearcher
		searcherColor models.Color
		nextColor     models.Color
	}
	type data struct {
		args     args
		wantBase bool
		wantErr  error
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
				searcher:      MockMoveSearcher{},
				searcherColor: models.White,
				nextColor:     models.White,
			},
			wantBase: true,
			wantErr:  nil,
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
				searcher:      MockMoveSearcher{},
				searcherColor: models.White,
				nextColor:     models.White,
			},
			wantBase: false,
			wantErr:  ErrCheck,
		},
	} {
		gotManual, gotErr := NewManual(
			data.args.storage,
			data.args.checker,
			data.args.searcher,
			data.args.searcherColor,
			data.args.nextColor,
		)

		hasBase := gotManual.Base != nil
		if hasBase != data.wantBase {
			test.Fail()
		}
		if gotErr != data.wantErr {
			test.Fail()
		}
	}
}

func TestManualApplyMove(test *testing.T) {
	type fields struct {
		storage       models.PieceStorage
		checker       MoveSearcher
		searcher      MoveSearcher
		searcherColor models.Color
		state         error
	}
	type args struct {
		move models.Move
	}
	type data struct {
		fields  fields
		args    args
		wantErr error
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
					checkMove: func(
						move models.Move,
					) error {
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

						return nil
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
				searcher:      MockMoveSearcher{},
				searcherColor: models.Black,
				state:         nil,
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
			wantErr: nil,
		},
		data{
			fields: fields{
				storage: MockPieceStorage{
					checkMove: func(
						move models.Move,
					) error {
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

						return errors.New("dummy")
					},
				},
				checker:       MockMoveSearcher{},
				searcher:      MockMoveSearcher{},
				searcherColor: models.Black,
				state:         nil,
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
			wantErr: errors.New("dummy"),
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
					checkMove: func(
						move models.Move,
					) error {
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

						return nil
					},
				},
				checker:       MockMoveSearcher{},
				searcher:      MockMoveSearcher{},
				searcherColor: models.White,
				state:         nil,
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
			wantErr: errors.New("opponent piece"),
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
					checkMove: func(
						move models.Move,
					) error {
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

						return nil
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
				searcher:      MockMoveSearcher{},
				searcherColor: models.Black,
				state:         nil,
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
			wantErr: ErrCheck,
		},
	} {
		manual := Manual{
			Base: &Base{
				storage: data.fields.storage,
				checker: data.fields.checker,
				state:   data.fields.state,
			},

			searcher: data.fields.searcher,
			searcherColor: data.fields.
				searcherColor,
		}
		gotErr :=
			manual.ApplyMove(data.args.move)

		if !reflect.DeepEqual(
			gotErr,
			data.wantErr,
		) {
			test.Fail()
		}
	}
}

func TestManualSearchMove(test *testing.T) {
	type fields struct {
		storage       models.PieceStorage
		checker       MoveSearcher
		searcher      MoveSearcher
		searcherColor models.Color
		state         error
	}
	type data struct {
		fields   fields
		wantMove models.Move
		wantErr  error
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
				searcher: MockMoveSearcher{
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

						move := models.Move{
							Start: models.Position{
								File: 1,
								Rank: 2,
							},
							Finish: models.Position{
								File: 3,
								Rank: 4,
							},
						}
						return move, nil
					},
				},
				searcherColor: models.Black,
				state:         nil,
			},
			wantMove: models.Move{
				Start: models.Position{
					File: 1,
					Rank: 2,
				},
				Finish: models.Position{
					File: 3,
					Rank: 4,
				},
			},
			wantErr: nil,
		},
		data{
			fields: fields{
				storage: MockPieceStorage{},
				checker: MockMoveSearcher{},
				searcher: MockMoveSearcher{
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
				searcherColor: models.Black,
				state:         nil,
			},
			wantMove: models.Move{},
			wantErr:  errors.New("dummy"),
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
				searcher: MockMoveSearcher{
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

						move := models.Move{
							Start: models.Position{
								File: 1,
								Rank: 2,
							},
							Finish: models.Position{
								File: 3,
								Rank: 4,
							},
						}
						return move, nil
					},
				},
				searcherColor: models.Black,
				state:         nil,
			},
			wantMove: models.Move{},
			wantErr:  ErrCheck,
		},
	} {
		manual := Manual{
			Base: &Base{
				storage: data.fields.storage,
				checker: data.fields.checker,
				state:   data.fields.state,
			},

			searcher: data.fields.searcher,
			searcherColor: data.fields.
				searcherColor,
		}
		gotMove, gotErr := manual.SearchMove()

		if !reflect.DeepEqual(
			gotMove,
			data.wantMove,
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
