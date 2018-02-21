package main

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

type Actor struct {
	Entity
	direction float64
}

type Food struct {
	Entity
}

func (actor *Actor) Move(amount float64) {
}

func (actor *Actor) Turn(amount float64) {
}
