package chessmodels

// SliceBoard ...
type SliceBoard struct {
	BaseBoard

	pieces []Piece
}

// NewSliceBoard ...
func NewSliceBoard(size Size, pieces []Piece) PieceStorage {
	extendedPieces := make([]Piece, size.PositionCount())
	for _, piece := range pieces {
		extendedPieces[size.PositionIndex(piece.Position())] = piece
	}

	baseBoard := NewBaseBoard(size)
	sliceBoard := SliceBoard{baseBoard, extendedPieces}
	return defaultBoardWrapper{sliceBoard}
}

// Piece ...
func (board SliceBoard) Piece(position Position) (piece Piece, ok bool) {
	piece = board.pieces[board.size.PositionIndex(position)]
	return piece, piece != nil
}

// ApplyMove ...
//
// It doesn't check that the move is correct.
func (board SliceBoard) ApplyMove(move Move) PieceStorage {
	pieceGroupCopy := make([]Piece, len(board.pieces))
	copy(pieceGroupCopy, board.pieces)

	startIndex, finishIndex :=
		board.size.PositionIndex(move.Start), board.size.PositionIndex(move.Finish)
	piece := pieceGroupCopy[startIndex]
	pieceGroupCopy[startIndex] = nil

	movedPiece := piece.ApplyPosition(move.Finish)
	pieceGroupCopy[finishIndex] = movedPiece

	sliceBoard := SliceBoard{board.BaseBoard, pieceGroupCopy}
	return defaultBoardWrapper{sliceBoard}
}
