package template

import (
	"fmt"
	"strings"
)

type Template struct {
	content string
}

func (t *Template) GenerateFile(valueMap map[string]string) (string, error) {
	generated := t.content

	for k, v := range valueMap {
		match := fmt.Sprintf("${%s}", k)
		if !strings.Contains(generated, match) {
			return "", fmt.Errorf("variable %s is not present in template", k)
		}
		generated = strings.ReplaceAll(generated, match, v)
	}

	return generated, nil
}

func NewTemplate(content string) *Template {
	return &Template{content: content}
}
