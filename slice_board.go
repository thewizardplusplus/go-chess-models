package chessmodels

// SliceBoard ...
type SliceBoard struct {
	size   Size
	pieces []Piece
}

// NewSliceBoard ...
func NewSliceBoard(size Size, pieces []Piece) PieceStorage {
	var extendedPieces []Piece
	pieceGroup := newPieceGroup(pieces)
	for _, position := range size.Positions() {
		piece := pieceGroup[position] // if the position is empty, the piece is nil
		extendedPieces = append(extendedPieces, piece)
	}

	return SliceBoard{size, extendedPieces}
}

// Size ...
func (board SliceBoard) Size() Size {
	return board.size
}

// Piece ...
func (board SliceBoard) Piece(position Position) (piece Piece, ok bool) {
	piece = board.pieces[board.index(position)]
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

	startIndex, finishIndex := board.index(move.Start), board.index(move.Finish)
	piece := pieceGroupCopy[startIndex]
	pieceGroupCopy[startIndex] = nil

	movedPiece := piece.ApplyPosition(move.Finish)
	pieceGroupCopy[finishIndex] = movedPiece

	return SliceBoard{board.size, pieceGroupCopy}
}

// CheckMove ...
//
// It doesn't check for a check before or after the move.
func (board SliceBoard) CheckMove(move Move) error {
	return CheckMove(board, move)
}

func (board SliceBoard) index(position Position) int {
	return board.size.Width*position.Rank + position.File
}
