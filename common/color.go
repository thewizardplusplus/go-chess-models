package common

// Color ...
type Color int

// ...
const (
	Black Color = iota
	White

	ColorCount
)

// Negative ...
func (color Color) Negative() Color {
	if color == Black {
		return White
	}

	return Black
}
