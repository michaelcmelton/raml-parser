package types

type ScalarType = string

const (
	STRING ScalarType = "string"
)

type BaseScalar struct {
	Name string
	Type ScalarType
}
