package libk8s

import (
	"testing"
)

func TestService(t *testing.T) {
	runScriptFile(t, "testdata/service_test.js")
}
