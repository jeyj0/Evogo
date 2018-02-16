package main

func main() {
}

type Neuron struct {
	value         float64
	bias          float64
	biasWeight    float64
	inConnections []*Connection
}

type Connection struct {
	inNeuron *Neuron
	weight   float64
}

type Net struct {
	inputNeurons  []*Neuron
	outputNeurons []*Neuron
}

func (n *Neuron) calculateValue() {
	newValue := 0.0
	weightSum := 0.0

	for _, connection := range n.inConnections {
		newValue += connection.inNeuron.value * connection.weight
		weightSum += connection.weight
	}
	newValue += n.bias * n.biasWeight
	newValue /= weightSum + n.biasWeight

	n.value = newValue
}

func (net *Net) calculateOutputs(inputs []float64) []float64 {
	for _, neuron := range net.outputNeurons {
		neuron.calculateValue()
	}
	return nil
}
