package boards

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

// SliceBoard ...
type SliceBoard struct {
	BaseBoard

	pieces []common.Piece
}

// NewSliceBoard ...
func NewSliceBoard(
	size common.Size,
	pieces []common.Piece,
) common.PieceStorage {
	extendedPieces := make([]common.Piece, size.PositionCount())
	for _, piece := range pieces {
		positionIndex := size.PositionIndex(piece.Position())
		extendedPieces[positionIndex] = piece
	}

	baseBoard := NewBaseBoard(size)
	sliceBoard := SliceBoard{baseBoard, extendedPieces}
	return WrapBasePieceStorage(sliceBoard)
}

// common.Piece ...
func (board SliceBoard) Piece(position common.Position) (
	piece common.Piece,
	ok bool,
) {
	positionIndex := board.Size().PositionIndex(position)
	piece = board.pieces[positionIndex]
	return piece, piece != nil
}

// ApplyMove ...
//
// It doesn't check that the move is correct.
func (board SliceBoard) ApplyMove(move common.Move) common.PieceStorage {
	pieceGroupCopy := make([]common.Piece, len(board.pieces))
	copy(pieceGroupCopy, board.pieces)

	startPositionIndex := board.Size().PositionIndex(move.Start)
	piece := pieceGroupCopy[startPositionIndex]
	pieceGroupCopy[startPositionIndex] = nil

	finishPositionIndex := board.Size().PositionIndex(move.Finish)
	movedPiece := piece.ApplyPosition(move.Finish)
	pieceGroupCopy[finishPositionIndex] = movedPiece

	sliceBoard := SliceBoard{board.BaseBoard, pieceGroupCopy}
	return WrapBasePieceStorage(sliceBoard)
}
