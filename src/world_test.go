package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	// given
	actor := Actor{Entity: Entity{x: 0, y: 0}, direction: 0}

	// when
	actor.Move(1)

	// then
	assert.Equal(t, actor.x, 0.0)
	assert.Equal(t, actor.y, 1.0)
	assert.Equal(t, 0.0, actor.direction)
}

func TestTurn90(t *testing.T) {
	// given
	actor := Actor{direction: 0}

	// when
	actor.Turn(90)

	// then
	assert.Equal(t, 90.0, actor.direction)
}

func TestTurn360(t *testing.T) {
	// given
	actor := Actor{direction: 0}

	// when
	actor.Turn(360)

	// then
	assert.Equal(t, 0.0, actor.direction)
}

func TestTurnNeg90(t *testing.T) {
	// given
	actor := Actor{direction: 0}

	// when
	actor.Turn(-90)

	// then
	assert.Equal(t, 270.0, actor.direction)
}
