package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const minActors int = 100
const runs int = 200

func main() {
	world := createWorld()
	mainLoop(world)
}

func createWorld() *World {
	world := World{width: 100, height: 100}

	actors := []*Actor{}
	for i := 0; i < minActors; i++ {
		x := 0.0
		y := 0.0
		net := GenerateFullyConnectedNeuralNet([]int{30, 50, 20})
		net.FillWeightsAndBiasesFromSeed(generateSeed(1660))

		actor := Actor{Entity: Entity{x: x, y: y, size: 1}, direction: Angle{0}, net: &net}
		actors = append(actors, &actor)
	}
	world.actors = actors

	return &world
}

func generateSeed(length int) []float64 {
	seed := []float64{}
	for i := 0; i < length; i++ {
		seed = append(seed, rand.Float64())
	}
	return seed
}

func mainLoop(world *World) {
	start := time.Now()
	for i := 0; i < runs; i++ {
		letActorsActSync(world.actors)
	}
	delta := time.Now().Sub(start).Seconds()
	fmt.Println("Duration (sync)  :", delta, "seconds | FPS:", float64(runs)/delta)

	start = time.Now()
	for i := 0; i < runs; i++ {
		letActorsActChrono(world.actors)
	}
	delta = time.Now().Sub(start).Seconds()
	fmt.Println("Duration (chrono):", delta, "seconds | FPS:", float64(runs)/delta)
}

func letActorsActSync(actors []*Actor) {

	waitGroup := new(sync.WaitGroup)

	for _, actor := range actors {
		waitGroup.Add(1)
		go func(net *Net) {
			defer waitGroup.Done()

			for _, n := range net.inputNeurons {
				n.value = rand.Float64()
			}
			net.CalculateOutputs()
		}(actor.net)
	}

	waitGroup.Wait()
}

func letActorsActChrono(actors []*Actor) {
	for _, actor := range actors {
		for _, n := range actor.net.inputNeurons {
			n.value = rand.Float64()
		}
		actor.net.CalculateOutputs()
	}
}
