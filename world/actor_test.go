package world_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	world "github.com/jeyj0/Evogo/world"
)

func TestMoveUp(t *testing.T) {
	// given
	actor := world.Actor{Entity: world.Entity{X: 0, Y: 0}, Direction: world.Angle{0}}

	// when
	actor.Move(1)

	// then
	assert.Equal(t, 0.0, actor.X)
	assert.Equal(t, 1.0, actor.Y)
}

func TestMoveUpRight(t *testing.T) {
	// given
	actor := world.Actor{Entity: world.Entity{X: 0, Y: 0}, Direction: world.Angle{45}}

	// when
	actor.Move(1)

	// then
	assert.Equal(t, 0.7071067811865475, actor.X)
	assert.Equal(t, 0.7071067811865476, actor.Y)
}

func TestMoveDown(t *testing.T) {
	// given
	actor := world.Actor{Entity: world.Entity{X: 0, Y: 0}, Direction: world.Angle{180}}

	// when
	actor.Move(1)

	// then
	const ZERO_BECAUSE_OF_FLOAT_CALC float64 = 1.2246467991473515e-16
	assert.Equal(t, ZERO_BECAUSE_OF_FLOAT_CALC, actor.X)
	assert.Equal(t, -1.0, actor.Y)
}

func TestTurn90(t *testing.T) {
	// given
	actor := world.Actor{Direction: world.Angle{0}}

	// when
	actor.Turn(90)

	// then
	assert.Equal(t, 90.0, actor.Direction.Deg)
}

func TestTurn360(t *testing.T) {
	// given
	actor := world.Actor{Direction: world.Angle{0}}

	// when
	actor.Turn(360)

	// then
	assert.Equal(t, 0.0, actor.Direction.Deg)
}

func TestTurnNeg90(t *testing.T) {
	// given
	actor := world.Actor{Direction: world.Angle{0}}

	// when
	actor.Turn(-90)

	// then
	assert.Equal(t, 270.0, actor.Direction.Deg)
}

func TestTurn361(t *testing.T) {
	// given
	actor := world.Actor{Direction: world.Angle{0}}

	// when
	actor.Turn(361)

	// then
	assert.Equal(t, 1.0, actor.Direction.Deg)
}

func TestTurnNeg361(t *testing.T) {
	// given
	actor := world.Actor{Direction: world.Angle{0}}

	// when
	actor.Turn(-361)

	// then
	assert.Equal(t, 359.0, actor.Direction.Deg)
}
