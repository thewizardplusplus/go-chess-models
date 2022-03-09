package pieces

import (
	"github.com/thewizardplusplus/go-chess-models/common"
)

// Base ...
type Base struct {
	kind     common.Kind
	color    common.Color
	position common.Position
}

// NewBase ...
func NewBase(
	kind common.Kind,
	color common.Color,
	position common.Position,
) Base {
	return Base{kind, color, position}
}

// Kind ...
func (piece Base) Kind() common.Kind {
	return piece.kind
}

// Color ...
func (piece Base) Color() common.Color {
	return piece.color
}

// Position ...
func (piece Base) Position() common.Position {
	return piece.position
}

// ApplyPosition ...
func (piece Base) ApplyPosition(position common.Position) Base {
	kind, color := piece.kind, piece.color
	return Base{kind, color, position}
}
