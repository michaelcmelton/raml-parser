package types

import "gopkg.in/yaml.v3"

type APIDefinition struct {
	Title string `yaml:"title"`
}

func (a *APIDefinition) UnmarshalYAML(node *yaml.Node) error {
	return nil
}
