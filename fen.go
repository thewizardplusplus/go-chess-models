package chessmodels

// PieceFactory ...
type PieceFactory func(
	kind Kind,
	color Color,
	position Position,
) (Piece, error)
