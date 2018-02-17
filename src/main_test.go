package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateValue(t *testing.T) {
	// given
	input1 := Neuron{value: 2, bias: 0, biasWeight: 0, inConnections: nil}
	input2 := Neuron{value: 6, bias: 0, biasWeight: 0, inConnections: nil}

	inConnections := []*Connection{
		&Connection{inNeuron: &input1, weight: 1},
		&Connection{inNeuron: &input2, weight: 2}}

	outNeuron := Neuron{value: 0, bias: 5, biasWeight: 1, inConnections: inConnections}

	// when
	outNeuron.calculateValue()

	// then
	if assert.NotNil(t, outNeuron.value) {
		assert.Equal(t, 4.75, outNeuron.value)
	}
}

func TestCalculateValueRecursive(t *testing.T) {
	// given
	input := Neuron{value: 1, bias: 0, biasWeight: 0, inConnections: nil}
	middle := Neuron{value: 0, bias: 2, biasWeight: 4, inConnections: []*Connection{&Connection{inNeuron: &input, weight: 4}}}
	output := Neuron{value: 0, bias: 5, biasWeight: 6, inConnections: []*Connection{&Connection{inNeuron: &middle, weight: 8}}}

	// (1*4 + 2*4) / (4 + 4) = 1.5 // middle
	// (1.5*8 + 5*6) / (8 + 6) = 3 // output

	// when
	output.calculateValueRecursive()

	// then
	assert.Equal(t, 3.0, output.value)
}

func TestCalculateOutputs(t *testing.T) {
	// given
	input1 := Neuron{value: 1, bias: 0, biasWeight: 0, inConnections: nil}
	input2 := Neuron{value: 2, bias: 0, biasWeight: 0, inConnections: nil}

	out1Connections := []*Connection{
		&Connection{inNeuron: &input1, weight: 3},
		&Connection{inNeuron: &input2, weight: 5}}

	out2Connections := []*Connection{
		&Connection{inNeuron: &input1, weight: 4},
		&Connection{inNeuron: &input2, weight: 6}}

	output1 := Neuron{value: 0, bias: 7, biasWeight: 8, inConnections: out1Connections}
	output2 := Neuron{value: 0, bias: 9, biasWeight: 10, inConnections: out2Connections}

	net := Net{inputNeurons: []*Neuron{&input1, &input2}, outputNeurons: []*Neuron{&output1, &output2}}

	// when
	net.calculateOutputs()

	// then
	assert.Equal(t, 4.3125, net.outputNeurons[0].value)
	assert.Equal(t, 5.3, net.outputNeurons[1].value)
}
