package types_test

import (
	"testing"

	"github.com/michaelcmelton/raml-parser/types"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_DateScalar_ValidShortand(t *testing.T) {
	sampleRaml := `birthday: date-only`

	var birthday *types.DateScalar
	err := yaml.Unmarshal([]byte(sampleRaml), &birthday)

	assert.NoError(t, err)
	assert.Equal(t, birthday.Name, "birthday")
	assert.Equal(t, birthday.Type, types.DATEONLY)
}
