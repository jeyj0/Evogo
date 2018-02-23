package world_test

import (
	"testing"

	"github.com/jeyj0/Evogo/world"
	"github.com/stretchr/testify/assert"
)

func TestRad(t *testing.T) {
	// given
	angle := world.Angle{90}

	// when
	rad := angle.Rad()

	// then
	assert.Equal(t, 1.5707963267948966, rad)
}
