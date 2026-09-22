package darts

import (
	"math"
)

func Score(x, y float64) int {
	if x > 10 || y > 10 {
		return 0
	}

	xAbs := math.Abs(x)
	yAbs := math.Abs(y)
	rs := math.Sqrt(xAbs*xAbs + yAbs*yAbs)
	if rs > 5 && rs <= 10 {
		return 1
	} else if rs > 1 && rs <= 5 {
		return 5
	}else if rs >= 0 && rs <= 1 {
		return 10
	}
	return 0
}
