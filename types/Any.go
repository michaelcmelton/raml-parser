package types

import "gopkg.in/yaml.v3"

type AnyType struct {
	BaseType
}

func (a *AnyType) UnmarshalYAML(node *yaml.Node) error {
	a.Name = node.Content[0].Value
	a.Type = ANY

	return nil
}
