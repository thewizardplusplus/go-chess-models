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
	return SliceBoard{baseBoard, extendedPieces}
}

// Piece ...
func (board SliceBoard) Piece(position Position) (piece Piece, ok bool) {
	piece = board.pieces[board.size.PositionIndex(position)]
	return piece, piece != nil
}

// Pieces ...
func (board SliceBoard) Pieces() []Piece {
	var pieces []Piece
	for _, piece := range board.pieces {
		if piece != nil {
			pieces = append(pieces, piece)
		}
	}

	return pieces
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

	return SliceBoard{board.BaseBoard, pieceGroupCopy}
}

// CheckMove ...
//
// It doesn't check for a check before or after the move.
func (board SliceBoard) CheckMove(move Move) error {
	return CheckMove(board, move)
}
