package libk8s

import (
	"testing"
)

func TestPod(t *testing.T) {
	runScriptFile(t, "testdata/pod_test.js")
}
