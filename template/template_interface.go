package template

type TemplateInterface interface {
	GenerateFile(map[string]string) (string, error)
}
