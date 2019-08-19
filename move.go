package chessmodels

// Move ...
type Move struct {
	Start  Position
	Finish Position
}

// IsZero ...
func (move Move) IsZero() bool {
	return move == Move{}
}
