package libk8s

import (
	"testing"
)

func TestJob(t *testing.T) {
	runScriptFile(t, "testdata/job_test.js")
}
