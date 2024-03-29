package test_random

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/random"
	"github.com/stretchr/testify/assert"
)

func TestNextDouble(t *testing.T) {
	value := random.Double.Next(0, 5)
	assert.True(t, value <= 5)

	value = random.Double.Next(2, 5)
	assert.True(t, value <= 5 && value >= 2)
}

func TestUpdateDouble(t *testing.T) {
	value := random.Double.Update(0, 5)
	assert.True(t, value <= 5 && value >= -5)

	value = random.Double.Update(5, 0)

	value = random.Double.Update(0, 0)
	assert.True(t, value == 0)
}
