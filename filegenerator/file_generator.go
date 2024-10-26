package filegenerator

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Loptt/infra-utils/template"
	"gopkg.in/yaml.v3"
)

// File represents a file to generate. It includes the name of the file and a
// collection of fields and values.
type File struct {
	Name   string
	Values map[string]string
}

// FileValues represents a list of file values. Each entry in the `Files`
// array represents a different file to generate. This struct is meant to be the
// result of parsing the yaml values file.
type FileValues struct {
	Files []File
}

func (fv FileValues) String() string {
	b := strings.Builder{}
	b.WriteString("Values:\n")

	for _, f := range fv.Files {
		b.Write([]byte(fmt.Sprintf("  File %s:\n", f.Name)))

		for k, v := range f.Values {
			b.Write([]byte(fmt.Sprintf("    Mapping %s => %s\n", k, v)))
		}
	}

	return b.String()
}

type FileGenerator struct {
	t  template.TemplateInterface
	fv *FileValues
}

type GeneratorResult struct {
	Name    string
	Content string
}

func (fg *FileGenerator) LoadValues(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to load values from %s: %v", path, err)
	}

	if err := yaml.Unmarshal(content, &fg.fv); err != nil {
		return fmt.Errorf("failed to parse yaml file %s: %v", path, err)
	}

	return nil
}

func (fg *FileGenerator) GenerateFiles() ([]GeneratorResult, error) {
	var results []GeneratorResult
	for _, f := range fg.fv.Files {
		content, err := fg.t.GenerateFile(f.Values)
		if err != nil {
			return results, fmt.Errorf("failed to generate file %s from template: %v", f.Name, err)
		}
		results = append(results, GeneratorResult{Name: f.Name, Content: content})
		log.Printf("Successfully generated file %s", f.Name)
	}

	return results, nil
}

func NewFileGenerator(t template.TemplateInterface) *FileGenerator {
	return &FileGenerator{t: t, fv: &FileValues{}}
}
