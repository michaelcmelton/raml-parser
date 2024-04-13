package types

import "gopkg.in/yaml.v3"

type BooleanScalar struct {
	Name string
	Type string
}

func (b *BooleanScalar) UnmarshalYAML(node *yaml.Node) error {
	b.Name = node.Content[0].Value
	b.Type = BOOLEAN

	return nil
}
