package libk8s

import (
	"testing"
)

func TestDiscovery(t *testing.T) {
	res := runScript(t, "kubernetes.discovery.serverVersion().gitVersion")
	if gv, _ := res.ToString(); len(gv) < 6 {
		t.Errorf("Expected semver, got %q", gv)
	}

	// discovery/fake does not return useful (simulated) data for the following
	// cases:
	// - ServerGroups
	// - SwaggerSchema
	// - ServerPrefferredResources
	// - ServerPrefferedNamespacdResources
	// - ServerResources
	// - SwggerSchema (returns empty)
	//
	// Definitive testing for all of these can be found in the _functional_tests
	// suite.

	/*
		res = runScript(t, "kubernetes.discovery.serverResources().length")
		if v, _ := res.ToInteger(); v == 0 {
			t.Error("Expected at least one server resource")
		}
	*/

	/*
		res = runScript(t, "kubernetes.discovery.swaggerSchema('batch', 'v1').apiVersion")
		t.Logf("Res: %v", res)
		if v, _ := res.ToString(); v != "batch/v1" {
			t.Error("Expected batch/v1, got %s", v)
		}
	*/
}
