package file

// FileManagerFake provides a fake implementation of the FileManagerInterface.
type FileManagerFake struct {
	Content string
	Err     error
}

// Read returns the predefined content speficied in `f.content`. If `f.err` is
// defined, then it returns the error instead.
func (f *FileManagerFake) Read(path string) (string, error) {
	if f.Err != nil {
		return "", f.Err
	}
	return f.Content, nil
}
