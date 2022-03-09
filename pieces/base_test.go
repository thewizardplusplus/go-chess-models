package pieces

import (
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

func TestNewBase(test *testing.T) {
	piece := NewBase(common.Pawn, common.White, common.Position{
		File: 2,
		Rank: 3,
	})

	expectedPiece := Base{
		kind:  common.Pawn,
		color: common.White,
		position: common.Position{
			File: 2,
			Rank: 3,
		},
	}
	if !reflect.DeepEqual(piece, expectedPiece) {
		test.Fail()
	}
}

func TestBaseKind(test *testing.T) {
	piece := Base{kind: common.Pawn}
	kind := piece.Kind()

	if kind != common.Pawn {
		test.Fail()
	}
}

func TestBaseColor(test *testing.T) {
	piece := Base{color: common.White}
	color := piece.Color()

	if color != common.White {
		test.Fail()
	}
}

func TestBasePosition(test *testing.T) {
	piece := Base{
		position: common.Position{
			File: 2,
			Rank: 3,
		},
	}
	position := piece.Position()

	expectedPosition := common.Position{
		File: 2,
		Rank: 3,
	}
	if !reflect.DeepEqual(position, expectedPosition) {
		test.Fail()
	}
}

func TestBaseApplyPosition(test *testing.T) {
	piece := Base{
		kind:  common.Pawn,
		color: common.White,
		position: common.Position{
			File: 2,
			Rank: 3,
		},
	}
	nextPiece := piece.ApplyPosition(common.Position{
		File: 4,
		Rank: 2,
	})

	expectedPiece := Base{
		kind:  common.Pawn,
		color: common.White,
		position: common.Position{
			File: 2,
			Rank: 3,
		},
	}
	if !reflect.DeepEqual(piece, expectedPiece) {
		test.Fail()
	}

	expectedNextPiece := Base{
		kind:  common.Pawn,
		color: common.White,
		position: common.Position{
			File: 4,
			Rank: 2,
		},
	}
	if !reflect.DeepEqual(nextPiece, expectedNextPiece) {
		test.Fail()
	}
}
