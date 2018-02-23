package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFullyConnectedNeuralNet1x1(t *testing.T) {
	// given
	input := Neuron{}
	output := Neuron{}
	conn := Connection{inNeuron: &input, outNeuron: &output}
	input.outConnections = []*Connection{&conn}
	output.inConnections = []*Connection{&conn}
	fullyConnected1x1 := Net{inputNeurons: []*Neuron{&input}, outputNeurons: []*Neuron{&output}}

	// when
	generated1x1 := GenerateFullyConnectedNeuralNet([]int{1, 1})

	// then
	assert.Equal(t, fullyConnected1x1, generated1x1)
}

func TestGenerateFullyConnectedNeuralNet1x2(t *testing.T) {
	// given
	input := Neuron{}
	out1 := Neuron{}
	out2 := Neuron{}
	conn1 := Connection{inNeuron: &input, outNeuron: &out1}
	conn2 := Connection{inNeuron: &input, outNeuron: &out2}
	input.outConnections = []*Connection{&conn1, &conn2}
	out1.inConnections = []*Connection{&conn1}
	out2.inConnections = []*Connection{&conn2}
	fullyConnected1x2 := Net{inputNeurons: []*Neuron{&input}, outputNeurons: []*Neuron{&out1, &out2}}

	// when
	generated1x2 := GenerateFullyConnectedNeuralNet([]int{1, 2})

	// then
	assert.Equal(t, fullyConnected1x2, generated1x2)
}

func TestGenerateFullyConnectedNeuralNet2x1(t *testing.T) {
	// given
	input1 := Neuron{}
	input2 := Neuron{}
	output := Neuron{}
	conn1 := Connection{inNeuron: &input1, outNeuron: &output}
	conn2 := Connection{inNeuron: &input2, outNeuron: &output}
	input1.outConnections = []*Connection{&conn1}
	input2.outConnections = []*Connection{&conn2}
	output.inConnections = []*Connection{&conn1, &conn2}
	fullyConnected2x1 := Net{inputNeurons: []*Neuron{&input1, &input2}, outputNeurons: []*Neuron{&output}}

	// when
	generated2x1 := GenerateFullyConnectedNeuralNet([]int{2, 1})

	// then
	assert.Equal(t, fullyConnected2x1, generated2x1)
}

func TestGenerateFullyConnectedNeuralNet1x1x1(t *testing.T) {
	// given
	input := Neuron{}
	middle := Neuron{}
	output := Neuron{}
	conn1 := Connection{inNeuron: &input, outNeuron: &middle}
	conn2 := Connection{inNeuron: &middle, outNeuron: &output}
	input.outConnections = []*Connection{&conn1}
	middle.inConnections = []*Connection{&conn1}
	middle.outConnections = []*Connection{&conn2}
	output.inConnections = []*Connection{&conn2}
	fullyConnected1x1x1 := Net{inputNeurons: []*Neuron{&input}, outputNeurons: []*Neuron{&output}}

	// when
	generated1x1x1 := GenerateFullyConnectedNeuralNet([]int{1, 1, 1})

	// then
	assert.Equal(t, 1, len(generated1x1x1.inputNeurons), "A 1x1x1 Net should have 1 input neuron")
	assert.Equal(t, 1, len(generated1x1x1.outputNeurons), "A 1x1x1 Net should have 1 output neuron")
	assert.Equal(t, 1, len(generated1x1x1.inputNeurons[0].outConnections), "A 1x1x1 Net should have 1 connection from input to middle")
	assert.Equal(t, 1, len(generated1x1x1.outputNeurons[0].inConnections), "A 1x1x1 Net should have 1 connection from middle to output")
	assert.Equal(t, fullyConnected1x1x1, generated1x1x1)
}

func TestGenerateFullyConnectedNeuralNet2x2x2(t *testing.T) {
	// given
	fullyConnected2x2x2 := givenFullyConnected2x2x2()

	// when
	generated2x2x2 := GenerateFullyConnectedNeuralNet([]int{2, 2, 2})

	// then
	assert.Equal(t, fullyConnected2x2x2, generated2x2x2)
}

