package boards

import (
	"errors"

	"github.com/thewizardplusplus/go-chess-models/common"
)

var (
	errBitBoardPieceGroupIterationStop = errors.New("iteration stop")
)

type bitBoardPieceGroupByColorAndKindHandler func(
	color common.Color,
	kind common.Kind,
	piecesByColorAndKind *bitBoardPieceGroupByColorAndKind,
) error

type bitBoardPieceGroup [common.ColorCount][common.KindCount]bitBoardPieceGroupByColorAndKind // nolint: lll

func (pieceGroup *bitBoardPieceGroup) IteratePiecesByColorAndKind(
	handler bitBoardPieceGroupByColorAndKindHandler,
) error {
	for colorAsInt := 0; colorAsInt < int(common.ColorCount); colorAsInt++ {
		for kindAsInt := 0; kindAsInt < int(common.KindCount); kindAsInt++ {
			color, kind := common.Color(colorAsInt), common.Kind(kindAsInt)
			piecesByColorAndKind := &pieceGroup[color][kind]
			if err := handler(color, kind, piecesByColorAndKind); err != nil {
				return err
			}
		}
	}

	return nil
}

func (pieceGroup *bitBoardPieceGroup) PieceByPosition(
	size common.Size,
	position common.Position,
	pieceFactory common.PieceFactory,
) (
	piece common.Piece,
	piecesByColorAndKind *bitBoardPieceGroupByColorAndKind,
	ok bool,
) {
	pieceGroup.IteratePiecesByColorAndKind(func( // nolint: errcheck
		color common.Color,
		kind common.Kind,
		specificPiecesByColorAndKind *bitBoardPieceGroupByColorAndKind,
	) error {
		if !specificPiecesByColorAndKind.IsPositionOccupied(size, position) {
			return nil
		}

		piece = pieceFactory(kind, color, position)
		piecesByColorAndKind = specificPiecesByColorAndKind
		ok = true

		return errBitBoardPieceGroupIterationStop
	})

	return piece, piecesByColorAndKind, ok
}

func (pieceGroup *bitBoardPieceGroup) SetValue(
	anotherPieceGroup *bitBoardPieceGroup,
) {
	anotherPieceGroup.IteratePiecesByColorAndKind(func( // nolint: errcheck
		color common.Color,
		kind common.Kind,
		piecesByColorAndKind *bitBoardPieceGroupByColorAndKind,
	) error {
		pieceGroup[color][kind].SetValue(piecesByColorAndKind)
		return nil
	})
}

func (pieceGroup *bitBoardPieceGroup) AddPiece(
	size common.Size,
	piece common.Piece,
) {
	pieceGroup[piece.Color()][piece.Kind()].
		SetPositionStatus(size, piece.Position(), occupiedPositionStatus)
}

func (pieceGroup *bitBoardPieceGroup) ClearPosition(
	size common.Size,
	position common.Position,
	pieceFactory common.PieceFactory,
) (piece common.Piece, ok bool) {
	piece, piecesByColorAndKind, ok :=
		pieceGroup.PieceByPosition(size, position, pieceFactory)
	if !ok {
		return nil, false
	}

	piecesByColorAndKind.SetPositionStatus(size, position, freePositionStatus)

	return piece, true
}
