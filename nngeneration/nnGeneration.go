package nngeneration

import (
	nn "github.com/jeyj0/Evogo/neuralnet"
)

func GenerateFullyConnectedNeuralNet(neuronsPerLayer []int) nn.Net {
	fullyConnectedNet := nn.Net{}

	var layer []*nn.Neuron

	fullyConnectedNet.InputNeurons = generateEmptyNeurons(neuronsPerLayer[0])
	var lastLayer []*nn.Neuron = fullyConnectedNet.InputNeurons

	for i := 1; i <= len(neuronsPerLayer)-1; i++ {
		layer = generateEmptyNeurons(neuronsPerLayer[1])
		fullyConnect(lastLayer, layer)
		lastLayer = layer
	}
	fullyConnectedNet.OutputNeurons = layer

	return fullyConnectedNet
}

func FillWeightsAndBiasesFromSeed(net *nn.Net, seed []float64) *nn.Net {
	var layer []*nn.Neuron
	var connections []*nn.Connection
	layer = net.InputNeurons
	seed = fillWeightsAndBiasesForLayer(layer, seed)

	connections, layer = getConnectionsAndNextLayer(layer)
	// for seedLen := -1; len(seed) > 0 && len(seed) != seedLen; seedLen = len(seed) {
	for len(seed) > 0 {
		seed = fillWeightsForConnections(connections, seed)
		seed = fillWeightsAndBiasesForLayer(layer, seed)
		connections, layer = getConnectionsAndNextLayer(layer)
	}
	// println(len(seed))
	return net
}

func generateEmptyNeurons(count int) []*nn.Neuron {
	neurons := []*nn.Neuron{}
	for i := 0; i < count; i++ {
		neurons = append(neurons, &nn.Neuron{})
	}
	return neurons
}

func fullyConnect(layer1 []*nn.Neuron, layer2 []*nn.Neuron) {
	for _, l1neuron := range layer1 {
		for _, l2neuron := range layer2 {
			connection := &nn.Connection{InNeuron: l1neuron, OutNeuron: l2neuron}
			l1neuron.OutConnections = append(l1neuron.OutConnections, connection)
			l2neuron.InConnections = append(l2neuron.InConnections, connection)
		}
	}
}

func generateWeightlessConnectionsFromNeurons(neurons []*nn.Neuron) []*nn.Connection {
	connections := []*nn.Connection{}
	for _, n := range neurons {
		connections = append(connections, &nn.Connection{InNeuron: n})
	}
	return connections
}

func fillWeightsAndBiasesForLayer(layer []*nn.Neuron, seed []float64) []float64 {
	for _, neuron := range layer {
		neuron.Bias = seed[0]
		neuron.BiasWeight = seed[1]
		seed = seed[2:]
	}
	return seed
}

func fillWeightsForConnections(connections []*nn.Connection, seed []float64) []float64 {
	for _, connection := range connections {
		connection.Weight = seed[0]
		seed = seed[1:]
	}
	return seed
}

func getConnectionsAndNextLayer(layer []*nn.Neuron) ([]*nn.Connection, []*nn.Neuron) {
	contains := func(layer []*nn.Neuron, neuron *nn.Neuron) bool {
		for _, n := range layer {
			if neuron == n {
				return true
			}
		}
		return false
	}

	nextLayer := []*nn.Neuron{}
	connections := []*nn.Connection{}
	for _, neuron := range layer {
		connections = append(connections, neuron.OutConnections...)
		for _, connection := range neuron.OutConnections {
			if !contains(nextLayer, connection.OutNeuron) {
				nextLayer = append(nextLayer, connection.OutNeuron)
			}
		}
	}

	return connections, nextLayer
}
