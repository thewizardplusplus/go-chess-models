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
	// checkmate or draw
	state error
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

// State ...
//
// Checkmate or draw.
func (game Game) State() error {
	return game.state
}

// ApplyMove ...
func (game *Game) ApplyMove(
	move Move,
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
	if err != nil {
		return Move{}, err // don't wrap
	}

	return move, nil
}

// caller code should guarantee
// piece existence at the move start
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
	// here error can be
	// checkmate or draw only
	game.state = err

	return nil
}
