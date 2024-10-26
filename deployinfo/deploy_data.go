package deployinfo

type DeployData struct {
	ProdLocations    []string `yaml:"prod-locations"`
	StagingLocations []string `yaml:"staging-locations"`
}
