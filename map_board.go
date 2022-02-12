package chessmodels

type pieceGroup map[Position]Piece

// MapBoard ...
type MapBoard struct {
	BaseBoard

	pieces pieceGroup
}

// NewMapBoard ...
func NewMapBoard(size Size, pieces []Piece) PieceStorage {
	pieceGroup := make(pieceGroup) // nolint: vetshadow
	for _, piece := range pieces {
		pieceGroup[piece.Position()] = piece
	}

	baseBoard := NewBaseBoard(size)
	mapBoard := MapBoard{baseBoard, pieceGroup}
	return DefaultBoardWrapper{mapBoard}
}

// Piece ...
func (board MapBoard) Piece(position Position) (piece Piece, ok bool) {
	piece, ok = board.pieces[position]
	return piece, ok
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

	mapBoard := MapBoard{board.BaseBoard, pieceGroupCopy}
	return DefaultBoardWrapper{mapBoard}
}
