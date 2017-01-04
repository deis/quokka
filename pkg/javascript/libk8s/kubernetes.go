package libk8s

import (
	"github.com/robertkrimen/otto"
	"github.com/technosophos/ottomatic"

	"k8s.io/client-go/kubernetes"
)

// Kubernetes is the top-level object exposed to JavaScript.
//
// It provides access to the latest supported versions of the built-in API, along
// with partiall tooling for ThirdPartyResources.
//
// Not all of the API is exposed. Emphasis has been given to commonly-used APIs
// that follow a consistent pattern.
//
// In many cases, the APIs function by executing Go functions and then marshaling
// the results to their JSON equivalent, which can then be loaded into the
// JavaScript runtime as-is.
//
// Objects created in JavaScript are passed into Go as maps, and are then
// typed by unmarshaling them into their respective JSON struct.
type Kubernetes struct {
	// WithNS initializes another Kubernetes object, but with a different namespace.
	WithNS func(string) otto.Value `otto:"withNS"`
	// GetNS returns the default namespace, if set.
	GetNS func() otto.Value `otto:"getNS"`
	// GetNamespaces returns an array of all of the namespace names.
	// This is bubbled to the top so that scripts may discover the available
	// namespaces before creating namespaced objects.
	GetNamespaces func() otto.Value `otto:"getNamespaces"`
	// Discovery returns the Discovery object for whatever the stable discovery
	// implementation is. This is bubbled to the top so that the top-level
	// client can make decisions about which API version to use before actually
	// (intentionally) using the API.
	Discovery *Discovery `otto:"discovery"`
	// Namespace gives full access to the Namespace group.
	// This uses the most stable version of Namespaces (currently CoreV1). It is
	// exposed here so that clients can determine information about namespace
	// before working with the core API.
	Namespace *Namespace `otto:"namespace"`
}

func NewKubernetes(vm *otto.Otto, c kubernetes.Interface, ns string) *Kubernetes {
	k := &Kubernetes{
		WithNS: func(ns string) otto.Value {
			kk := NewClientSet(vm, c, ns)
			obj, err := vm.Object(`ns = {}`)
			poe(err)
			err = ottomatic.RegisterTo("x", kk, vm, obj)
			poe(err)
			res, err := obj.Get("x")
			poe(err)
			return res
		},
		GetNS:     func() otto.Value { r, _ := vm.ToValue(ns); return r },
		Discovery: NewDiscovery(c, vm),
		Namespace: NewNamespace(c, vm),
		GetNamespaces: func() otto.Value {
			return listAllNamespaces(c, vm)
		},
	}
	return k
}

// Register registers the top-level Kubernetes API objects with the JS runtime.
func Register(vm *otto.Otto) error {
	c, err := kubeClient()
	if err != nil {
		return err
	}
	return RegisterWithClient(vm, c)
}

// RegisterWithClient registers the Kubernetes object to an existing Kubernetes client.
func RegisterWithClient(vm *otto.Otto, c kubernetes.Interface) error {
	k := NewKubernetes(vm, c, "default")
	return ottomatic.Register("kubernetes", k, vm)
}
