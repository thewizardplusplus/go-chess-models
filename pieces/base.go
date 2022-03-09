package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/common"
)

// Base ...
type Base struct {
	kind     common.Kind
	color    common.Color
	position models.Position
}

// NewBase ...
func NewBase(
	kind common.Kind,
	color common.Color,
	position models.Position,
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
func (piece Base) Position() models.Position {
	return piece.position
}

// ApplyPosition ...
func (piece Base) ApplyPosition(position models.Position) Base {
	kind, color := piece.kind, piece.color
	return Base{kind, color, position}
}
