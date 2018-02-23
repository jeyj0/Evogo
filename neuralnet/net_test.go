package neuralnet_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	nn "github.com/jeyj0/Evogo/neuralnet"
)

func TestCalculateOutputs(t *testing.T) {
	// given
	input1 := nn.Neuron{Value: 1}
	input2 := nn.Neuron{Value: 2}

	out1Connections := []*nn.Connection{
		&nn.Connection{InNeuron: &input1, Weight: 3},
		&nn.Connection{InNeuron: &input2, Weight: 5}}

	out2Connections := []*nn.Connection{
		&nn.Connection{InNeuron: &input1, Weight: 4},
		&nn.Connection{InNeuron: &input2, Weight: 6}}

	output1 := nn.Neuron{Bias: 7, BiasWeight: 8, InConnections: out1Connections}
	output2 := nn.Neuron{Bias: 9, BiasWeight: 10, InConnections: out2Connections}

	net := nn.Net{InputNeurons: []*nn.Neuron{&input1, &input2}, OutputNeurons: []*nn.Neuron{&output1, &output2}}

	// when
	net.CalculateOutputs()

	// then
	assert.Equal(t, 4.3125, net.OutputNeurons[0].Value)
	assert.Equal(t, 5.3, net.OutputNeurons[1].Value)
}

func TestCalculateOutputsWithInputBias(t *testing.T) {
	// given
	input1 := nn.Neuron{Value: 1, Bias: 2, BiasWeight: 1} // 1.5
	input2 := nn.Neuron{Value: 2, Bias: 4, BiasWeight: 1} // 3

	out1Connections := []*nn.Connection{
		&nn.Connection{InNeuron: &input1, Weight: 1}, // 1.5
		&nn.Connection{InNeuron: &input2, Weight: 1}} // 3
	out2Connections := []*nn.Connection{
		&nn.Connection{InNeuron: &input1, Weight: 2}, // 3
		&nn.Connection{InNeuron: &input2, Weight: 1}} // 3

	output1 := nn.Neuron{Bias: 1, BiasWeight: 2, InConnections: out1Connections} // 1.625
	output2 := nn.Neuron{Bias: 1, BiasWeight: 1, InConnections: out2Connections} // 1.75

	net := nn.Net{InputNeurons: []*nn.Neuron{&input1, &input2}, OutputNeurons: []*nn.Neuron{&output1, &output2}}

	// when
	net.CalculateOutputs()

	// then
	assert.Equal(t, 1.625, net.OutputNeurons[0].Value)
	assert.Equal(t, 1.75, net.OutputNeurons[1].Value)
}
