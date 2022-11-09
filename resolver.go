package ramlparser

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

var (
	ErrOpenFile string = "error occurred opening file: (%w)"
	ErrReadFile string = "error occurred reading file: (%w)"
)

// FileExists is an function extraction for checking if a file exists at a given path.
func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return errors.Is(err, fs.ErrNotExist)
}

// ResolveFile gathers file contents from a file passed in and
// returns a byte slice of the file's contents.
func ResolveFile(filepath string) ([]byte, error) {
	if !FileExists(filepath) {
		return nil, fmt.Errorf(ErrFileNotFound, filepath)
	}

	file, err := os.Open(filepath)

	if err != nil {
		return nil, fmt.Errorf(ErrOpenFile, err)
	}

	defer file.Close()

	contents, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, fmt.Errorf(ErrReadFile, err)
	}

	return contents, nil
}
