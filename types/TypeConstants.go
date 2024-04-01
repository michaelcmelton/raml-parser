package types

const (
	STRING       = "string"
	NUMBER       = "number"
	BOOLEAN      = "boolean"
	DATEONLY     = "date-only"
	TIMEONLY     = "time-only"
	DATETIMEONLY = "datetime-only"
	DATETIME     = "datetime"
	FILE         = "file"
	INTEGER      = "integer"
	NIL          = "nil"
	ANY          = "any"
)

type BaseType struct {
	Name string
	Type string
}
