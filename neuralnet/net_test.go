package neuralnet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateOutputs(t *testing.T) {
	// given
	input1 := Neuron{Value: 1}
	input2 := Neuron{Value: 2}

	out1Connections := []*Connection{
		&Connection{InNeuron: &input1, Weight: 3},
		&Connection{InNeuron: &input2, Weight: 5}}

	out2Connections := []*Connection{
		&Connection{InNeuron: &input1, Weight: 4},
		&Connection{InNeuron: &input2, Weight: 6}}

	output1 := Neuron{Bias: 7, BiasWeight: 8, InConnections: out1Connections}
	output2 := Neuron{Bias: 9, BiasWeight: 10, InConnections: out2Connections}

	net := Net{InputNeurons: []*Neuron{&input1, &input2}, OutputNeurons: []*Neuron{&output1, &output2}}

	// when
	net.CalculateOutputs()

	// then
	assert.Equal(t, 4.3125, net.OutputNeurons[0].Value)
	assert.Equal(t, 5.3, net.OutputNeurons[1].Value)
}

func TestCalculateOutputsWithInputBias(t *testing.T) {
	// given
	input1 := Neuron{Value: 1, Bias: 2, BiasWeight: 1} // 1.5
	input2 := Neuron{Value: 2, Bias: 4, BiasWeight: 1} // 3

	out1Connections := []*Connection{
		&Connection{InNeuron: &input1, Weight: 1}, // 1.5
		&Connection{InNeuron: &input2, Weight: 1}} // 3
	out2Connections := []*Connection{
		&Connection{InNeuron: &input1, Weight: 2}, // 3
		&Connection{InNeuron: &input2, Weight: 1}} // 3

	output1 := Neuron{Bias: 1, BiasWeight: 2, InConnections: out1Connections} // 1.625
	output2 := Neuron{Bias: 1, BiasWeight: 1, InConnections: out2Connections} // 1.75

	net := Net{InputNeurons: []*Neuron{&input1, &input2}, OutputNeurons: []*Neuron{&output1, &output2}}

	// when
	net.CalculateOutputs()

	// then
	assert.Equal(t, 1.625, net.OutputNeurons[0].Value)
	assert.Equal(t, 1.75, net.OutputNeurons[1].Value)
}
