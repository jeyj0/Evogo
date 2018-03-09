package world

import "strconv"

type World struct {
	Actors []*Actor
	Food   []*Entity
	Width  float64
	Height float64
}

func (world World) ToJSON() string {
	json := "{"

	json += "\"width\":" + strconv.FormatFloat(world.Width, 'f', -1, 64)
	json += ",\"height\":" + strconv.FormatFloat(world.Height, 'f', -1, 64)

	json += ",\"actors\":["
	for _, actor := range world.Actors {
		json += actor.ToJSON() + ","
	}
	if len(world.Actors) > 0 {
		json = json[:len(json)-1]
	}
	json += "]"

	json += ",\"foods\":["
	for _, food := range world.Food {
		json += food.ToJSON() + ","
	}
	if len(world.Food) > 0 {
		json = json[:len(json)-1]
	}
	json += "]"

	json += "}"
	return json
}
