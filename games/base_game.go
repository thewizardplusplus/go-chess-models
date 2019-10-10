package games

import (
	"errors"

	models "github.com/thewizardplusplus/go-chess-models"
)

// ...
var (
	ErrCheck     = errors.New("check")
	ErrCheckmate = errors.New("checkmate")
	ErrDraw      = errors.New("draw")
)

// MoveSearcher ...
type MoveSearcher interface {
	SearchMove(
		storage models.PieceStorage,
		color models.Color,
	) (models.Move, error)
}

// BaseGame ...
type BaseGame struct {
	storage models.PieceStorage
	checker MoveSearcher
	// ErrCheckmate or ErrDraw
	state error
}

// NewBaseGame ...
func NewBaseGame(
	storage models.PieceStorage,
	checker MoveSearcher,
) *BaseGame {
	return &BaseGame{
		storage: storage,
		checker: checker,
	}
}

// Storage ...
func (
	game BaseGame,
) Storage() models.PieceStorage {
	return game.storage
}

// State ...
//
// ErrCheckmate or ErrDraw.
func (game BaseGame) State() error {
	return game.state
}

// ApplyMove ...
//
// It DOESN'T check that the move is correct.
//
// It DOES check of storage state
// after the move.
func (game *BaseGame) ApplyMove(
	move models.Move,
) error {
	// disable move if the game already is
	// in ErrCheckmate or ErrDraw states
	if game.state != nil {
		return game.state // don't wrap
	}

	nextStorage := game.storage.ApplyMove(move)
	nextColor :=
		game.moveColor(move).Negative()
	_, err := game.checker.SearchMove(
		nextStorage,
		nextColor,
	)
	if err == models.ErrKingCapture {
		return ErrCheck
	}

	game.storage = nextStorage
	// here error can be only
	// ErrCheckmate or ErrDraw
	game.state = err

	return nil
}

// caller code should guarantee
// piece existence at the move start
func (game BaseGame) moveColor(
	move models.Move,
) models.Color {
	piece, _ := game.storage.Piece(move.Start)
	return piece.Color()
}