func givenFullyConnected2x2x2() Net {
	input1 := Neuron{}
	input2 := Neuron{}
	middle1 := Neuron{}
	middle2 := Neuron{}
	out1 := Neuron{}
	out2 := Neuron{}

	connI1M1 := &Connection{inNeuron: &input1, outNeuron: &middle1}
	connI1M2 := &Connection{inNeuron: &input1, outNeuron: &middle2}
	connI2M1 := &Connection{inNeuron: &input2, outNeuron: &middle1}
	connI2M2 := &Connection{inNeuron: &input2, outNeuron: &middle2}
	connM1O1 := &Connection{inNeuron: &middle1, outNeuron: &out1}
	connM1O2 := &Connection{inNeuron: &middle1, outNeuron: &out2}
	connM2O1 := &Connection{inNeuron: &middle2, outNeuron: &out1}
	connM2O2 := &Connection{inNeuron: &middle2, outNeuron: &out2}

	input1.outConnections = []*Connection{connI1M1, connI1M2}
	input2.outConnections = []*Connection{connI2M1, connI2M2}
	middle1.outConnections = []*Connection{connM1O1, connM1O2}
	middle2.outConnections = []*Connection{connM2O1, connM2O2}

	middle1.inConnections = []*Connection{connI1M1, connI2M1}
	middle2.inConnections = []*Connection{connI1M2, connI2M2}
	out1.inConnections = []*Connection{connM1O1, connM2O1}
	out2.inConnections = []*Connection{connM1O2, connM2O2}

	return Net{inputNeurons: []*Neuron{&input1, &input2}, outputNeurons: []*Neuron{&out1, &out2}}
}

func TestFillWeightsAndBiasesFromSeed(t *testing.T) {
	// given
	seed := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fullyConnected2x2x2 := givenFullyConnectedAndBiasWeightFilled2x2x2()

	// when
	filled2x2x2 := GenerateFullyConnectedNeuralNet([]int{2, 2, 2})
	filled2x2x2.FillWeightsAndBiasesFromSeed(seed)

	// then
	assert.Equal(t, fullyConnected2x2x2, filled2x2x2)
}

func givenFullyConnectedAndBiasWeightFilled2x2x2() Net {
	input1 := Neuron{bias: 1, biasWeight: 2}
	input2 := Neuron{bias: 3, biasWeight: 4}
	middle1 := Neuron{bias: 9, biasWeight: 10}
	middle2 := Neuron{bias: 11, biasWeight: 12}
	out1 := Neuron{bias: 17, biasWeight: 18}
	out2 := Neuron{bias: 19, biasWeight: 20}

	connI1M1 := &Connection{inNeuron: &input1, outNeuron: &middle1, weight: 5}
	connI1M2 := &Connection{inNeuron: &input1, outNeuron: &middle2, weight: 6}
	connI2M1 := &Connection{inNeuron: &input2, outNeuron: &middle1, weight: 7}
	connI2M2 := &Connection{inNeuron: &input2, outNeuron: &middle2, weight: 8}
	connM1O1 := &Connection{inNeuron: &middle1, outNeuron: &out1, weight: 13}
	connM1O2 := &Connection{inNeuron: &middle1, outNeuron: &out2, weight: 14}
	connM2O1 := &Connection{inNeuron: &middle2, outNeuron: &out1, weight: 15}
	connM2O2 := &Connection{inNeuron: &middle2, outNeuron: &out2, weight: 16}

	input1.outConnections = []*Connection{connI1M1, connI1M2}
	input2.outConnections = []*Connection{connI2M1, connI2M2}
	middle1.outConnections = []*Connection{connM1O1, connM1O2}
	middle2.outConnections = []*Connection{connM2O1, connM2O2}

	middle1.inConnections = []*Connection{connI1M1, connI2M1}
	middle2.inConnections = []*Connection{connI1M2, connI2M2}
	out1.inConnections = []*Connection{connM1O1, connM2O1}
	out2.inConnections = []*Connection{connM1O2, connM2O2}

	return Net{inputNeurons: []*Neuron{&input1, &input2}, outputNeurons: []*Neuron{&out1, &out2}}
}
