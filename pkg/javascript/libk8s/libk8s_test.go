package libk8s

import (
	"io/ioutil"
	"testing"

	"github.com/robertkrimen/otto"
	"github.com/technosophos/ottomatic"
	"k8s.io/client-go/kubernetes/fake"
)

func TestRegister(t *testing.T) {
	vm := otto.New()
	if err := Register(vm); err != nil {
		t.Fatal(err)
	}

	if _, err := vm.Get("kubernetes"); err != nil {
		t.Fatalf("Error getting kube: %s", err)
	}
}

func TestRegisterWithClient(t *testing.T) {

	c := fake.NewSimpleClientset()

	vm := otto.New()
	if err := RegisterWithClient(vm, c); err != nil {
		t.Fatal(err)
	}

	if _, err := vm.Get("kubernetes"); err != nil {
		t.Fatalf("Error getting kube: %s", err)
	}

	if _, err := vm.Run("ver = kubernetes.discovery.serverVersion();"); err != nil {
		t.Errorf("script did not exit cleanly: %s", err)
	}

	if d, err := ottomatic.DeepGet("ver.gitVersion", vm); err != nil {
		t.Error(err)
	} else if gitVer, _ := d.ToString(); len(gitVer) == 0 {
		t.Errorf("Expected string %q to be longer than 0", gitVer)
	} else {
		t.Logf("Git version: %q", gitVer)
	}
}

func fakeVM(t *testing.T) *otto.Otto {
	c := fake.NewSimpleClientset()

	vm := otto.New()
	if err := RegisterWithClient(vm, c); err != nil {
		t.Fatal(err)
	}
	return vm
}

func runScript(t *testing.T, script string) otto.Value {
	vm := fakeVM(t)
	out, err := vm.Run(script)
	if err != nil {
		t.Errorf("script failed with %q:\n\t%s", err, script)
	}
	return out
}

func runScriptFile(t *testing.T, file string) otto.Value {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatalf("failed to read %s: %s", file, err)
	}
	return runScript(t, string(data))
}
