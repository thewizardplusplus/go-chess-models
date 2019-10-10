package chessmodels

import (
	"errors"
)

// MoveSearcher ...
type MoveSearcher interface {
	SearchMove(
		storage PieceStorage,
		color Color,
	) (Move, error)
}

// Game ...
type Game struct {
	storage       PieceStorage
	searcher      MoveSearcher
	searcherColor Color
	checker       MoveSearcher
}

// NewGame ...
func NewGame(
	storage PieceStorage,
	searcher MoveSearcher,
	searcherColor Color,
	checker MoveSearcher,
) *Game {
	return &Game{
		storage:       storage,
		searcher:      searcher,
		searcherColor: searcherColor,
		checker:       checker,
	}
}

// Storage ...
func (game Game) Storage() PieceStorage {
	return game.storage
}

// ApplyMove ...
func (game Game) ApplyMove(
	move Move,
) error {
	err := game.storage.CheckMove(move)
	if err != nil {
		return err // don't wrap
	}

	// CheckMove call guarantees no error here
	piece, _ := game.storage.Piece(move.Start)
	userColor := game.searcherColor.Negative()
	if piece.Color() != userColor {
		return errors.New("opponent piece")
	}

	storage := game.storage.ApplyMove(move)
	_, err = game.checker.SearchMove(
		storage,
		game.searcherColor,
	)
	if err != nil {
		return err // don't wrap
	}

	game.storage = storage
	return nil
}
