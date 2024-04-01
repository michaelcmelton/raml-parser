package types

import (
	"fmt"
	"regexp"

	"gopkg.in/yaml.v3"
)

type StringScalar struct {
	BaseType
	Pattern   *regexp.Regexp
	MinLength int
	MaxLength int
}

type stringScalarFacets struct {
	Pattern   string `yaml:"pattern,omitempty"`
	MinLength int    `yaml:"minLength,omitempty"`
	MaxLength int    `yaml:"maxLength,omitempty"`
}

func (s *StringScalar) UnmarshalYAML(node *yaml.Node) error {

	// Set Defaults according to string scalar spec
	s.MinLength = 0
	s.MaxLength = 2147483647

	s.Name = node.Content[0].Value
	s.Type = STRING
	properties := node.Content[1]

	if properties.Tag == "!!str" {
		return nil
	}

	var facets *stringScalarFacets

	err := properties.Decode(&facets)

	if err != nil {
		return err
	}

	if facets.MinLength < 0 {
		return fmt.Errorf(ErrValidation, s.Name, "minimum length, if provided, must be equal to or greater than zero")
	}

	if facets.MaxLength < 0 {
		return fmt.Errorf(ErrValidation, s.Name, "maximum length, if provided, must be equal to or greater than zero")
	}

	if facets.MinLength != 0 {
		s.MinLength = facets.MinLength
	}
	if facets.MaxLength != 0 {
		s.MaxLength = facets.MaxLength
	}

	if facets.Pattern != "" {
		re := regexp.MustCompile(facets.Pattern)
		s.Pattern = re
	}

	return nil
}
