package chessmodels

import (
	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
)

// SliceBoard ...
type SliceBoard struct {
	boards.BaseBoard

	pieces []common.Piece
}

// NewSliceBoard ...
func NewSliceBoard(size common.Size, pieces []common.Piece) common.PieceStorage {
	extendedPieces := make([]common.Piece, size.PositionCount())
	for _, piece := range pieces {
		extendedPieces[size.PositionIndex(piece.Position())] = piece
	}

	baseBoard := boards.NewBaseBoard(size)
	sliceBoard := SliceBoard{baseBoard, extendedPieces}
	return boards.DefaultBoardWrapper{sliceBoard}
}

// common.Piece ...
func (board SliceBoard) Piece(position common.Position) (piece common.Piece, ok bool) {
	piece = board.pieces[board.Size().PositionIndex(position)]
	return piece, piece != nil
}

// ApplyMove ...
//
// It doesn't check that the move is correct.
func (board SliceBoard) ApplyMove(move common.Move) common.PieceStorage {
	pieceGroupCopy := make([]common.Piece, len(board.pieces))
	copy(pieceGroupCopy, board.pieces)

	startIndex, finishIndex :=
		board.Size().PositionIndex(move.Start), board.Size().PositionIndex(move.Finish)
	piece := pieceGroupCopy[startIndex]
	pieceGroupCopy[startIndex] = nil

	movedPiece := piece.ApplyPosition(move.Finish)
	pieceGroupCopy[finishIndex] = movedPiece

	sliceBoard := SliceBoard{board.BaseBoard, pieceGroupCopy}
	return boards.DefaultBoardWrapper{sliceBoard}
}
