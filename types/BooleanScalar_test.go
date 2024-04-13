package types_test

import (
	"testing"

	"github.com/michaelcmelton/raml-parser/types"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_BooleanScalar_Shorthand(t *testing.T) {
	sampleRaml := `SampleBoolean: boolean`
	var booleanScalar *types.BooleanScalar
	yaml.Unmarshal([]byte(sampleRaml), &booleanScalar)

	assert.Equal(t, booleanScalar.Name, "SampleBoolean")
	assert.Equal(t, booleanScalar.Type, types.BOOLEAN)
}

func Test_BooleanScalar_Longform(t *testing.T) {
	sampleRaml := `SampleBoolean:
      type: boolean`
	var booleanScalar *types.BooleanScalar
	err := yaml.Unmarshal([]byte(sampleRaml), &booleanScalar)
	assert.NoError(t, err)
	assert.Equal(t, booleanScalar.Name, "SampleBoolean")
	assert.Equal(t, booleanScalar.Type, types.BOOLEAN)
}
