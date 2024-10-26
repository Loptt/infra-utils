package deployinfo

// DeployInfoInterface is an abstract representation of an object used to
// manage deploy information.
type DeployInfoInterface interface {
	ProdLocations() []string
	StagingLocations() []string
}
