package neuralnet_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"

// 	nn "github.com/jeyj0/Evogo/neuralnet"
// )

// func TestCalculateValue(t *testing.T) {
// 	// given
// 	input1 := nn.Neuron{Value: 2}
// 	input2 := nn.Neuron{Value: 6}

// 	inConnections := []*nn.Connection{
// 		&nn.Connection{InNeuron: &input1, Weight: 1},
// 		&nn.Connection{InNeuron: &input2, Weight: 2}}

// 	outNeuron := nn.Neuron{Bias: 5, BiasWeight: 1, InConnections: inConnections}

// 	// when
// 	outNeuron.calculateValue()

// 	// then
// 	assert.Equal(t, 4.75, outNeuron.Value)
// }

// func TestCalculateValueInputNeuron(t *testing.T) {
// 	// given
// 	neuron := nn.Neuron{Value: 4, Bias: 1, BiasWeight: 2}

// 	// when
// 	neuron.calculateValue()

// 	// then
// 	assert.Equal(t, 2.0, neuron.Value)
// }

// func TestCalculateValueRecursive(t *testing.T) {
// 	// given
// 	input := nn.Neuron{Value: 1}
// 	middle := nn.Neuron{Bias: 2, BiasWeight: 4, InConnections: []*nn.Connection{&nn.Connection{InNeuron: &input, Weight: 4}}}
// 	output := nn.Neuron{Bias: 5, BiasWeight: 6, InConnections: []*nn.Connection{&nn.Connection{InNeuron: &middle, Weight: 8}}}

// 	// (1*4 + 2*4) / (4 + 4) = 1.5 // middle
// 	// (1.5*8 + 5*6) / (8 + 6) = 3 // output

// 	// when
// 	output.calculateValueRecursive()

// 	// then
// 	assert.Equal(t, 3.0, output.Value)
// }

// func TestCalculateValueRecursiveWithInputBias(t *testing.T) {
// 	// given
// 	input := nn.Neuron{Value: 1, Bias: 2, BiasWeight: 4}
// 	middle := nn.Neuron{Bias: 6, BiasWeight: 10, InConnections: []*nn.Connection{&nn.Connection{InNeuron: &input, Weight: 10}}}
// 	output := nn.Neuron{Bias: 3, BiasWeight: 14, InConnections: []*nn.Connection{&nn.Connection{InNeuron: &middle, Weight: 16}}}

// 	// input: 1.8
// 	// middle: 3.9
// 	// output: 3.48

// 	// when
// 	output.calculateValueRecursive()

// 	// then
// 	assert.Equal(t, 3.48, output.Value)
// }
