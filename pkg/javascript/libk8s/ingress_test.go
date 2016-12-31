package libk8s

import (
	"testing"
)

func TestIngress(t *testing.T) {
	runScriptFile(t, "testdata/ingress_test.js")
}
