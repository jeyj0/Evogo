package main

import "math/rand"

func combineSeeds(seed1, seed2 []float64) []float64 {
	newSeed := []float64{}
	for i, seed1Val := range seed1 {
		if rand.Intn(2) == 0 {
			newSeed = append(newSeed, seed1Val)
		} else {
			newSeed = append(newSeed, seed2[i])
		}
	}
	return newSeed
}
