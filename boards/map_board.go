package boards

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

type pieceGroup map[common.Position]common.Piece

// MapBoard ...
type MapBoard struct {
	BaseBoard

	pieces pieceGroup
}

// NewMapBoard ...
func NewMapBoard(size common.Size, pieces []common.Piece) common.PieceStorage {
	pieceGroup := make(pieceGroup, len(pieces))
	for _, piece := range pieces {
		pieceGroup[piece.Position()] = piece
	}

	baseBoard := NewBaseBoard(size)
	mapBoard := MapBoard{baseBoard, pieceGroup}
	return WrapBasePieceStorage(mapBoard)
}

// common.Piece ...
func (board MapBoard) Piece(position common.Position) (
	piece common.Piece,
	ok bool,
) {
	piece, ok = board.pieces[position]
	return piece, ok
}

// ApplyMove ...
//
// It doesn't check that the move is correct.
func (board MapBoard) ApplyMove(move common.Move) common.PieceStorage {
	piece := board.pieces[move.Start]
	movedPiece := piece.ApplyPosition(move.Finish)

	pieceGroupCopy := pieceGroup{move.Finish: movedPiece}
	for position, piece := range board.pieces {
		if position != move.Start && position != move.Finish {
			pieceGroupCopy[position] = piece
		}
	}

	mapBoard := MapBoard{board.BaseBoard, pieceGroupCopy}
	return WrapBasePieceStorage(mapBoard)
}
