package games

import (
	"errors"

	models "github.com/thewizardplusplus/go-chess-models"
)

// ManualGame ...
type ManualGame struct {
	*BaseGame

	searcher      MoveSearcher
	searcherColor models.Color
}

// NewManualGame ...
func NewManualGame(
	storage models.PieceStorage,
	checker MoveSearcher,
	searcher MoveSearcher,
	searcherColor models.Color,
	nextColor models.Color,
) (*ManualGame, error) {
	baseGame, err :=
		NewBaseGame(storage, checker, nextColor)
	if err != nil {
		return nil, err // don't wrap
	}

	game := &ManualGame{
		BaseGame: baseGame,

		searcher:      searcher,
		searcherColor: searcherColor,
	}
	return game, nil
}

// ApplyMove ...
//
// It DOES check that the move is correct.
//
// It DOES check of storage state
// after the move.
func (game *ManualGame) ApplyMove(
	move models.Move,
) error {
	// disable move if the game already is
	// in ErrCheckmate or ErrDraw states
	if game.state != nil {
		return game.state // don't wrap
	}

	err := game.storage.CheckMove(move)
	if err != nil {
		return err // don't wrap
	}

	moveColor := game.moveColor(move)
	userColor := game.searcherColor.Negative()
	if moveColor != userColor {
		return errors.New("opponent piece")
	}

	return game.BaseGame.ApplyMove(move)
}

// SearchMove ...
func (game *ManualGame) SearchMove() (
	models.Move,
	error,
) {
	// disable move if the game already is
	// in ErrCheckmate or ErrDraw states
	if game.state != nil {
		return models.Move{},
			game.state // don't wrap
	}

	move, err := game.searcher.SearchMove(
		game.storage,
		game.searcherColor,
	)
	if err != nil {
		return models.Move{}, err // don't wrap
	}

	err = game.BaseGame.ApplyMove(move)
	if err != nil {
		return models.Move{}, err // don't wrap
	}

	return move, nil
}
