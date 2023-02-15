package boards

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

// BitBoard ...
type BitBoard struct {
	BaseBoard

	pieces       *bitBoardPieceGroup
	pieceFactory common.PieceFactory
}

// NewBitBoard ...
func NewBitBoard(
	size common.Size,
	pieces []common.Piece,
	pieceFactory common.PieceFactory,
) common.PieceStorage {
	pieceGroup := new(bitBoardPieceGroup)
	for _, piece := range pieces {
		pieceGroup.AddPiece(size, piece)
	}

	baseBoard := NewBaseBoard(size)
	bitBoard := BitBoard{baseBoard, pieceGroup, pieceFactory}
	return WrapBasePieceStorage(bitBoard)
}

// common.Piece ...
func (board BitBoard) Piece(position common.Position) (
	piece common.Piece,
	ok bool,
) {
	piece, _, ok =
		board.pieces.PieceByPosition(board.Size(), position, board.pieceFactory)
	return piece, ok
}

// ApplyMove ...
//
// It doesn't check that the move is correct.
func (board BitBoard) ApplyMove(move common.Move) common.PieceStorage {
	pieceGroupCopy := new(bitBoardPieceGroup)
	pieceGroupCopy.SetValue(board.pieces)

	piece, _ :=
		pieceGroupCopy.ClearPosition(board.Size(), move.Start, board.pieceFactory)
	pieceGroupCopy.ClearPosition(board.Size(), move.Finish, board.pieceFactory)

	movedPiece := piece.ApplyPosition(move.Finish)
	pieceGroupCopy.AddPiece(board.Size(), movedPiece)

	bitBoard := BitBoard{board.BaseBoard, pieceGroupCopy, board.pieceFactory}
	return WrapBasePieceStorage(bitBoard)
}
