package test_random

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/random"
	"github.com/stretchr/testify/assert"
)

func TestNextFloat(t *testing.T) {
	value := random.Float.Next(0, 5)
	assert.True(t, value <= 5)

	value = random.Float.Next(2, 5)
	assert.True(t, value <= 5 && value >= 2)
}

func TestUpdateFloat(t *testing.T) {
	value := random.Float.Update(0, 5)
	assert.True(t, value <= 5 && value >= -5)

	value = random.Float.Update(5, 0)

	value = random.Float.Update(0, 0)
	assert.True(t, value == 0)
}
