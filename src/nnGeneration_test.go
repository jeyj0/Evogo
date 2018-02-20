package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFullyConnectedNeuralNet(t *testing.T) {
	// given
	input1 := Neuron{}
	input2 := Neuron{}
	middle1 := Neuron{inConnections: []*Connection{&Connection{inNeuron: &input1}, &Connection{inNeuron: &input2}}}
	middle2 := Neuron{inConnections: []*Connection{&Connection{inNeuron: &input1}, &Connection{inNeuron: &input2}}}
	out1 := Neuron{inConnections: []*Connection{&Connection{inNeuron: &middle1}, &Connection{inNeuron: &middle2}}}
	out2 := Neuron{inConnections: []*Connection{&Connection{inNeuron: &middle1}, &Connection{inNeuron: &middle2}}}
	fullyConnected2x2x2 := Net{inputNeurons: []*Neuron{&input1, &input2}, outputNeurons: []*Neuron{&out1, &out2}}

	// when
	generated2x2x2 := GenerateFullyConnectedNeuralNet([]int{2, 2, 2})

	// then
	assert.Equal(t, fullyConnected2x2x2, generated2x2x2)
}
