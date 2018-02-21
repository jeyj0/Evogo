package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombineSeeds(t *testing.T) {
	// given
	rand.Seed(1)
	seed1 := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	seed2 := []float64{8, 7, 6, 5, 4, 3, 2, 1}

	// when
	seedsCombined := CombineSeeds(seed1, seed2)

	// then
	assert.Equal(t, []float64{8, 7, 6, 5, 4, 6, 2, 8}, seedsCombined)
}
