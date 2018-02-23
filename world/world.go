package world

import (
	"math"

	nn "github.com/jeyj0/Evogo/neuralnet"
)

type World struct {
	Actors []*Actor
	Food   []*Entity
	Width  float64
	Height float64
}

type Entity struct {
	X     float64
	Y     float64
	Size  float64
	Color [3]uint8
}

type Angle struct {
	Deg float64
}

func (angle Angle) Rad() float64 {
	return angle.Deg * math.Pi / 180.0
}

type Actor struct {
	Entity
	Net       *nn.Net
	Direction Angle
}

func (actor *Actor) Move(amount float64) {
	actor.X += math.Sin(actor.Direction.Rad()) * amount
	actor.Y += math.Cos(actor.Direction.Rad()) * amount
}

func (actor *Actor) Turn(amount float64) {
	actor.Direction.Deg = math.Mod(math.Mod((actor.Direction.Deg+amount), 360)+360, 360)
}
