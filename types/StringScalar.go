package types

import (
	"regexp"

	"gopkg.in/yaml.v3"
)

type StringScalar struct {
	BaseScalar
	Pattern   *regexp.Regexp
	MinLength int
	MaxLength int
}

type StringScalarFacets struct {
	Pattern   string `yaml:"pattern,omitempty"`
	MinLength int    `yaml:"minLength,omitempty"`
	MaxLength int    `yaml:"maxLength,omitempty"`
}

func (s *StringScalar) UnmarshalYAML(node *yaml.Node) error {
	var facets *StringScalarFacets
	s.Name = node.Content[0].Value
	s.Type = STRING
	properties := node.Content[1]

	err := properties.Decode(&facets)

	if err != nil {
		return err
	}

	s.MinLength = facets.MinLength
	s.MaxLength = facets.MaxLength

	re := regexp.MustCompile(facets.Pattern)

	s.Pattern = re

	return nil
}
