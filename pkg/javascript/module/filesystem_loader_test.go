package module

import (
	"testing"
)

var _ FileLoader = FilesystemLoader("")

func TestFilesystemLoader(t *testing.T) {
	f := FilesystemLoader("./testdata")

	if f.IsDir("nosuchthing") {
		t.Error("Found nonexistent directory")
	}
	if f.IsFile("nosuchthing") {
		t.Error("Found nonexistent file")
	}

	if !f.IsFile("test_simplemodule.js") {
		t.Errorf("Expected to find test_simplemodule.js in %s", f)
	}
	if !f.IsDir("simplemodule") {
		t.Errorf("Expected to find simplemodule/ in %s", f)
	}
}
