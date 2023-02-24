package chessmodels

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

// MoveGenerator ...
type MoveGenerator struct{}

// MovesForColor ...
//
// It doesn't guarantee an order of returned moves.
//
// It doesn't take into account possible checks and can generate such moves.
//
// It returns an error only on a king capture.
func (generator MoveGenerator) MovesForColor(
	storage common.PieceStorage,
	color common.Color,
) ([]common.Move, error) {
	var moves []common.Move
	for _, piece := range storage.Pieces() {
		if piece.Color() != color {
			continue
		}

		positionMoves, err :=
			generator.MovesForPosition(storage, piece.Position())
		if err != nil {
			return nil, err
		}

		moves = append(moves, positionMoves...)
	}

	return moves, nil
}

// MovesForPosition ...
//
// It doesn't take into account possible checks and can generate such moves.
//
// It returns an error only on a king capture.
func (generator MoveGenerator) MovesForPosition(
	storage common.PieceStorage,
	position common.Position,
) ([]common.Move, error) {
	var moves []common.Move
	if err := storage.Size().IteratePositions(func(finish common.Position) error {
		move := common.Move{position, finish}
		if err := storage.CheckMove(move); err != nil {
			// if the move captures a king, break a generating
			if err == common.ErrKingCapture {
				return err
			}

			// on other errors just skip this move
			return nil
		}

		moves = append(moves, move)
		return nil
	}); err != nil {
		return nil, err
	}

	return moves, nil
}

// PerftMoveGenerator ...
type PerftMoveGenerator interface {
	MovesForColor(storage common.PieceStorage, color common.Color) (
		[]common.Move,
		error,
	)
}

// PerftHandler ...
type PerftHandler func(move common.Move, count int, deep int)

// Perft ...
func Perft(
	generator PerftMoveGenerator,
	storage common.PieceStorage,
	color common.Color,
	deep int,
	handler PerftHandler,
) int {
	// check for a check should be first, including before a termination check,
	// because a terminated evaluation doesn't make sense for a check position
	moves, err := generator.MovesForColor(storage, color)
	if err != nil {
		return 0
	}

	if deep == 0 {
		return 1
	}

	var totalMoveCount int
	for _, move := range moves {
		nextStorage := storage.ApplyMove(move)
		nextColor := color.Negative()
		moveCount := Perft(
			generator,
			nextStorage,
			nextColor,
			deep-1,
			handler,
		)
		if handler != nil {
			handler(move, moveCount, deep)
		}

		totalMoveCount += moveCount
	}

	return totalMoveCount
}
