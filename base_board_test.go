package chessmodels

import (
	"reflect"
	"testing"

	"github.com/thewizardplusplus/go-chess-models/common"
)

func TestNewBaseBoard(test *testing.T) {
	baseBoard := NewBaseBoard(common.Size{5, 5})

	expectedBaseBoard := BaseBoard{
		size: common.Size{5, 5},
	}
	if !reflect.DeepEqual(baseBoard, expectedBaseBoard) {
		test.Fail()
	}
}

func TestBaseBoardSize(test *testing.T) {
	baseBoard := NewBaseBoard(common.Size{5, 5})
	size := baseBoard.Size()

	if !reflect.DeepEqual(size, common.Size{5, 5}) {
		test.Fail()
	}
}
