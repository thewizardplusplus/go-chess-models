package pieces

import (
	"math"
)

func steps(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}
