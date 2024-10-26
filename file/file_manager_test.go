package file

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

const testFile = "test.txt"

func TestRead(t *testing.T) {
	tmpDir := os.TempDir()
	tmpFile := filepath.Join(tmpDir, testFile)
	defer func() {
		if err := os.Remove(tmpFile); err != nil {
			log.Fatalf("failed to remove file %s as part of cleanup: %v", tmpFile, err)
		}
	}()

	tests := []struct {
		description string
		setup       func(string) error
		f           FileManagerInterface
		path        string
		input       string
		want        string
		want_err    error
	}{
		{
			description: "Test read contents from file",
			setup: func(contents string) error {
				if err := os.WriteFile(tmpFile, []byte(contents), os.FileMode(0777)); err != nil {
					return fmt.Errorf("failed to setup test file: %v", err)
				}
				return nil
			},
			f:        &FileManager{},
			path:     tmpFile,
			input:    "Hello world!",
			want:     "Hello world!",
			want_err: nil,
		},
		{
			description: "Test read from nonexistent file",
			setup: func(contents string) error {
				// Do setup anything here.
				return nil
			},
			f:        &FileManager{},
			path:     "/fake/path.txt",
			input:    "Hello world!",
			want:     "",
			want_err: errors.New("Error!"),
		},
		{
			description: "Test read contents with newline at the end",
			setup: func(contents string) error {
				if err := os.WriteFile(tmpFile, []byte(contents), os.FileMode(0777)); err != nil {
					return fmt.Errorf("failed to setup test file: %v", err)
				}
				return nil
			},
			f:        &FileManager{},
			path:     tmpFile,
			input:    "Hola mundo!\n",
			want:     "Hola mundo!",
			want_err: nil,
		},
	}

	for i, test := range tests {
		test.setup(test.input)
		got, err := test.f.Read(test.path)

		if test.want_err != nil {
			if err == nil {
				t.Errorf("Test #%d %s: want err %v, got %v", i, test.description, test.want_err, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test #%d %s: found error %v", i, test.description, err)
			} else if got != test.want {
				t.Errorf("Test #%d %s: got %v, want %v", i, test.description, got, test.want)
			}
		}
	}
}
