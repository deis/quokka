package module

import (
	"testing"

	"github.com/robertkrimen/otto"
)

func TestModule_MinimalModule(t *testing.T) {

	loader := NewStrategicLoader(FilesystemLoader("./testdata"))
	for _, tt := range []struct {
		name string
		file string
	}{
		{"Minimal", "test_minimal.js"},
		{"Simple function", "test_simplemodule.js"},
	} {
		vm := otto.New()
		if err := Register(vm, loader); err != nil {
			t.Fatal(err)
		}

		script, err := loader.Load(tt.file)
		if err != nil {
			t.Fatalf("%s: %s", tt.name, err)
		}

		if _, err := vm.Run(script); err != nil {
			t.Fatalf("%s: %s", tt.name, err)
		}
	}

}

func TestModule_SimpleModule(t *testing.T) {

	loader := NewStrategicLoader(FilesystemLoader("./testdata"))

	vm := otto.New()
	if err := Register(vm, loader); err != nil {
		t.Fatal(err)
	}

	script, err := loader.Load("test_simplemodule.js")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := vm.Run(script); err != nil {
		t.Fatal(err)
	}
}
