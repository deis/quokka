package libk8s

import (
	"testing"
)

func TestThirdPartyResource(t *testing.T) {
	runScriptFile(t, "testdata/thirdpartyresource_test.js")
}
