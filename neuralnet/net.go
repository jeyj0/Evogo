package neuralnet

type Net struct {
	InputNeurons  []*Neuron
	OutputNeurons []*Neuron
}

func (net *Net) CalculateOutputs() {
	net.reset()
	for _, outputNeuron := range net.OutputNeurons {
		outputNeuron.calculateValueRecursive()
	}
}

func (net *Net) reset() {
	resetRecursive(net.OutputNeurons)
}

func resetRecursive(neurons []*Neuron) {
	for _, neuron := range neurons {
		neuron.isCalculated = false
		nextNeurons := []*Neuron{}
		for _, connection := range neuron.InConnections {
			nextNeurons = append(nextNeurons, connection.InNeuron)
		}
		resetRecursive(nextNeurons)
	}
}
