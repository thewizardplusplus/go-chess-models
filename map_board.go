package chessmodels

// MapBoard ...
type MapBoard struct {
	size   Size
	pieces pieceGroup
}

// NewMapBoard ...
func NewMapBoard(size Size, pieces []Piece) PieceStorage {
	pieceGroup := newPieceGroup(pieces)
	return MapBoard{size, pieceGroup}
}

// Size ...
func (board MapBoard) Size() Size {
	return board.size
}

// Piece ...
func (board MapBoard) Piece(position Position) (piece Piece, ok bool) {
	piece, ok = board.pieces[position]
	return piece, ok
}

// Pieces ...
//
// It doesn't guarantee an order of returned pieces.
func (board MapBoard) Pieces() []Piece {
	var pieces []Piece
	for _, piece := range board.pieces {
		pieces = append(pieces, piece)
	}

	return pieces
}

// ApplyMove ...
//
// It doesn't check that the move is correct.
func (board MapBoard) ApplyMove(move Move) PieceStorage {
	piece := board.pieces[move.Start]
	movedPiece := piece.ApplyPosition(move.Finish)

	pieceGroupCopy := pieceGroup{move.Finish: movedPiece}
	for position, piece := range board.pieces {
		if position != move.Start && position != move.Finish {
			pieceGroupCopy[position] = piece
		}
	}

	return MapBoard{board.size, pieceGroupCopy}
}

// CheckMove ...
//
// It doesn't check for a check before or after the move.
func (board MapBoard) CheckMove(move Move) error {
	return CheckMove(board, move)
}
