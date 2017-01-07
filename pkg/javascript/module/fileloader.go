package module

// FileLoader describes something able to load a file by name.
//
// This returns the data as a []byte, and an error under any condition in which
// the file cannot be completely read.
//
// It should not be assumed that the file returned from a loader can be found
// on the filesystem.
type FileLoader interface {
	// Load loads a file.
	//
	// If the file is not found, this returns os.ErrNotExist. Other errors
	// are possible for other read failures.
	Load(filename string) ([]byte, error)
	IsFile(filename string) bool
	IsDir(filename string) bool
}
