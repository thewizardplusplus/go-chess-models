package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

type Base struct {
	kind     models.Kind
	color    models.Color
	position models.Position
}

func (piece Base) Kind() models.Kind {
	return piece.kind
}

func (piece Base) Color() models.Color {
	return piece.color
}

func (
	piece Base,
) Position() models.Position {
	return piece.position
}

func (piece Base) ApplyPosition(
	position models.Position,
) Base {
	kind, color := piece.kind, piece.color
	return Base{kind, color, position}
}
