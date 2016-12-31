package libk8s

import (
	"testing"
)

func TestDeployment(t *testing.T) {
	runScriptFile(t, "testdata/deployment_test.js")
}
