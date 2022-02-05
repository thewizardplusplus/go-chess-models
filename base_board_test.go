package chessmodels

import (
	"reflect"
	"testing"
)

func TestNewBaseBoard(test *testing.T) {
	baseBoard := NewBaseBoard(Size{5, 5})

	expectedBaseBoard := BaseBoard{
		size: Size{5, 5},
	}
	if !reflect.DeepEqual(baseBoard, expectedBaseBoard) {
		test.Fail()
	}
}

func TestBaseBoardSize(test *testing.T) {
	baseBoard := NewBaseBoard(Size{5, 5})
	size := baseBoard.Size()

	if !reflect.DeepEqual(size, Size{5, 5}) {
		test.Fail()
	}
}
