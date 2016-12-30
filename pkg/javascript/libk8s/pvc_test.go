package libk8s

import (
	"testing"
)

func TestPersistentVolumeClaim(t *testing.T) {
	runScriptFile(t, "testdata/persistentvolumeclaim_test.js")
}
