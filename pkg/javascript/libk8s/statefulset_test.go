package libk8s

import (
	"testing"
)

func TestStatefulSet(t *testing.T) {
	runScriptFile(t, "testdata/statefulset_test.js")
}
