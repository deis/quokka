package libk8s

import (
	"testing"
)

func TestServiceAccount(t *testing.T) {
	runScriptFile(t, "testdata/serviceaccount_test.js")
}
