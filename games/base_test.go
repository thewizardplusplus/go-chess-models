package games

import (
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
