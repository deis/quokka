package module

import (
	"fmt"
	"os"
)

// StrategicLoader implements the different load strategies (directory, file, etc.)
//
// This can be wrapped around a basic loader, and it will use internal strategies
// to load files according to the Module interface for Node.js.
type StrategicLoader struct {
	basicLoader FileLoader
}

func NewStrategicLoader(loader FileLoader) *StrategicLoader {
	return &StrategicLoader{basicLoader: loader}
}
func (s *StrategicLoader) IsFile(name string) bool { return s.basicLoader.IsFile(name) }
func (s *StrategicLoader) IsDir(name string) bool  { return s.basicLoader.IsDir(name) }

func (s *StrategicLoader) Load(filename string) ([]byte, error) {
	loader := s.basicLoader
	// Right now, we're not very strategic; we just load *.js files.
	// TODO: handle directories by loading package.json. See https://nodejs.org/api/modules.html
	if loader.IsFile(filename) {
		return loader.Load(filename)
	}
	// TODO: Need to make sure filename isn't a directory before we go around
	// adding extensions to it.
	if loader.IsFile(filename + ".js") {
		return loader.Load(filename + ".js")
	}
	return []byte{}, fmt.Errorf("%s %s", os.ErrNotExist, filename)
}
