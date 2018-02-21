package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombineSeeds(t *testing.T) {
	// given
	seed1 := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	seed2 := []float64{8, 7, 6, 5, 4, 3, 2, 1}

	// when
	seedsCombined := combineSeeds(seed1, seed2)

	// then
	assert.Equal(t, []float64{1, 2, 3, 5, 4, 3, 2, 1}, seedsCombined)
}
