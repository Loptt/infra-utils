package file

// FileManagerInterface defines an interface to interact with the file system.
type FileManagerInterface interface {
	Read(string) (string, error)
}
