package world_test

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/jeyj0/Evogo/world"
	"github.com/stretchr/testify/assert"
)

func TestToJSON(t *testing.T) {
	// given
	world := givenWorldInstance()
	json_expected := givenJSON()

	// when
	json := world.ToJSON()
	fmt.Println(json)

	// then
	assert.Equal(t, json_expected, json)
}

func givenWorldInstance() world.World {
	actor := &world.Actor{}
	actors := []*world.Actor{actor}
	food := &world.Entity{}
	foods := []*world.Entity{food}
	return world.World{
		Width:  100,
		Height: 100,
		Actors: actors,
		Food:   foods,
	}
}

func givenJSON() string {
	return unWhitespaceString(`{
		"width": 100,
		"height": 100,
		"actors": [
			{
				"entity": {
					"color": {
						"red": 0,
						"green": 0,
						"blue": 0
					},
					"size": 0,
					"x": 0,
					"y": 0
				},
				"direction": 0
			}
		],
		"foods": [
			{
				"color": {
					"red": 0,
					"green": 0,
					"blue": 0
				},
				"size": 0,
				"x": 0,
				"y": 0
			}
		]
	}`)
}

func unWhitespaceString(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
