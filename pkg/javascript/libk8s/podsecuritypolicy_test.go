package libk8s

import (
	"testing"
)

func TestPodSecurityPolicy(t *testing.T) {
	runScriptFile(t, "testdata/podsecuritypolicy_test.js")
}
