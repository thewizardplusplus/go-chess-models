package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

// Base ...
type Base struct {
	kind     models.Kind
	color    models.Color
	position models.Position
}

// NewBase ...
func NewBase(
	kind models.Kind,
	color models.Color,
	position models.Position,
) Base {
	return Base{kind, color, position}
}

// Kind ...
func (piece Base) Kind() models.Kind {
	return piece.kind
}

// Color ...
func (piece Base) Color() models.Color {
	return piece.color
}

// Position ...
func (
	piece Base,
) Position() models.Position {
	return piece.position
}

// ApplyPosition ...
func (piece Base) ApplyPosition(
	position models.Position,
) Base {
	kind, color := piece.kind, piece.color
	return Base{kind, color, position}
}
