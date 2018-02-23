package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveUp(t *testing.T) {
	// given
	actor := Actor{Entity: Entity{x: 0, y: 0}, direction: Angle{0}}

	// when
	actor.Move(1)

	// then
	assert.Equal(t, 0.0, actor.x)
	assert.Equal(t, 1.0, actor.y)
}

func TestMoveUpRight(t *testing.T) {
	// given
	actor := Actor{Entity: Entity{x: 0, y: 0}, direction: Angle{45}}

	// when
	actor.Move(1)

	// then
	assert.Equal(t, 0.7071067811865475, actor.x)
	assert.Equal(t, 0.7071067811865476, actor.y)
}

func TestMoveDown(t *testing.T) {
	// given
	actor := Actor{Entity: Entity{x: 0, y: 0}, direction: Angle{180}}

	// when
	actor.Move(1)

	// then
	const ZERO_BECAUSE_OF_FLOAT_CALC float64 = 1.2246467991473515e-16
	assert.Equal(t, ZERO_BECAUSE_OF_FLOAT_CALC, actor.x)
	assert.Equal(t, -1.0, actor.y)
}

func TestTurn90(t *testing.T) {
	// given
	actor := Actor{direction: Angle{0}}

	// when
	actor.Turn(90)

	// then
	assert.Equal(t, 90.0, actor.direction.deg)
}

func TestTurn360(t *testing.T) {
	// given
	actor := Actor{direction: Angle{0}}

	// when
	actor.Turn(360)

	// then
	assert.Equal(t, 0.0, actor.direction.deg)
}

func TestTurnNeg90(t *testing.T) {
	// given
	actor := Actor{direction: Angle{0}}

	// when
	actor.Turn(-90)

	// then
	assert.Equal(t, 270.0, actor.direction.deg)
}

func TestTurn361(t *testing.T) {
	// given
	actor := Actor{direction: Angle{0}}

	// when
	actor.Turn(361)

	// then
	assert.Equal(t, 1.0, actor.direction.deg)
}

func TestTurnNeg361(t *testing.T) {
	// given
	actor := Actor{direction: Angle{0}}

	// when
	actor.Turn(-361)

	// then
	assert.Equal(t, 359.0, actor.direction.deg)
}
