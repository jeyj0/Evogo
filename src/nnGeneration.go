package main

func GenerateFullyConnectedNeuralNet(neuronsPerLayer []int) Net {
	fullyConnectedNet := Net{}

	var layer []*Neuron

	fullyConnectedNet.inputNeurons = generateEmptyNeurons(neuronsPerLayer[0])
	var lastLayer []*Neuron = fullyConnectedNet.inputNeurons

	for i := 1; i <= len(neuronsPerLayer)-1; i++ {
		layer = generateEmptyNeurons(neuronsPerLayer[1])
		fullyConnect(lastLayer, layer)
		lastLayer = layer
	}
	fullyConnectedNet.outputNeurons = layer

	return fullyConnectedNet
}

func (net *Net) FillWeightsAndBiasesFromSeed(seed []float64) {
	var layer []*Neuron
	var connections []*Connection
	layer = net.inputNeurons
	seed = fillWeightsAndBiasesForLayer(layer, seed)

	connections, layer = getConnectionsAndNextLayer(layer)
	for len(seed) > 0 {
		seed = fillWeightsForConnections(connections, seed)
		seed = fillWeightsAndBiasesForLayer(layer, seed)
		connections, layer = getConnectionsAndNextLayer(layer)
	}
}

func generateEmptyNeurons(count int) []*Neuron {
	neurons := []*Neuron{}
	for i := 0; i < count; i++ {
		neurons = append(neurons, &Neuron{})
	}
	return neurons
}

func fullyConnect(layer1 []*Neuron, layer2 []*Neuron) {
	for _, l1neuron := range layer1 {
		for _, l2neuron := range layer2 {
			connection := &Connection{inNeuron: l1neuron, outNeuron: l2neuron}
			l1neuron.outConnections = append(l1neuron.outConnections, connection)
			l2neuron.inConnections = append(l2neuron.inConnections, connection)
		}
	}
}

func generateWeightlessConnectionsFromNeurons(neurons []*Neuron) []*Connection {
	connections := []*Connection{}
	for _, n := range neurons {
		connections = append(connections, &Connection{inNeuron: n})
	}
	return connections
}

func fillWeightsAndBiasesForLayer(layer []*Neuron, seed []float64) []float64 {
	for _, neuron := range layer {
		neuron.bias = seed[0]
		neuron.biasWeight = seed[1]
		seed = seed[2:]
	}
	return seed
}

func fillWeightsForConnections(connections []*Connection, seed []float64) []float64 {
	for _, connection := range connections {
		connection.weight = seed[0]
		seed = seed[1:]
	}
	return seed
}

func getConnectionsAndNextLayer(layer []*Neuron) ([]*Connection, []*Neuron) {
	contains := func(layer []*Neuron, neuron *Neuron) bool {
		for _, n := range layer {
			if neuron == n {
				return true
			}
		}
		return false
	}

	nextLayer := []*Neuron{}
	connections := []*Connection{}
	for _, neuron := range layer {
		connections = append(connections, neuron.outConnections...)
		for _, connection := range neuron.outConnections {
			if !contains(nextLayer, connection.outNeuron) {
				nextLayer = append(nextLayer, connection.outNeuron)
			}
		}
	}

	return connections, nextLayer
}
