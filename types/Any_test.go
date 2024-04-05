package types_test

import (
	"testing"

	"github.com/michaelcmelton/raml-parser/types"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_Any(t *testing.T) {
	sampleRaml := `ExampleType: any`

	var anyType types.AnyType
	err := yaml.Unmarshal([]byte(sampleRaml), &anyType)
	assert.NoError(t, err)

	assert.Equal(t, anyType.Type, types.ANY)
	assert.Equal(t, anyType.Name, "ExampleType")
}
