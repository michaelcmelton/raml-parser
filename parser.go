package ramlparser

import (
	"fmt"

	"github.com/michaelcmelton/raml-parser/types"
)

var (
	ErrFileNotFound string = "no file found at the path provided: %s"
)

// ParseSpecification is responsible for parsing the file passed,
// and returning the full RAML specification contained in the file.
func ParseSpecification(filePath string) (*types.APISpecification, error) {

	// Create spec variable.
	var ramlSpec *types.APISpecification
	// Check if file exists at path passed in.
	if !FileExists(filePath) {
		return nil, fmt.Errorf(fmt.Sprintf(ErrFileNotFound, filePath))
	}
	// Open file.
	// Resolve all includes and append.
	// Parse file and return spec.
	return ramlSpec, nil
}
