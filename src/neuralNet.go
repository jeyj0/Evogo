package main

type Neuron struct {
	value          float64
	bias           float64
	biasWeight     float64
	inConnections  []*Connection
	outConnections []*Connection
	isCalculated   bool
}

type Connection struct {
	inNeuron  *Neuron
	outNeuron *Neuron
	weight    float64
}

type Net struct {
	inputNeurons  []*Neuron
	outputNeurons []*Neuron
}

func (n *Neuron) calculateValue() {
	if !n.isCalculated {
		newValue := 0.0
		weightSum := 0.0

		if n.inConnections == nil {
			newValue = n.value
			weightSum = 1.0
		} else {
			for _, connection := range n.inConnections {
				newValue += connection.inNeuron.value * connection.weight
				weightSum += connection.weight
			}
		}
		newValue += n.bias * n.biasWeight
		newValue /= weightSum + n.biasWeight

		n.value = newValue

		n.isCalculated = true
	}
}

func (n *Neuron) calculateValueRecursive() {
	if n.inConnections != nil {
		for _, connection := range n.inConnections {
			connection.inNeuron.calculateValueRecursive()
		}
	}

	n.calculateValue()
}

func (net *Net) reset() {
	resetRecursive(net.outputNeurons)
}

func resetRecursive(neurons []*Neuron) {
	for _, neuron := range neurons {
		neuron.isCalculated = false
		nextNeurons := []*Neuron{}
		for _, connection := range neuron.inConnections {
			nextNeurons = append(nextNeurons, connection.inNeuron)
		}
		resetRecursive(nextNeurons)
	}
}

func (net *Net) CalculateOutputs() {
	for _, outputNeuron := range net.outputNeurons {
		outputNeuron.calculateValueRecursive()
	}
}
