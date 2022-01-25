package chessmodels

// PieceStorage ...
type PieceStorage interface {
	Size() Size
	Piece(position Position) (piece Piece, ok bool)
	Pieces() []Piece

	// It shouldn't check that the move is correct.
	ApplyMove(move Move) PieceStorage

	// It shouldn't check for a check before or after the move.
	CheckMove(move Move) error
}

type pieceGroup map[Position]Piece

// Board ...
type Board struct {
	size   Size
	pieces pieceGroup
}

// NewBoard ...
func NewBoard(size Size, pieces []Piece) PieceStorage {
	pieceGroup := make(pieceGroup)
	for _, piece := range pieces {
		pieceGroup[piece.Position()] = piece
	}

	return Board{size, pieceGroup}
}

// Size ...
func (board Board) Size() Size {
	return board.size
}

// Piece ...
func (board Board) Piece(position Position) (piece Piece, ok bool) {
	piece, ok = board.pieces[position]
	return piece, ok
}

// Pieces ...
//
// It doesn't guarantee an order of returned pieces.
func (board Board) Pieces() []Piece {
	var pieces []Piece
	for _, piece := range board.pieces {
		pieces = append(pieces, piece)
	}

	return pieces
}

// ApplyMove ...
//
// It doesn't check that the move is correct.
func (board Board) ApplyMove(move Move) PieceStorage {
	piece := board.pieces[move.Start]
	movedPiece := piece.ApplyPosition(move.Finish)

	pieceGroupCopy := pieceGroup{move.Finish: movedPiece}
	for position, piece := range board.pieces {
		if position != move.Start && position != move.Finish {
			pieceGroupCopy[position] = piece
		}
	}

	return Board{board.size, pieceGroupCopy}
}

// CheckMove ...
//
// It doesn't check for a check before or after the move.
func (board Board) CheckMove(move Move) error {
	return CheckMove(board, move)
}
