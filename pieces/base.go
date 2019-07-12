package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

type base struct {
	kind     models.Kind
	color    models.Color
	position models.Position
}

// Kind ...
func (piece base) Kind() models.Kind {
	return piece.kind
}

// Color ...
func (piece base) Color() models.Color {
	return piece.color
}

// Position ...
func (
	piece base,
) Position() models.Position {
	return piece.position
}

// ApplyPosition ...
func (piece base) ApplyPosition(
	position models.Position,
) base {
	kind, color := piece.kind, piece.color
	return base{kind, color, position}
}
