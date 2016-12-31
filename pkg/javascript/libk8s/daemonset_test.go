package libk8s

import (
	"testing"
)

func TestDaemonSet(t *testing.T) {
	runScriptFile(t, "testdata/daemonset_test.js")
}
