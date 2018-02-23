package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateValue(t *testing.T) {
	// given
	input1 := Neuron{value: 2}
	input2 := Neuron{value: 6}

	inConnections := []*Connection{
		&Connection{inNeuron: &input1, weight: 1},
		&Connection{inNeuron: &input2, weight: 2}}

	outNeuron := Neuron{bias: 5, biasWeight: 1, inConnections: inConnections}

	// when
	outNeuron.calculateValue()

	// then
	assert.Equal(t, 4.75, outNeuron.value)
}

func TestCalculateValueInputNeuron(t *testing.T) {
	// given
	neuron := Neuron{value: 4, bias: 1, biasWeight: 2}

	// when
	neuron.calculateValue()

	// then
	assert.Equal(t, 2.0, neuron.value)
}

func TestCalculateValueRecursive(t *testing.T) {
	// given
	input := Neuron{value: 1}
	middle := Neuron{bias: 2, biasWeight: 4, inConnections: []*Connection{&Connection{inNeuron: &input, weight: 4}}}
	output := Neuron{bias: 5, biasWeight: 6, inConnections: []*Connection{&Connection{inNeuron: &middle, weight: 8}}}

	// (1*4 + 2*4) / (4 + 4) = 1.5 // middle
	// (1.5*8 + 5*6) / (8 + 6) = 3 // output

	// when
	output.calculateValueRecursive()

	// then
	assert.Equal(t, 3.0, output.value)
}

func TestCalculateValueRecursiveWithInputBias(t *testing.T) {
	// given
	input := Neuron{value: 1, bias: 2, biasWeight: 4}
	middle := Neuron{bias: 6, biasWeight: 10, inConnections: []*Connection{&Connection{inNeuron: &input, weight: 10}}}
	output := Neuron{bias: 3, biasWeight: 14, inConnections: []*Connection{&Connection{inNeuron: &middle, weight: 16}}}

	// input: 1.8
	// middle: 3.9
	// output: 3.48

	// when
	output.calculateValueRecursive()

	// then
	assert.Equal(t, 3.48, output.value)
}

func TestCalculateOutputs(t *testing.T) {
	// given
	input1 := Neuron{value: 1}
	input2 := Neuron{value: 2}

	out1Connections := []*Connection{
		&Connection{inNeuron: &input1, weight: 3},
		&Connection{inNeuron: &input2, weight: 5}}

	out2Connections := []*Connection{
		&Connection{inNeuron: &input1, weight: 4},
		&Connection{inNeuron: &input2, weight: 6}}

	output1 := Neuron{bias: 7, biasWeight: 8, inConnections: out1Connections}
	output2 := Neuron{bias: 9, biasWeight: 10, inConnections: out2Connections}

	net := Net{inputNeurons: []*Neuron{&input1, &input2}, outputNeurons: []*Neuron{&output1, &output2}}

	// when
	net.CalculateOutputs()

	// then
	assert.Equal(t, 4.3125, net.outputNeurons[0].value)
	assert.Equal(t, 5.3, net.outputNeurons[1].value)
}

func TestCalculateOutputsWithInputBias(t *testing.T) {
	// given
	input1 := Neuron{value: 1, bias: 2, biasWeight: 1} // 1.5
	input2 := Neuron{value: 2, bias: 4, biasWeight: 1} // 3

	out1Connections := []*Connection{
		&Connection{inNeuron: &input1, weight: 1}, // 1.5
		&Connection{inNeuron: &input2, weight: 1}} // 3
	out2Connections := []*Connection{
		&Connection{inNeuron: &input1, weight: 2}, // 3
		&Connection{inNeuron: &input2, weight: 1}} // 3

	output1 := Neuron{bias: 1, biasWeight: 2, inConnections: out1Connections} // 1.625
	output2 := Neuron{bias: 1, biasWeight: 1, inConnections: out2Connections} // 1.75

	net := Net{inputNeurons: []*Neuron{&input1, &input2}, outputNeurons: []*Neuron{&output1, &output2}}

	// when
	net.CalculateOutputs()

	// then
	assert.Equal(t, 1.625, net.outputNeurons[0].value)
	assert.Equal(t, 1.75, net.outputNeurons[1].value)
}
