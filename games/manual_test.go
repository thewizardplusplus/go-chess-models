package games

import (
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
