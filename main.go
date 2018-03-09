package main

import (
	"math/rand"
	"sync"

	nn "github.com/jeyj0/Evogo/neuralnet"
	nng "github.com/jeyj0/Evogo/nngeneration"
	"github.com/jeyj0/Evogo/websocket"
	w "github.com/jeyj0/Evogo/world"
)

const minActors int = 100
const runs int = 200

func main() {
	world := createWorld()
	ws := websocket.Init()
	mainLoop(world, ws)
}

func createWorld() *w.World {
	world := w.World{Width: 100, Height: 100}

	actors := []*w.Actor{}
	for i := 0; i < minActors; i++ {
		var x float64
		var y float64
		net := nng.GenerateFullyConnectedNeuralNet([]int{30, 50, 20})
		nng.FillWeightsAndBiasesFromSeed(&net, generateSeed(1660))

		actor := w.Actor{Entity: w.Entity{X: x, Y: y, Size: 1}, Direction: w.Angle{}, Net: &net}
		actors = append(actors, &actor)
	}
	world.Actors = actors

	return &world
}

func generateSeed(length int) []float64 {
	seed := []float64{}
	for i := 0; i < length; i++ {
		seed = append(seed, rand.Float64())
	}
	return seed
}

func mainLoop(world *w.World, socket *websocket.Socket) {
	dataJSON := "{}"
	for i := 0; i < runs; i++ {
		letActorsActSync(world.Actors)

		dataJSON = world.ToJSON()

		socket.Message = dataJSON
	}
}

func letActorsActSync(actors []*w.Actor) {

	waitGroup := new(sync.WaitGroup)

	for _, actor := range actors {
		waitGroup.Add(1)
		go func(net *nn.Net) {
			defer waitGroup.Done()

			for _, n := range net.InputNeurons {
				n.Value = rand.Float64()
			}
			net.CalculateOutputs()
		}(actor.Net)
	}

	waitGroup.Wait()
}

func letActorsActChrono(actors []*w.Actor) {
	for _, actor := range actors {
		for _, n := range actor.Net.InputNeurons {
			n.Value = rand.Float64()
		}
		actor.Net.CalculateOutputs()
	}
}
