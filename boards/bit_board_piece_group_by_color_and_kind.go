package boards

import (
	"math/big"

	"github.com/thewizardplusplus/go-chess-models/common"
)

type positionStatus int

const (
	freePositionStatus positionStatus = iota
	occupiedPositionStatus
)

type bitBoardPieceGroupByColorAndKind big.Int

func (pieceGroup *bitBoardPieceGroupByColorAndKind) ToBigInt() *big.Int {
	return (*big.Int)(pieceGroup)
}

func (pieceGroup *bitBoardPieceGroupByColorAndKind) IsPositionOccupied(
	size common.Size,
	position common.Position,
) bool {
	positionIndex := size.PositionIndex(position)
	return pieceGroup.ToBigInt().Bit(positionIndex) == 1
}

func (pieceGroup *bitBoardPieceGroupByColorAndKind) SetValue(
	anotherPieceGroup *bitBoardPieceGroupByColorAndKind,
) {
	pieceGroup.ToBigInt().Set(anotherPieceGroup.ToBigInt())
}

func (pieceGroup *bitBoardPieceGroupByColorAndKind) SetPositionStatus(
	size common.Size,
	position common.Position,
	positionStatus positionStatus,
) {
	var positionValue uint
	switch positionStatus {
	case freePositionStatus:
		positionValue = 0
	case occupiedPositionStatus:
		positionValue = 1
	}

	positionIndex := size.PositionIndex(position)
	pieceGroup.ToBigInt().
		SetBit(pieceGroup.ToBigInt(), positionIndex, positionValue)
}
