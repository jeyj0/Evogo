package world

import "strconv"

type Entity struct {
	X     float64
	Y     float64
	Size  float64
	Color [3]uint8
}

func (e Entity) ToJSON() string {
	json := "{"
	json += "\"color\":{" +
		"\"red\":" + strconv.FormatUint(uint64(e.Color[0]), 8) +
		",\"green\":" + strconv.FormatUint(uint64(e.Color[1]), 8) +
		",\"blue\":" + strconv.FormatUint(uint64(e.Color[2]), 8) +
		"}"
	json += ",\"size\":" + strconv.FormatFloat(e.Size, 'f', -1, 64)
	json += ",\"x\":" + strconv.FormatFloat(e.X, 'f', -1, 64)
	json += ",\"y\":" + strconv.FormatFloat(e.Y, 'f', -1, 64)
	json += "}"
	return json
}
