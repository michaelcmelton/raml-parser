package types_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/michaelcmelton/raml-parser/types"
	"gopkg.in/yaml.v3"
)

func Test_StringScalar_Full(t *testing.T) {
	sampleRAML := `EmailAddress:
    type: string
    pattern: ^.+@.+\..+$
    minLength: 3
    maxLength: 320`

	var stringScalar types.StringScalar

	err := yaml.Unmarshal([]byte(sampleRAML), &stringScalar)

	assert.NoError(t, err)
	assert.Equal(t, 320, stringScalar.MaxLength)
	assert.Equal(t, 3, stringScalar.MinLength)
	assert.NotNil(t, stringScalar.Pattern)
}

func Test_StringScalar_Abbreviated(t *testing.T) {
	sampleRAML := `EmailAddress: string`

	var stringScalar types.StringScalar

	err := yaml.Unmarshal([]byte(sampleRAML), &stringScalar)

	assert.NoError(t, err)
	assert.Equal(t, "string", stringScalar.Type)
	assert.Equal(t, 0, stringScalar.MinLength)
	assert.Equal(t, 2147483647, stringScalar.MaxLength)
	assert.Nil(t, stringScalar.Pattern)
}

func Test_StringScalar_InvalidRAML(t *testing.T) {
	sampleRAML := `EmailAddress:
    type: string
    pattern: ^.+@.+\..+$
    minLength: -1
    maxLength: 320`

	var stringScalar types.StringScalar

	err := yaml.Unmarshal([]byte(sampleRAML), &stringScalar)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "minimum length, if provided, must be equal to or greater than zero")
}
