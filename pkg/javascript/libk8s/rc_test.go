package libk8s

import (
	"testing"
)

func TestReplicationController(t *testing.T) {
	runScriptFile(t, "testdata/replicationcontroller_test.js")
}
