package libk8s

import (
	"testing"

	"github.com/robertkrimen/otto"
	"github.com/technosophos/ottomatic"
	"k8s.io/client-go/kubernetes/fake"
)

func TestNewKubernetes(t *testing.T) {
	vm := otto.New()
	c := fake.NewSimpleClientset()
	ns := "foobar"

	k := NewKubernetes(vm, c, ns)
	nsv := k.GetNS()
	if kns, _ := nsv.ToString(); kns != ns {
		t.Errorf("Expected ns=%q, got %q", kns)
	}

	if k.GetNamespaces == nil {
		t.Errorf("GetNamespaces not initialized")
	}

	if k.Discovery == nil {
		t.Errorf("Discovery not initialized")
	}

	if k.Namespace == nil {
		t.Errorf("Namespace not initialized")
	}

	if k.WithNS == nil {
		t.Errorf("WithNS not initialized")
	}

	// This will panic if broken.
	k.GetNamespaces()
}

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
