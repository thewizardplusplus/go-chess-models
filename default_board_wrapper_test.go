package chessmodels

type MockBasePieceStorage struct {
	size Size

	piece func(position Position) (piece Piece, ok bool)
}

func (storage MockBasePieceStorage) Size() Size {
	return storage.size
}

func (storage MockBasePieceStorage) Piece(
	position Position,
) (piece Piece, ok bool) {
	if storage.piece == nil {
		panic("not implemented")
	}

	return storage.piece(position)
}

func (storage MockBasePieceStorage) ApplyMove(move Move) PieceStorage {
	panic("not implemented")
}
