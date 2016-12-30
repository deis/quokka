package libk8s

import (
	"testing"
)

func TestConfigMap(t *testing.T) {
	runScriptFile(t, "testdata/configmap_test.js")
}
