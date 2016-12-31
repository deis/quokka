package libk8s

import (
	"testing"
)

func TestReplicaSet(t *testing.T) {
	runScriptFile(t, "testdata/replicaset_test.js")
}
