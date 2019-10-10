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
	userColor     Color
	checker       MoveSearcher
}

// NewGame ...
func NewGame(
	storage PieceStorage,
	searcher MoveSearcher,
	searcherColor Color,
	checker MoveSearcher,
) *Game {
	userColor := searcherColor.Negative()
	return &Game{
		storage:       storage,
		searcher:      searcher,
		searcherColor: searcherColor,
		userColor:     userColor,
		checker:       checker,
	}
}

// Storage ...
func (game Game) Storage() PieceStorage {
	return game.storage
}

// ApplyMove ...
func (game *Game) ApplyMove(
	move Move,
) error {
	err := game.storage.CheckMove(move)
	if err != nil {
		return err // don't wrap
	}

	color := game.moveColor(move)
	if color != game.userColor {
		return errors.New("opponent piece")
	}

	return game.tryApplyMove(move)
}

// SearchMove ...
func (game *Game) SearchMove() (
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

	err = game.tryApplyMove(move)
	return move, err
}

func (game Game) moveColor(
	move Move,
) Color {
	piece, _ := game.storage.Piece(move.Start)
	return piece.Color()
}

func (game *Game) tryApplyMove(
	move Move,
) error {
	color := game.moveColor(move)
	storage := game.storage.ApplyMove(move)
	_, err :=
		game.checker.SearchMove(storage, color)
	if err == ErrKingCapture {
		return ErrCheck
	}

	game.storage = storage
	// here error can be checkmate or draw
	// only
	return err
}
