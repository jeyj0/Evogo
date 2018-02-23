package neuralnet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateValue(t *testing.T) {
	// given
	input1 := Neuron{Value: 2}
	input2 := Neuron{Value: 6}

	inConnections := []*Connection{
		&Connection{InNeuron: &input1, Weight: 1},
		&Connection{InNeuron: &input2, Weight: 2}}

	outNeuron := Neuron{Bias: 5, BiasWeight: 1, InConnections: inConnections}

	// when
	outNeuron.calculateValue()

	// then
	assert.Equal(t, 4.75, outNeuron.Value)
}

func TestCalculateValueInputNeuron(t *testing.T) {
	// given
	neuron := Neuron{Value: 4, Bias: 1, BiasWeight: 2}

	// when
	neuron.calculateValue()

	// then
	assert.Equal(t, 2.0, neuron.Value)
}

func TestCalculateValueRecursive(t *testing.T) {
	// given
	input := Neuron{Value: 1}
	middle := Neuron{Bias: 2, BiasWeight: 4, InConnections: []*Connection{&Connection{InNeuron: &input, Weight: 4}}}
	output := Neuron{Bias: 5, BiasWeight: 6, InConnections: []*Connection{&Connection{InNeuron: &middle, Weight: 8}}}

	// (1*4 + 2*4) / (4 + 4) = 1.5 // middle
	// (1.5*8 + 5*6) / (8 + 6) = 3 // output

	// when
	output.calculateValueRecursive()

	// then
	assert.Equal(t, 3.0, output.Value)
}

func TestCalculateValueRecursiveWithInputBias(t *testing.T) {
	// given
	input := Neuron{Value: 1, Bias: 2, BiasWeight: 4}
	middle := Neuron{Bias: 6, BiasWeight: 10, InConnections: []*Connection{&Connection{InNeuron: &input, Weight: 10}}}
	output := Neuron{Bias: 3, BiasWeight: 14, InConnections: []*Connection{&Connection{InNeuron: &middle, Weight: 16}}}

	// input: 1.8
	// middle: 3.9
	// output: 3.48

	// when
	output.calculateValueRecursive()

	// then
	assert.Equal(t, 3.48, output.Value)
}
