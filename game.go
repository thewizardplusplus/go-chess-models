package chessmodels

import (
	"errors"
)

// ...
var (
	ErrCheck = errors.New("check")
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
) Game {
	return Game{
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
	if err == ErrKingCapture {
		return ErrCheck
	}

	game.storage = storage
	// here error can be checkmate or draw
	// only
	return err
}

// SearchMove ...
func (game Game) SearchMove() (
	Move,
	error,
) {
	move, err := game.searcher.SearchMove(
		game.storage,
		game.searcherColor,
	)
	if err != nil {
		return Move{}, err // don't wrap
	}

	game.storage =
		game.storage.ApplyMove(move)
	userColor := game.searcherColor.Negative()
	_, err = game.checker.SearchMove(
		game.storage,
		userColor,
	)

	// here error can be checkmate or draw
	// only
	return move, err
}
