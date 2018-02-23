package nngeneration

import (
	"testing"

	nn "github.com/jeyj0/Evogo/neuralnet"
	"github.com/stretchr/testify/assert"
)

func TestGenerateFullyConnectedNeuralNet1x1(t *testing.T) {
	// given
	input := nn.Neuron{}
	output := nn.Neuron{}
	conn := nn.Connection{InNeuron: &input, OutNeuron: &output}
	input.OutConnections = []*nn.Connection{&conn}
	output.InConnections = []*nn.Connection{&conn}
	fullyConnected1x1 := nn.Net{InputNeurons: []*nn.Neuron{&input}, OutputNeurons: []*nn.Neuron{&output}}

	// when
	generated1x1 := GenerateFullyConnectedNeuralNet([]int{1, 1})

	// then
	assert.Equal(t, fullyConnected1x1, generated1x1)
}

func TestGenerateFullyConnectedNeuralNet1x2(t *testing.T) {
	// given
	input := nn.Neuron{}
	out1 := nn.Neuron{}
	out2 := nn.Neuron{}
	conn1 := nn.Connection{InNeuron: &input, OutNeuron: &out1}
	conn2 := nn.Connection{InNeuron: &input, OutNeuron: &out2}
	input.OutConnections = []*nn.Connection{&conn1, &conn2}
	out1.InConnections = []*nn.Connection{&conn1}
	out2.InConnections = []*nn.Connection{&conn2}
	fullyConnected1x2 := nn.Net{InputNeurons: []*nn.Neuron{&input}, OutputNeurons: []*nn.Neuron{&out1, &out2}}

	// when
	generated1x2 := GenerateFullyConnectedNeuralNet([]int{1, 2})

	// then
	assert.Equal(t, fullyConnected1x2, generated1x2)
}

func TestGenerateFullyConnectedNeuralNet2x1(t *testing.T) {
	// given
	input1 := nn.Neuron{}
	input2 := nn.Neuron{}
	output := nn.Neuron{}
	conn1 := nn.Connection{InNeuron: &input1, OutNeuron: &output}
	conn2 := nn.Connection{InNeuron: &input2, OutNeuron: &output}
	input1.OutConnections = []*nn.Connection{&conn1}
	input2.OutConnections = []*nn.Connection{&conn2}
	output.InConnections = []*nn.Connection{&conn1, &conn2}
	fullyConnected2x1 := nn.Net{InputNeurons: []*nn.Neuron{&input1, &input2}, OutputNeurons: []*nn.Neuron{&output}}

	// when
	generated2x1 := GenerateFullyConnectedNeuralNet([]int{2, 1})

	// then
	assert.Equal(t, fullyConnected2x1, generated2x1)
}

