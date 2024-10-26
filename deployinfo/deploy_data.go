package deployinfo

// ComposeInformation represents the data require to manage the compose files
// present at each deploy location.
type ComposeInformation struct {
	ValuesFile   string `yaml:"values-file"`
	TemplateFile string `yaml:"template-file"`
}

// DeployData struct represents the structure of a deploy information yaml
// file. And thus, it is required to be unmarshable by this structure.
type DeployData struct {
	ProdLocations    []string           `yaml:"prod-locations"`
	StagingLocations []string           `yaml:"staging-locations"`
	CI               ComposeInformation `yaml:"compose-information"`
}
