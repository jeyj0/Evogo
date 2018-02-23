package neuralnet

type Neuron struct {
	Value          float64
	Bias           float64
	BiasWeight     float64
	InConnections  []*Connection
	OutConnections []*Connection
	isCalculated   bool
}

type Connection struct {
	InNeuron  *Neuron
	OutNeuron *Neuron
	Weight    float64
}

type Net struct {
	InputNeurons  []*Neuron
	OutputNeurons []*Neuron
}

func (n *Neuron) calculateValue() {
	if !n.isCalculated {
		newValue := 0.0
		weightSum := 0.0

		if n.InConnections == nil {
			newValue = n.Value
			weightSum = 1.0
		} else {
			for _, connection := range n.InConnections {
				newValue += connection.InNeuron.Value * connection.Weight
				weightSum += connection.Weight
			}
		}
		newValue += n.Bias * n.BiasWeight
		newValue /= weightSum + n.BiasWeight

		n.Value = newValue

		n.isCalculated = true
	}
}

func (n *Neuron) calculateValueRecursive() {
	if n.InConnections != nil {
		for _, connection := range n.InConnections {
			connection.InNeuron.calculateValueRecursive()
		}
	}

	n.calculateValue()
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

func (net *Net) CalculateOutputs() {
	net.reset()
	for _, outputNeuron := range net.OutputNeurons {
		outputNeuron.calculateValueRecursive()
	}
}
