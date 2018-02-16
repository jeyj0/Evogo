package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateValue(t *testing.T) {
	// given
	input1 := Neuron{value: 2.0, bias: 0, biasWeight: 0, inConnections: nil}
	input2 := Neuron{value: 6.0, bias: 0, biasWeight: 0, inConnections: nil}

	var inConnections []*Connection
	inConnections = append(inConnections, &Connection{inNeuron: &input1, weight: 1.0})
	inConnections = append(inConnections, &Connection{inNeuron: &input2, weight: 2.0})

	outNeuron := Neuron{value: 0.0, bias: 5.0, biasWeight: 1.0, inConnections: inConnections}

	// when
	outNeuron.calculateValue()

	// then
	if assert.NotNil(t, outNeuron.value) {
		assert.Equal(t, 4.75, outNeuron.value)
	}
}

func TestCalculateOutputs(t *testing.T) {
	// TODO
}
