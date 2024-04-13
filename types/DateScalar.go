package types

import "gopkg.in/yaml.v3"

type DateScalar struct {
	Name   string
	Type   string
	Format string
}

type dateScalarFacets struct {
	Type   string `yaml:"type"`
	Format string `yaml:"format,omitempty"`
}

func (d *DateScalar) UnmarshalYAML(node *yaml.Node) error {
	d.Name = node.Content[0].Value

	properties := node.Content[1]

	if node.Content[1].Tag == "!!str" {
		d.Type = getTypeString(node.Content[1].Value)
		return nil
	}
	var dateFacets *dateScalarFacets

	err := properties.Decode(&dateFacets)

	if err != nil {
		return err
	}

	d.Type = getTypeString(dateFacets.Type)

	if d.Type == DATETIME {
		d.Format = dateFacets.Format
	}

	return nil
}

func getTypeString(dateType string) string {
	switch dateType {
	case "date-only":
		return DATEONLY
	case "time-only":
		return TIMEONLY
	case "datetime-only":
		return DATETIMEONLY
	}

	return DATETIME
}
