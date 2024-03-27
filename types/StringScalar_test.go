package types_test

import (
	"testing"

	"github.com/michaelcmelton/raml-parser/types"
	"gopkg.in/yaml.v3"
)

func Test_StringScalar(t *testing.T) {
	sampleRAML := `EmailAddress:
    type: string
    pattern: ^.+@.+\..+$
    minLength: 3
    maxLength: 320`

	var stringScalar types.StringScalar

	err := yaml.Unmarshal([]byte(sampleRAML), &stringScalar)

	if err != nil {
		t.Fatalf("error unmarshaling string scalar: %s", err)
	}

	if stringScalar.MaxLength != 320 {
		t.Fatalf("expected %d, got %d", 320, stringScalar.MaxLength)
	}

	if stringScalar.MinLength != 3 {
		t.Fatalf("expected %d, got %d", 3, stringScalar.MinLength)
	}
}
