package file

import (
	"fmt"
	"os"
)

// FileManager object provices concrete implementations to perform file actions
// (read, write, etc.)
type FileManager struct{}

// Read function reads contents from the file specified in `path` and returns
// it as a string.
func (fm *FileManager) Read(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s with error: %v", path, err)
	}

	lastChar := data[len(data)-1]

	// If the last character is a newline, we want to delete it to prevent
	// issues when reading tokens.
	if lastChar == '\n' || lastChar == '\r' {
		data = data[:len(data)-1]
	}

	return string(data), nil
}
