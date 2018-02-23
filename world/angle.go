package world

import "math"

type Angle struct {
	Deg float64
}

func (angle Angle) Rad() float64 {
	return angle.Deg * math.Pi / 180.0
}
