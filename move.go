package chessmodels

// Move ...
type Move struct {
	Start  Position
	Finish Position
}

// IsZero ...
//
// It checks that all fields of the move are zero.
func (move Move) IsZero() bool {
	return move == Move{}
}

// IsEmpty ...
//
// It checks that the start of the move equals its finish.
func (move Move) IsEmpty() bool {
	return move.Start == move.Finish
}
