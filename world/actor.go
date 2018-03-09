package world

import (
	"math"
	"strconv"

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

func (actor Actor) ToJSON() string {
	json := "{"
	json += "\"entity\":" + actor.Entity.ToJSON()
	json += ",\"direction\":" + strconv.FormatFloat(actor.Direction.Deg, 'f', -1, 64)
	json += "}"
	return json
}
