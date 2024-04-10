package types_test

import (
	"testing"

	"github.com/michaelcmelton/raml-parser/types"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_IntegerScalar_NoFacets(t *testing.T) {
	sampleRaml := `SampleNumber: integer`

	var integerScalar *types.IntegerScalar

	err := yaml.Unmarshal([]byte(sampleRaml), &integerScalar)
	assert.NoError(t, err)
	assert.Equal(t, integerScalar.Minimum, float64(0))
	assert.Equal(t, integerScalar.Maximum, float64(0))
}
