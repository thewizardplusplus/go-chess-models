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
//
// After this call you should check
// a state of the game.
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
//
// After this call you should check
// a state of the game.
func (game Manual) ApplyMove(
	move models.Move,
) error {
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
//
// After this call you should check
// a state of the game.
func (game Manual) SearchMove() models.Move {
	// thanks to the rest code
	// this call can't return any error
	move, _ := game.searcher.SearchMove(
		game.storage,
		game.searcherColor,
	)

	// thanks to the rest code
	// this call can't return any error
	game.Base.ApplyMove(move)

	return move
}
