package main

import (
	"math"
)

type World struct {
	entities []*Entity
	width    float64
	height   float64
}

type Entity struct {
	x     float64
	y     float64
	size  float64
	color [3]uint8
}

type Angle struct {
	deg float64
}

func (angle Angle) Rad() float64 {
	return angle.deg * math.Pi / 180.0
}

type Actor struct {
	Entity
	direction Angle
}

type Food struct {
	Entity
}

func (actor *Actor) Move(amount float64) {
	actor.x += math.Sin(actor.direction.Rad()) * amount
	actor.y += math.Cos(actor.direction.Rad()) * amount
}

func (actor *Actor) Turn(amount float64) {
	actor.direction.deg = math.Mod(math.Mod((actor.direction.deg+amount), 360)+360, 360)
}
