package filegenerator

type FileGeneratorInterface interface {
	LoadValues(string) error
	GenerateFiles() ([]GeneratorResult, error)
}
