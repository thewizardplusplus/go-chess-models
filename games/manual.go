package games

import (
	"errors"

	models "github.com/thewizardplusplus/go-chess-models"
)

// Manual ...
type Manual struct {
	*Base

	searcher      MoveSearcher
	searcherColor models.Color
}

// NewManual ...
func NewManual(
	storage models.PieceStorage,
	checker MoveSearcher,
	searcher MoveSearcher,
	searcherColor models.Color,
	nextColor models.Color,
) (Manual, error) {
	base, err :=
		NewBase(storage, checker, nextColor)
	if err != nil {
		return Manual{}, err // don't wrap
	}

	game := Manual{
		Base: base,

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
func (game Manual) ApplyMove(
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

	return game.Base.ApplyMove(move)
}

// SearchMove ...
func (game Manual) SearchMove() (
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

	err = game.Base.ApplyMove(move)
	if err != nil {
		return models.Move{}, err // don't wrap
	}

	return move, nil
}
