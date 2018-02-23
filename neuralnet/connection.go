package neuralnet

type Connection struct {
	InNeuron  *Neuron
	OutNeuron *Neuron
	Weight    float64
}
