package world

type World struct {
	Actors []*Actor
	Food   []*Entity
	Width  float64
	Height float64
}
