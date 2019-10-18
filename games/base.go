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
	// It should return only following errors:
	// * models.ErrKingCapture;
	// * ErrCheckmate;
	// * ErrDraw.
	SearchMove(
		storage models.PieceStorage,
		color models.Color,
	) (models.Move, error)
}

// Base ...
type Base struct {
	storage models.PieceStorage
	checker MoveSearcher
	// ErrCheckmate or ErrDraw
	state error
}

// NewBase ...
//
// After this call you should check
// a state of the game.
func NewBase(
	storage models.PieceStorage,
	checker MoveSearcher,
	nextColor models.Color,
) (*Base, error) {
	game := &Base{checker: checker}
	err := game.tryUpdateStorage(
		storage,
		nextColor,
	)
	if err != nil {
		return nil, err // don't wrap
	}

	return game, nil
}

// Storage ...
func (
	game Base,
) Storage() models.PieceStorage {
	return game.storage
}

// State ...
//
// ErrCheckmate or ErrDraw.
func (game Base) State() error {
	return game.state
}

// ApplyMove ...
//
// It DOESN'T check that the move is correct.
//
// It DOES check of storage state
// after the move.
//
// After this call you should check
// a state of the game.
func (game *Base) ApplyMove(
	move models.Move,
) error {
	nextStorage :=
		game.storage.ApplyMove(move)
	nextColor :=
		game.moveColor(move).Negative()
	return game.tryUpdateStorage(
		nextStorage,
		nextColor,
	)
}

// caller code should guarantee
// piece existence at the move start
func (game Base) moveColor(
	move models.Move,
) models.Color {
	piece, _ := game.storage.Piece(move.Start)
	return piece.Color()
}

// it checks storage state before update
func (game *Base) tryUpdateStorage(
	nextStorage models.PieceStorage,
	nextColor models.Color,
) error {
	_, err := game.checker.SearchMove(
		nextStorage,
		nextColor,
	)
	if err == models.ErrKingCapture {
		return ErrCheck // don't wrap
	}

	game.storage = nextStorage
	// here error can be only
	// ErrCheckmate or ErrDraw
	game.state = err

	return nil
}