func TestGenerateFullyConnectedNeuralNet1x1x1(t *testing.T) {
	// given
	input := nn.Neuron{}
	middle := nn.Neuron{}
	output := nn.Neuron{}
	conn1 := nn.Connection{InNeuron: &input, OutNeuron: &middle}
	conn2 := nn.Connection{InNeuron: &middle, OutNeuron: &output}
	input.OutConnections = []*nn.Connection{&conn1}
	middle.InConnections = []*nn.Connection{&conn1}
	middle.OutConnections = []*nn.Connection{&conn2}
	output.InConnections = []*nn.Connection{&conn2}
	fullyConnected1x1x1 := nn.Net{InputNeurons: []*nn.Neuron{&input}, OutputNeurons: []*nn.Neuron{&output}}

	// when
	generated1x1x1 := GenerateFullyConnectedNeuralNet([]int{1, 1, 1})

	// then
	assert.Equal(t, 1, len(generated1x1x1.InputNeurons), "A 1x1x1 Net should have 1 input neuron")
	assert.Equal(t, 1, len(generated1x1x1.OutputNeurons), "A 1x1x1 Net should have 1 output neuron")
	assert.Equal(t, 1, len(generated1x1x1.InputNeurons[0].OutConnections), "A 1x1x1 Net should have 1 connection from input to middle")
	assert.Equal(t, 1, len(generated1x1x1.OutputNeurons[0].InConnections), "A 1x1x1 Net should have 1 connection from middle to output")
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

func givenFullyConnected2x2x2() nn.Net {
	input1 := nn.Neuron{}
	input2 := nn.Neuron{}
	middle1 := nn.Neuron{}
	middle2 := nn.Neuron{}
	out1 := nn.Neuron{}
	out2 := nn.Neuron{}

	connI1M1 := &nn.Connection{InNeuron: &input1, OutNeuron: &middle1}
	connI1M2 := &nn.Connection{InNeuron: &input1, OutNeuron: &middle2}
	connI2M1 := &nn.Connection{InNeuron: &input2, OutNeuron: &middle1}
	connI2M2 := &nn.Connection{InNeuron: &input2, OutNeuron: &middle2}
	connM1O1 := &nn.Connection{InNeuron: &middle1, OutNeuron: &out1}
	connM1O2 := &nn.Connection{InNeuron: &middle1, OutNeuron: &out2}
	connM2O1 := &nn.Connection{InNeuron: &middle2, OutNeuron: &out1}
	connM2O2 := &nn.Connection{InNeuron: &middle2, OutNeuron: &out2}

	input1.OutConnections = []*nn.Connection{connI1M1, connI1M2}
	input2.OutConnections = []*nn.Connection{connI2M1, connI2M2}
	middle1.OutConnections = []*nn.Connection{connM1O1, connM1O2}
	middle2.OutConnections = []*nn.Connection{connM2O1, connM2O2}

	middle1.InConnections = []*nn.Connection{connI1M1, connI2M1}
	middle2.InConnections = []*nn.Connection{connI1M2, connI2M2}
	out1.InConnections = []*nn.Connection{connM1O1, connM2O1}
	out2.InConnections = []*nn.Connection{connM1O2, connM2O2}

	return nn.Net{InputNeurons: []*nn.Neuron{&input1, &input2}, OutputNeurons: []*nn.Neuron{&out1, &out2}}
}

func TestFillWeightsAndBiasesFromSeed(t *testing.T) {
	// given
	seed := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fullyConnected2x2x2 := givenFullyConnectedAndBiasWeightFilled2x2x2()

	// when
	filled2x2x2 := GenerateFullyConnectedNeuralNet([]int{2, 2, 2})
	FillWeightsAndBiasesFromSeed(&filled2x2x2, seed)

	// then
	assert.Equal(t, fullyConnected2x2x2, filled2x2x2)
}

func givenFullyConnectedAndBiasWeightFilled2x2x2() nn.Net {
	input1 := nn.Neuron{Bias: 1, BiasWeight: 2}
	input2 := nn.Neuron{Bias: 3, BiasWeight: 4}
	middle1 := nn.Neuron{Bias: 9, BiasWeight: 10}
	middle2 := nn.Neuron{Bias: 11, BiasWeight: 12}
	out1 := nn.Neuron{Bias: 17, BiasWeight: 18}
	out2 := nn.Neuron{Bias: 19, BiasWeight: 20}

	connI1M1 := &nn.Connection{InNeuron: &input1, OutNeuron: &middle1, Weight: 5}
	connI1M2 := &nn.Connection{InNeuron: &input1, OutNeuron: &middle2, Weight: 6}
	connI2M1 := &nn.Connection{InNeuron: &input2, OutNeuron: &middle1, Weight: 7}
	connI2M2 := &nn.Connection{InNeuron: &input2, OutNeuron: &middle2, Weight: 8}
	connM1O1 := &nn.Connection{InNeuron: &middle1, OutNeuron: &out1, Weight: 13}
	connM1O2 := &nn.Connection{InNeuron: &middle1, OutNeuron: &out2, Weight: 14}
	connM2O1 := &nn.Connection{InNeuron: &middle2, OutNeuron: &out1, Weight: 15}
	connM2O2 := &nn.Connection{InNeuron: &middle2, OutNeuron: &out2, Weight: 16}

	input1.OutConnections = []*nn.Connection{connI1M1, connI1M2}
	input2.OutConnections = []*nn.Connection{connI2M1, connI2M2}
	middle1.OutConnections = []*nn.Connection{connM1O1, connM1O2}
	middle2.OutConnections = []*nn.Connection{connM2O1, connM2O2}

	middle1.InConnections = []*nn.Connection{connI1M1, connI2M1}
	middle2.InConnections = []*nn.Connection{connI1M2, connI2M2}
	out1.InConnections = []*nn.Connection{connM1O1, connM2O1}
	out2.InConnections = []*nn.Connection{connM1O2, connM2O2}

	return nn.Net{InputNeurons: []*nn.Neuron{&input1, &input2}, OutputNeurons: []*nn.Neuron{&out1, &out2}}
}
