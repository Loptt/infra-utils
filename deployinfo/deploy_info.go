package deployinfo

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// DeployInfo represents information for a product's deployment.
type DeployInfo struct {
	data DeployData
}

// DeployLocations returns all the deploy locations for the prod environment.
func (d *DeployInfo) ProdLocations() []string {
	return d.data.ProdLocations
}

// DeployLocations returns all the deploy locations for the staging
// environment.
func (d *DeployInfo) StagingLocations() []string {
	return d.data.StagingLocations
}

// ComposeInformation returns all the data necessary to manage compose files at
// the deploy locations.
func (d *DeployInfo) ComposeInformation() ComposeInformation {
	return d.data.CI
}

// String representation of DeployInfo. It is human readable but not machine
// readable.
func (d DeployInfo) String() string {
	b := strings.Builder{}
	b.WriteString("Deploy Info:\n")

	b.WriteString("  Prod Locations:\n")

	for i, v := range d.ProdLocations() {
		b.Write([]byte(fmt.Sprintf("    Location %d: %s\n", i, v)))
	}

	b.WriteString("  Staging Locations:\n")

	for i, v := range d.StagingLocations() {
		b.Write([]byte(fmt.Sprintf("    Location %d: %s\n", i, v)))
	}

	b.WriteString("  Compose Information:\n")
	b.WriteString(fmt.Sprintf("    Values File: %s\n", d.data.CI.ValuesFile))
	b.WriteString(fmt.Sprintf("    Template File: %s\n", d.data.CI.TemplateFile))

	return b.String()
}

func parseDeployData(rawData string) (DeployData, error) {
	data := DeployData{}
	if err := yaml.Unmarshal([]byte(rawData), &data); err != nil {
		return data, fmt.Errorf("failed to unmarshal raw data: %v", err)
	}

	return data, nil
}

// NewdeplotInfo creates a new DeployInfo struct based on a YAML file specified
// by the path argument. The file argument specifies a FileManager object to use.
func NewDeployInfo(content string) (*DeployInfo, error) {
	data, err := parseDeployData(content)
	if err != nil {
		return nil, err
	}

	return &DeployInfo{data: data}, nil
}
