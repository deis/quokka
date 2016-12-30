package libk8s

import (
	"testing"
)

func TestNamespace(t *testing.T) {
	runScriptFile(t, "testdata/namespace_test.js")
}
