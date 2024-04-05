package types

import "gopkg.in/yaml.v3"

type IntegerScalar struct {
	BaseType
	numberScalarFacets
}

func (i *IntegerScalar) UnmarshalYAML(node *yaml.Node) error {
	i.Name = node.Content[0].Value
	i.Type = INTEGER

	properties := node.Content[1]
	if properties.Tag == "!!str" {
		return nil
	}

	var facets *numberScalarFacets

	err := properties.Decode(&facets)

	if err != nil {
		return err
	}

	i.numberScalarFacets = *facets

	return nil
}
