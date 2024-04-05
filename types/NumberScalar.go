package types

import "gopkg.in/yaml.v3"

type NumberScalar struct {
	BaseType
	numberScalarFacets
}

type numberScalarFacets struct {
	Minimum    float64 `yaml:"minimum,omitempty"`
	Maximum    float64 `yaml:"maximum,omitempty"`
	Format     string  `yaml:"format,omitempty"`
	MultipleOf float64 `yaml:"multipleOf,omitempty"`
}

func (n *NumberScalar) UnmarshalYAML(node *yaml.Node) error {
	n.Name = node.Content[0].Value
	n.Type = NUMBER
	properties := node.Content[1]
	if properties.Tag == "!!str" {
		return nil
	}

	var facets *numberScalarFacets

	err := properties.Decode(&facets)

	if err != nil {
		return err
	}

	n.numberScalarFacets = *facets

	return nil
}
