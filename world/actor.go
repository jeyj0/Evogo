package world

import (
	"math"

	nn "github.com/jeyj0/Evogo/neuralnet"
)

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
