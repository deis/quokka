package module

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// FilesystemLoader is a simple filesystem-based FileLoader
//
// Functions here will evaluate path. If the path is absolute, the path is
// evaluated as-is. If the path is relative, it is joined with the parent
// path.
type FilesystemLoader string

func (f FilesystemLoader) IsDir(path string) bool {
	path = f.adjust(path)
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func (f FilesystemLoader) IsFile(path string) bool {
	path = f.adjust(path)
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	// FIXME
	return !fi.IsDir()
}

func (f FilesystemLoader) Load(path string) ([]byte, error) {
	path = f.adjust(path)
	return ioutil.ReadFile(path)
}

func (f FilesystemLoader) adjust(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(string(f), path)
}

func (f FilesystemLoader) String() string { return string(f) }
