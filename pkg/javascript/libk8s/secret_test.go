package libk8s

import (
	"testing"
)

func TestSecret(t *testing.T) {
	runScriptFile(t, "testdata/secret_test.js")
}
