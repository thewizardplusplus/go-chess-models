package pieces

import (
	models "github.com/thewizardplusplus/go-chess-models"
)

// Factory ...
type Factory func(
	color models.Color,
	position models.Position,
) models.Piece
