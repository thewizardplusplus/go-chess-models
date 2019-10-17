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
		args      args
		wantBase  bool
		wantState error
		wantErr   error
	}

	for _, data := range []data{} {
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
