package main

func GenerateFullyConnectedNeuralNet(neuronsPerLayer []int) Net {
	fullyConnectedNet := Net{}

	lastNewNeurons := generateEmptyNeurons(neuronsPerLayer[len(neuronsPerLayer)-1])
	fullyConnectedNet.outputNeurons = lastNewNeurons
	var newNeurons []*Neuron

	for i := len(neuronsPerLayer) - 1; i > 0; i-- {
		newNeurons = generateEmptyNeurons(neuronsPerLayer[i-1])
		if i > 0 {
			connections := generateWeightlessConnectionsFromNeurons(newNeurons)
			for _, neuron := range lastNewNeurons {
				neuron.attachConnections(connections)
			}
		}
		lastNewNeurons = newNeurons
	}

	fullyConnectedNet.inputNeurons = lastNewNeurons

	return fullyConnectedNet
}

func generateEmptyNeurons(count int) []*Neuron {
	neurons := []*Neuron{}
	for i := 0; i < count; i++ {
		neurons = append(neurons, &Neuron{})
	}
	return neurons
}

func (neuron *Neuron) attachConnections(connections []*Connection) {
	neuron.inConnections = append(neuron.inConnections, connections...)
}

func generateWeightlessConnectionsFromNeurons(neurons []*Neuron) []*Connection {
	connections := []*Connection{}
	for _, n := range neurons {
		connections = append(connections, &Connection{inNeuron: n})
	}
	return connections
}
