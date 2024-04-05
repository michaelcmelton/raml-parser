package types_test

import (
	"testing"

	"github.com/michaelcmelton/raml-parser/types"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_NumberScalar_AllFacets(t *testing.T) {
	sampleRaml := `Weight:
    type: number
    minimum: -1.1
    maximum: 20.9
    format: float
    multipleOf: 1.1`
	var number *types.NumberScalar

	err := yaml.Unmarshal([]byte(sampleRaml), &number)

	assert.NoError(t, err)

	assert.Equal(t, number.Name, "Weight")
	assert.Equal(t, number.Format, "float")
	assert.Equal(t, number.MultipleOf, 1.1)
	assert.Equal(t, number.Minimum, -1.1)
	assert.Equal(t, number.Maximum, 20.9)
}

func Test_NumberScalar_SomeFacets(t *testing.T) {
	sampleRaml := `Weight:
    type: number
    format: float
    multipleOf: 1.1`
	var number *types.NumberScalar

	err := yaml.Unmarshal([]byte(sampleRaml), &number)

	assert.NoError(t, err)

	assert.Equal(t, number.Name, "Weight")
	assert.Equal(t, number.Format, "float")
	assert.Equal(t, number.MultipleOf, 1.1)
	assert.Equal(t, number.Minimum, float64(0))
	assert.Equal(t, number.Maximum, float64(0))
}

func Test_NumberScalar_NoFacets(t *testing.T) {
	sampleRaml := `Weight: number`
	var number *types.NumberScalar

	err := yaml.Unmarshal([]byte(sampleRaml), &number)

	assert.NoError(t, err)

	assert.Equal(t, number.Name, "Weight")
	assert.Equal(t, number.Format, "")
	assert.Equal(t, number.MultipleOf, float64(0))
	assert.Equal(t, number.Minimum, float64(0))
	assert.Equal(t, number.Maximum, float64(0))
}
