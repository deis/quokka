package libk8s

import (
	"io/ioutil"
	"testing"

	"github.com/robertkrimen/otto"
	"k8s.io/client-go/kubernetes/fake"
)

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
